package bitcoin

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"math/big"
	"nn-blockchain-api/pkg/errors"
)

//go:generate mockgen -source=transaction.go -destination=mocks/transaction_mock.go

type UnspentList struct {
	TxId         string `json:"txid"`
	Vout         int64  `json:"vout"`
	ScriptPubKey string `json:"scriptPubKey"`
}

type DecodedTx struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`

		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm     string `json:"asm"`
			Hex     string `json:"hex"`
			Address string `json:"address"`
			Type    string `json:"type"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
}

type UTXO []struct {
	TxId     string
	Vout     int64
	Amount   int64
	PKScript string
}

type TransactionService interface {
	CreateTransaction(utxos UTXO, fromAddress, toAddress string, amount int64, network string) (*string, *float64, error)
	//CreateTransaction(inputs []map[string]interface{}, outputs []map[string]string, network string) (string, error)
	DecodeTransaction(tx string, network string) (*DecodedTx, error)
	FundForTransaction(createdTx, changeAddress, network string) (string, *float64, error)
	SignTransaction(tx, privateKey string, utxos UTXO, network string) (string, error)
	SendTransaction(signedTx, network string) (string, error)
}

type transactionService struct {
	btcClient IBtcClient
}

func NewTransactionService(btcClient IBtcClient) (TransactionService, error) {
	if btcClient == nil {
		return nil, errors.NewInternal("invalid btc client")
	}
	return &transactionService{btcClient: btcClient}, nil
}

//func (svc *transactionService) CreateTransaction(client IBtcClient, inputs []map[string]interface{}, outputs []map[string]string, network string) (string, error) {
//	req := BaseRequest{
//		JsonRpc: "2.0",
//		Method:  "createrawtransaction",
//		Params:  []interface{}{inputs, outputs},
//	}
//
//	msg := struct {
//		Result string `json:"result"`
//		Error  struct {
//			Message string `json:"message"`
//		} `json:"error"`
//	}{}
//
//	body, err := client.EncodeBaseRequest(req)
//	if err != nil {
//		return "", err
//	}
//
//	response, err := client.Send(body, "", network)
//	if err != nil {
//		return "", err
//	}
//
//	defer response.Body.Close()
//
//	err = json.NewDecoder(response.Body).Decode(&msg)
//	if err != nil {
//		return "", err
//	}
//
//	if msg.Error.Message != "" {
//		return "", errors.NewInternal(msg.Error.Message)
//	}
//
//	return msg.Result, nil
//}

func (svc *transactionService) CreateTransaction(utxos UTXO, fromAddress, toAddress string, amount int64, network string) (*string, *float64, error) {
	chainParams := &chaincfg.TestNet3Params

	// Get fee
	feeRate, err := GetCurrentFeeRate(svc.btcClient, network)
	if err != nil {
		return nil, nil, errors.NewInternal(err.Error())
	}

	// Calculate all unspent amount
	utxosAmount := big.NewInt(0)
	for idx := range utxos {
		utxosAmount.Add(utxosAmount, new(big.Int).SetInt64(utxos[idx].Amount))
	}

	// Init transaction
	tx := wire.NewMsgTx(2)

	// prepare transaction inputs
	sourceUtxosAmount := big.NewInt(0)
	var sourceUTXOs []*UnspentList
	for idx := range utxos {
		hashStr := utxos[idx].TxId
		sourceUtxosAmount.Add(sourceUtxosAmount, new(big.Int).SetInt64(utxos[idx].Amount))

		sourceUTXOHash, err := chainhash.NewHashFromStr(hashStr)
		if err != nil {
			return nil, nil, errors.NewInternal(err.Error())
		}

		if amount <= sourceUtxosAmount.Int64() {
			sourceUTXOIndex := uint32(utxos[idx].Vout)
			sourceUTXO := wire.NewOutPoint(sourceUTXOHash, sourceUTXOIndex)
			sourceUTXOs = append(sourceUTXOs, &UnspentList{
				TxId:         utxos[idx].TxId,
				Vout:         utxos[idx].Vout,
				ScriptPubKey: utxos[idx].PKScript,
			})
			sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)

			tx.AddTxIn(sourceTxIn)
			break
		}

		sourceUTXOIndex := uint32(utxos[idx].Vout)
		sourceUTXO := wire.NewOutPoint(sourceUTXOHash, sourceUTXOIndex)
		sourceUTXOs = append(sourceUTXOs, &UnspentList{
			TxId:         utxos[idx].TxId,
			Vout:         utxos[idx].Vout,
			ScriptPubKey: utxos[idx].PKScript,
		})
		sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)

		tx.AddTxIn(sourceTxIn)
	}

	// create the transaction outputs
	destAddress, err := btcutil.DecodeAddress(toAddress, chainParams)
	if err != nil {
		return nil, nil, errors.NewInternal(err.Error())
	}

	destScript, err := txscript.PayToAddrScript(destAddress)
	if err != nil {
		return nil, nil, errors.NewInternal(err.Error())
	}

	// tx out to send btc to user
	destOutput := wire.NewTxOut(amount, destScript)
	tx.AddTxOut(destOutput)

	change := new(big.Int).Set(sourceUtxosAmount)
	change = new(big.Int).Sub(change, new(big.Int).SetInt64(amount))
	//change = new(big.Int).Sub(change, totalFee)
	if change.Cmp(big.NewInt(0)) == -1 {
		//svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, nil, errors.NewInternal(err.Error())
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
		changeSendToAddress, err := btcutil.DecodeAddress(fromAddress, chainParams)
		if err != nil {
			//svc.log.WithContext(ctx).Errorf(err.Error())
			return nil, nil, errors.NewInternal(err.Error())
		}

		changeSendToScript, err := txscript.PayToAddrScript(changeSendToAddress)
		if err != nil {
			//svc.log.WithContext(ctx).Errorf(err.Error())
			return nil, nil, errors.NewInternal(err.Error())
		}

		//tx out to send change back to us
		changeOutput := wire.NewTxOut(change.Int64(), changeSendToScript)
		tx.AddTxOut(changeOutput)
	}

	// calculate fees
	txByteSize := big.NewInt(int64(len(tx.TxIn)*180 + len(tx.TxOut)*34 + 10 + len(tx.TxIn)))
	totalFee := new(big.Int).Mul(feeRate, txByteSize)

	// Need add fee to spend amount and then compare
	if (amount - totalFee.Int64()) >= sourceUtxosAmount.Int64() {
		//log.Fatal(errors.New("your balance too low for this transaction"))
		//svc.log.WithContext(ctx).Errorf("your balance too low for this transaction")
		return nil, nil, errors.NewInternal("your balance too low for this transaction")
	}

	//log.Printf("%-18s %s\n", "total fee:", totalFee)

	// Change amount of source output transaction
	tx.TxOut[0].Value = amount - totalFee.Int64()

	// Transaction Hash
	notSignedTxBuf := bytes.NewBuffer(make([]byte, 0, tx.SerializeSize()))
	err = tx.Serialize(notSignedTxBuf)
	if err != nil {
		//log.Fatal(err)
		//svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, nil, errors.NewInternal(err.Error())
	}

	createdTx := hex.EncodeToString(notSignedTxBuf.Bytes())
	btcAmountFee := btcutil.Amount(totalFee.Int64()).ToBTC()
	return &createdTx, &btcAmountFee, nil
}

