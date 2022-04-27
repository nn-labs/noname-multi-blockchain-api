package bitcoin

import (
	"context"
	"nn-blockchain-api/pkg/errors"
	rpc_bitcoin "nn-blockchain-api/pkg/rpc/bitcoin"

	"github.com/sirupsen/logrus"
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

	WalletInfo(ctx context.Context, dto *WalletDTO) (*WalletInfoDTO, error)
	CreateWallet(ctx context.Context, dto *CreateWalletDTO) (*CreatedWalletInfoDTO, error)
	LoadWaller(ctx context.Context, dto *LoadWalletDTO) (*LoadWalletInfoDTO, error)
	ImportAddress(ctx context.Context, dto *ImportAddressDTO) (*ImportAddressInfoDTO, error)
	RescanWallet(ctx context.Context, dto *RescanWalletDTO) (*RescanWalletInfoDTO, error)
	ListUnspent(ctx context.Context, dto *ListUnspentDTO) (*ListUnspentInfoDTO, error)
}

type service struct {
	log       *logrus.Logger
	btcRpcSvc rpc_bitcoin.Service
}

func NewService(btcRpcSvc rpc_bitcoin.Service, log *logrus.Logger) (Service, error) {
	if btcRpcSvc == nil {
		return nil, errors.NewInternal("invalid btc service")
	}
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{btcRpcSvc: btcRpcSvc, log: log}, nil
}

func (svc *service) StatusNode(ctx context.Context, dto *StatusNodeDTO) (*StatusNodeInfoDTO, error) {
	status, err := svc.btcRpcSvc.Status(ctx, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf("failed check node status")
		return nil, errors.WithMessage(ErrFailedGetStatusNode, err.Error())
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
	tx, fee, err := svc.btcRpcSvc.CreateTransaction(ctx, rpc_bitcoin.UTXO(dto.Utxo), dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedCreateTx, err.Error())
	}

	return &CreatedRawTransactionDTO{
		Tx:  *tx,
		Fee: *fee,
	}, nil
}

func (svc *service) DecodeTransaction(ctx context.Context, dto *DecodeRawTransactionDTO) (*DecodedRawTransactionDTO, error) {
	decodedTx, err := svc.btcRpcSvc.DecodeTransaction(ctx, dto.Tx, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedDecodeTx, err.Error())
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
	tx, fee, err := svc.btcRpcSvc.FundForTransaction(ctx, dto.CreatedTxHex, dto.ChangeAddress, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedFundForTx, err.Error())
	}

	return &FundedRawTransactionDTO{
		Tx:  tx,
		Fee: *fee,
	}, nil
}

func (svc *service) SignTransaction(ctx context.Context, dto *SignRawTransactionDTO) (*SignedRawTransactionDTO, error) {
	//var utxos []map[string]interface{}
	//for _, s := range dto.Utxo {
	//	utxos = append(utxos, map[string]interface{}{"txid": s.TxId, "vout": s.Vout, "scriptPubKey": s.PKScript, "amount": s.Amount})
	//}

	tx, err := svc.btcRpcSvc.SignTransaction(ctx, dto.Tx, dto.PrivateKey, rpc_bitcoin.UTXO(dto.Utxo), dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedSignTx, err.Error())
	}

	return &SignedRawTransactionDTO{
		Hash: tx,
	}, nil
}

func (svc *service) SendTransaction(ctx context.Context, dto *SendRawTransactionDTO) (*SentRawTransactionDTO, error) {
	txId, err := svc.btcRpcSvc.SendTransaction(ctx, dto.SignedTx, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedSendTx, err.Error())
	}

	return &SentRawTransactionDTO{
		TxId: txId,
	}, nil
}

func (svc *service) WalletInfo(ctx context.Context, dto *WalletDTO) (*WalletInfoDTO, error) {
	info, err := svc.btcRpcSvc.WalletInfo(ctx, dto.WalletId, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedGetWalletInfo, err.Error())
	}

	return &WalletInfoDTO{
		Walletname:            info.Walletname,
		Walletversion:         info.Walletversion,
		Format:                info.Format,
		Balance:               info.Balance,
		UnconfirmedBalance:    info.UnconfirmedBalance,
		ImmatureBalance:       info.ImmatureBalance,
		Txcount:               info.Txcount,
		Keypoololdest:         info.Keypoololdest,
		Keypoolsize:           info.Keypoolsize,
		Hdseedid:              info.Hdseedid,
		KeypoolsizeHdInternal: info.KeypoolsizeHdInternal,
		Paytxfee:              info.Paytxfee,
		PrivateKeysEnabled:    info.PrivateKeysEnabled,
		AvoidReuse:            info.AvoidReuse,
		Scanning:              info.Scanning,
		Descriptors:           info.Descriptors,
	}, nil
}

func (svc *service) CreateWallet(ctx context.Context, dto *CreateWalletDTO) (*CreatedWalletInfoDTO, error) {
	walletId, err := svc.btcRpcSvc.CreateWallet(ctx, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedCreateWallet, err.Error())
	}

	return &CreatedWalletInfoDTO{WalletId: walletId}, nil
}

func (svc *service) LoadWaller(ctx context.Context, dto *LoadWalletDTO) (*LoadWalletInfoDTO, error) {
	err := svc.btcRpcSvc.LoadWallet(ctx, dto.WalletId, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedLoadWallet, err.Error())
	}

	return &LoadWalletInfoDTO{
		Message: "successful",
	}, nil
}

func (svc *service) ImportAddress(ctx context.Context, dto *ImportAddressDTO) (*ImportAddressInfoDTO, error) {
	err := svc.btcRpcSvc.ImportAddress(ctx, dto.Address, dto.WalletId, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedImportAddress, err.Error())
	}

	return &ImportAddressInfoDTO{
		Message: "successful",
	}, nil
}

func (svc *service) RescanWallet(ctx context.Context, dto *RescanWalletDTO) (*RescanWalletInfoDTO, error) {
	err := svc.btcRpcSvc.RescanWallet(ctx, dto.WalletId, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedRescanWallet, err.Error())
	}

	return &RescanWalletInfoDTO{
		Status:  "scanning has been started",
		Message: "if you want to check status of scan, you can use /wallet-info endpoint",
	}, nil
}

func (svc *service) ListUnspent(ctx context.Context, dto *ListUnspentDTO) (*ListUnspentInfoDTO, error) {
	list, err := svc.btcRpcSvc.ListUnspent(ctx, dto.Address, dto.WalletId, dto.Network)
	if err != nil {
		svc.log.WithContext(ctx).Errorf(err.Error())
		return nil, errors.WithMessage(ErrFailedGetUnspent, err.Error())
	}

	var result []*UnspentInfoDTO
	for _, unspent := range list {
		result = append(result, &UnspentInfoDTO{
			Txid:          unspent.Txid,
			Vout:          unspent.Vout,
			Address:       unspent.Address,
			Label:         unspent.Label,
			ScriptPubKey:  unspent.ScriptPubKey,
			Amount:        unspent.Amount,
			Confirmations: unspent.Confirmations,
			Spendable:     unspent.Spendable,
			Solvable:      unspent.Solvable,
			Safe:          unspent.Safe,
		})
	}

	return &ListUnspentInfoDTO{Result: result}, err
}
