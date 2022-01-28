package bitcoin

import (
	"bytes"
	"context"
	"encoding/hex"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/sirupsen/logrus"
	"math/big"
	"nn-blockchain-api/pkg/errors"
	"nn-blockchain-api/pkg/rpc/bitcoin"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type UnspentList struct {
	TxId         string `json:"txid"`
	Vout         int64  `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
}

type Service interface {
	StatusNode(ctx context.Context, dto *StatusNodeDTO) (*StatusNodeInfoDTO, error)
	CreateTransaction(ctx context.Context, dto *CreateRawTransactionDTO) (*CreatedRawTransactionDTO, error)
	DecodeTransaction(ctx context.Context, dto *DecodeRawTransactionDTO) (*DecodedRawTransactionDTO, error)
	FoundForRawTransaction(ctx context.Context, dto *FundForRawTransactionDTO) (*FundedRawTransactionDTO, error)
	SignTransaction(ctx context.Context, dto *SignRawTransactionDTO) (*SignedRawTransactionDTO, error)
	SendTransaction(ctx context.Context, dto *SendRawTransactionDTO) (*SentRawTransactionDTO, error)
	ImportAddress(ctx context.Context, dto *ImportAddressDTO) error
}

type service struct {
	log       *logrus.Logger
	btcClient bitcoin.IBtcClient
}

func NewService(log *logrus.Logger, btcClient bitcoin.IBtcClient) (Service, error) {
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	if btcClient == nil {
		return nil, errors.NewInternal("invalid btc client")
	}
	return &service{log: log, btcClient: btcClient}, nil
}

func (svc *service) StatusNode(ctx context.Context, dto *StatusNodeDTO) (*StatusNodeInfoDTO, error) {
	status, err := bitcoin.Status(svc.btcClient, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf("failed check node status")
		return nil, errors.NewInternal("failed check node status")
	}

	return &StatusNodeInfoDTO{
		Chain:                status.Chain,
		Blocks:               status.Blocks,
		Headers:              status.Headers,
		Verificationprogress: status.Verificationprogress,
		Softforks:            status.Softforks,
		Warnings:             status.Warnings,
	}, nil
}

func (svc *service) CreateTransaction(ctx context.Context, dto *CreateRawTransactionDTO) (*CreatedRawTransactionDTO, error) {
	chainParams := &chaincfg.TestNet3Params

	// Get fee
	feeRate, err := bitcoin.GetCurrentFeeRate(svc.btcClient, dto.Network)
	//log.Printf("%-18s %s\n", "current fee rate:", feeRate)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	// Calculate all unspent amount
	utxosAmount := big.NewInt(0)
	for idx := range dto.Utxo {
		utxosAmount.Add(utxosAmount, new(big.Int).SetInt64(dto.Utxo[idx].Amount))
	}

	// Init transaction
	tx := wire.NewMsgTx(2)

	// prepare transaction inputs
	sourceUtxosAmount := big.NewInt(0)
	var sourceUTXOs []*UnspentList
	for idx := range dto.Utxo {
		hashStr := dto.Utxo[idx].TxId
		sourceUtxosAmount.Add(sourceUtxosAmount, new(big.Int).SetInt64(dto.Utxo[idx].Amount))

		sourceUTXOHash, err := chainhash.NewHashFromStr(hashStr)
		if err != nil {
			svc.log.WithContext(ctx).Errorf(err.Error())
			return nil, errors.NewInternal(err.Error())
		}

		if dto.Amount <= sourceUtxosAmount.Int64() {
			sourceUTXOIndex := uint32(dto.Utxo[idx].Vout)
			sourceUTXO := wire.NewOutPoint(sourceUTXOHash, sourceUTXOIndex)
			sourceUTXOs = append(sourceUTXOs, &UnspentList{
				TxId:         dto.Utxo[idx].TxId,
				Vout:         dto.Utxo[idx].Vout,
				ScriptPubKey: dto.Utxo[idx].PKScript,
			})
			sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)

			tx.AddTxIn(sourceTxIn)
			break
		}

		sourceUTXOIndex := uint32(dto.Utxo[idx].Vout)
		sourceUTXO := wire.NewOutPoint(sourceUTXOHash, sourceUTXOIndex)
		sourceUTXOs = append(sourceUTXOs, &UnspentList{
			TxId:         dto.Utxo[idx].TxId,
			Vout:         dto.Utxo[idx].Vout,
			ScriptPubKey: dto.Utxo[idx].PKScript,
		})
		sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)

		tx.AddTxIn(sourceTxIn)
	}

	// create the transaction outputs
	destAddress, err := btcutil.DecodeAddress(dto.ToAddress, chainParams)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	destScript, err := txscript.PayToAddrScript(destAddress)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	// tx out to send btc to user
	destOutput := wire.NewTxOut(dto.Amount, destScript)
	tx.AddTxOut(destOutput)

	change := new(big.Int).Set(sourceUtxosAmount)
	change = new(big.Int).Sub(change, new(big.Int).SetInt64(dto.Amount))
	//change = new(big.Int).Sub(change, totalFee)
	if change.Cmp(big.NewInt(0)) == -1 {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	if change.Int64() != 0 {
		// our fee address
		//feeSendToAddress, err := btcutil.DecodeAddress(fromWalletPublicAddress, chainParams)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//feeSendToScript, err := txscript.PayToAddrScript(feeSendToAddress)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		////tx out to send change back to us
		//feeOutput := wire.NewTxOut(changeFee.Int64(), feeSendToScript)
		//tx.AddTxOut(feeOutput)

		// our change address
		changeSendToAddress, err := btcutil.DecodeAddress(dto.FromAddress, chainParams)
		if err != nil {
			svc.log.WithContext(ctx).Errorf(err.Error())
			return nil, errors.NewInternal(err.Error())
		}

		changeSendToScript, err := txscript.PayToAddrScript(changeSendToAddress)
		if err != nil {
			svc.log.WithContext(ctx).Errorf(err.Error())
			return nil, errors.NewInternal(err.Error())
		}

		//tx out to send change back to us
		changeOutput := wire.NewTxOut(change.Int64(), changeSendToScript)
		tx.AddTxOut(changeOutput)
	}

	// calculate fees
	txByteSize := big.NewInt(int64(len(tx.TxIn)*180 + len(tx.TxOut)*34 + 10 + len(tx.TxIn)))
	totalFee := new(big.Int).Mul(feeRate, txByteSize)

	// Need add fee to spend amount and then compare
	if (dto.Amount - totalFee.Int64()) >= sourceUtxosAmount.Int64() {
		//log.Fatal(errors.New("your balance too low for this transaction"))
		svc.log.WithContext(ctx).Errorf("your balance too low for this transaction")
		return nil, errors.NewInternal("your balance too low for this transaction")
	}

	//log.Printf("%-18s %s\n", "total fee:", totalFee)

	// Change amount of source output transaction
	tx.TxOut[0].Value = dto.Amount - totalFee.Int64()

	// Transaction Hash
	notSignedTxBuf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	err = tx.Serialize(notSignedTxBuf)
	if err != nil {
		//log.Fatal(err)
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	return &CreatedRawTransactionDTO{
		Tx:  hex.EncodeToString(notSignedTxBuf.Bytes()),
		Fee: btcutil.Amount(totalFee.Int64()).ToBTC(),
	}, nil
}

func (svc *service) DecodeTransaction(ctx context.Context, dto *DecodeRawTransactionDTO) (*DecodedRawTransactionDTO, error) {
	decodedTx, err := bitcoin.DecodeTx(svc.btcClient, dto.Tx, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	return &DecodedRawTransactionDTO{
		Txid:     decodedTx.Txid,
		Hash:     decodedTx.Hash,
		Version:  decodedTx.Version,
		Size:     decodedTx.Size,
		Vsize:    decodedTx.Vsize,
		Weight:   decodedTx.Weight,
		Locktime: decodedTx.Locktime,
		Vin:      decodedTx.Vin,
		Vout:     decodedTx.Vout,
	}, nil
}

func (svc *service) FoundForRawTransaction(ctx context.Context, dto *FundForRawTransactionDTO) (*FundedRawTransactionDTO, error) {
	tx, fee, err := bitcoin.FundForRawTransaction(svc.btcClient, dto.CreatedTxHex, dto.ChangeAddress, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	return &FundedRawTransactionDTO{
		Tx:  tx,
		Fee: *fee,
	}, nil
}

func (svc *service) SignTransaction(ctx context.Context, dto *SignRawTransactionDTO) (*SignedRawTransactionDTO, error) {
	var utxos []map[string]interface{}
	for _, s := range dto.Utxo {
		utxos = append(utxos, map[string]interface{}{"txid": s.TxId, "vout": s.Vout, "scriptPubKey": s.PKScript, "amount": s.Amount})
	}

	tx, err := bitcoin.SignTx(svc.btcClient, dto.Tx, dto.PrivateKey, utxos, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	return &SignedRawTransactionDTO{
		Hash: tx,
	}, nil
}

func (svc *service) SendTransaction(ctx context.Context, dto *SendRawTransactionDTO) (*SentRawTransactionDTO, error) {
	txId, err := bitcoin.SendTx(svc.btcClient, dto.SignedTx, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.NewInternal(err.Error())
	}

	return &SentRawTransactionDTO{
		TxId: txId,
	}, nil
}

func (svc *service) ImportAddress(ctx context.Context, dto *ImportAddressDTO) error {
	err := bitcoin.ImportAddress(svc.btcClient, dto.Address, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return errors.NewInternal(err.Error())
	}

	return nil
}