func (svc *transactionService) DecodeTransaction(tx string, network string) (*DecodedTx, error) {
	msg := struct {
		Result DecodedTx `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "decoderawtransaction",
		Params:  []interface{}{tx},
	}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := svc.btcClient.Send(body, "", network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.NewInternal(msg.Error.Message)
	}

	return &DecodedTx{
		Txid:     msg.Result.Txid,
		Hash:     msg.Result.Hash,
		Version:  msg.Result.Version,
		Size:     msg.Result.Size,
		Vsize:    msg.Result.Vsize,
		Weight:   msg.Result.Weight,
		Locktime: msg.Result.Locktime,
		Vin:      msg.Result.Vin,
		Vout:     msg.Result.Vout,
	}, nil
}

func (svc *transactionService) FundForTransaction(createdTx, changeAddress, network string) (string, *float64, error) {
	subtractFeeFromOutputs := []int64{0}

	params := map[string]interface{}{
		"changeAddress":          changeAddress,
		"subtractFeeFromOutputs": subtractFeeFromOutputs,
	}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "fundrawtransaction",
		Params:  []interface{}{createdTx, params},
	}

	msg := struct {
		Result struct {
			Hex string  `json:"hex"`
			Fee float64 `json:"fee"`
		} `json:"result"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return "", nil, err
	}

	response, err := svc.btcClient.Send(body, "", network)
	if err != nil {
		return "", nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", nil, err
	}

	if msg.Error.Message != "" {
		return "", nil, errors.NewInternal(msg.Error.Message)
	}

	return msg.Result.Hex, &msg.Result.Fee, nil
}

func (svc *transactionService) SignTransaction(tx, privateKey string, utxos UTXO, network string) (string, error) {
	privateKeyArray := []string{privateKey}
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "signrawtransactionwithkey",
		Params:  []interface{}{tx, privateKeyArray, utxos},
	}

	msg := struct {
		Result struct {
			Hex      string `json:"hex"`
			Complete bool   `json:"complete"`
		} `json:"result"`
		Error struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return "", err
	}

	response, err := svc.btcClient.Send(body, "", network)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Error.Message != "" {
		return "", errors.NewInternal(msg.Error.Message)
	}

	if !msg.Result.Complete {
		return "", errors.NewInternal("signing transaction not complete. Please try again")
	}

	return msg.Result.Hex, nil
}

func (svc *transactionService) SendTransaction(signedTx, network string) (string, error) {
	msg := struct {
		Result string `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "sendrawtransaction",
		Params:  []interface{}{signedTx},
	}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return "", err
	}

	response, err := svc.btcClient.Send(body, "", network)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Error.Message != "" {
		return "", errors.NewInternal(msg.Error.Message)
	}

	return msg.Result, nil
}
