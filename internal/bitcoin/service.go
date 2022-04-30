package bitcoin

import (
	"context"
	"github.com/sirupsen/logrus"
	"nn-blockchain-api/pkg/errors"
	bitcoin_rpc "nn-blockchain-api/pkg/rpc/bitcoin"
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
	btcRpcSvc bitcoin_rpc.Service
	log       *logrus.Logger
}

func NewService(btcRpcSvc bitcoin_rpc.Service, log *logrus.Logger) (Service, error) {
	if btcRpcSvc == nil {
		return nil, errors.NewInternal("invalid btc rpc service")
	}
	if log == nil {
		return nil, errors.NewInternal("invalid logger")
	}
	return &service{btcRpcSvc: btcRpcSvc, log: log}, nil
}

func (s *service) StatusNode(ctx context.Context, dto *StatusNodeDTO) (*StatusNodeInfoDTO, error) {
	status, err := s.btcRpcSvc.Status(ctx, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed check node status: %v", err)
		return nil, errors.WithMessage(ErrFailedGetStatusNode, err.Error())
		//return nil, ErrFailedGetStatusNode
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

func (s *service) CreateTransaction(ctx context.Context, dto *CreateRawTransactionDTO) (*CreatedRawTransactionDTO, error) {
	tx, fee, err := s.btcRpcSvc.CreateTransaction(ctx, bitcoin_rpc.UTXO(dto.Utxo), dto.FromAddress, dto.ToAddress, dto.Amount, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed create transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedCreateTx, err.Error())
		//return nil, ErrFailedCreateTx
	}

	return &CreatedRawTransactionDTO{
		Tx:  *tx,
		Fee: *fee,
	}, nil
}

func (s *service) DecodeTransaction(ctx context.Context, dto *DecodeRawTransactionDTO) (*DecodedRawTransactionDTO, error) {
	decodedTx, err := s.btcRpcSvc.DecodeTransaction(ctx, dto.Tx, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed decode transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedDecodeTx, err.Error())
		//return nil, ErrFailedDecodeTx
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

func (s *service) FoundForRawTransaction(ctx context.Context, dto *FundForRawTransactionDTO) (*FundedRawTransactionDTO, error) {
	tx, fee, err := s.btcRpcSvc.FundForTransaction(ctx, dto.CreatedTxHex, dto.ChangeAddress, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed found for transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedFundForTx, err.Error())
		//return nil, ErrFailedFundForTx
	}

	return &FundedRawTransactionDTO{
		Tx:  tx,
		Fee: *fee,
	}, nil
}

func (s *service) SignTransaction(ctx context.Context, dto *SignRawTransactionDTO) (*SignedRawTransactionDTO, error) {
	//var utxos []map[string]interface{}
	//for _, s := range dto.Utxo {
	//	utxos = append(utxos, map[string]interface{}{"txid": s.TxId, "vout": s.Vout, "scriptPubKey": s.PKScript, "amount": s.Amount})
	//}

	tx, err := s.btcRpcSvc.SignTransaction(ctx, dto.Tx, dto.PrivateKey, bitcoin_rpc.UTXO(dto.Utxo), dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed sign transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedSignTx, err.Error())
		//return nil, ErrFailedSignTx
	}

	return &SignedRawTransactionDTO{
		Hash: tx,
	}, nil
}

func (s *service) SendTransaction(ctx context.Context, dto *SendRawTransactionDTO) (*SentRawTransactionDTO, error) {
	txId, err := s.btcRpcSvc.SendTransaction(ctx, dto.SignedTx, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed send transaction: %v", err)
		return nil, errors.WithMessage(ErrFailedSendTx, err.Error())
		//return nil, ErrFailedSendTx
	}

	return &SentRawTransactionDTO{
		TxId: txId,
	}, nil
}

func (s *service) WalletInfo(ctx context.Context, dto *WalletDTO) (*WalletInfoDTO, error) {
	info, err := s.btcRpcSvc.WalletInfo(ctx, dto.WalletId, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed get wallet info: %v", err)
		return nil, errors.WithMessage(ErrFailedGetWalletInfo, err.Error())
		//return nil, ErrFailedGetWalletInfo
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

func (s *service) CreateWallet(ctx context.Context, dto *CreateWalletDTO) (*CreatedWalletInfoDTO, error) {
	walletId, err := s.btcRpcSvc.CreateWallet(ctx, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed create wallet: %v", err)
		return nil, errors.WithMessage(ErrFailedCreateWallet, err.Error())
		//return nil, ErrFailedCreateWallet
	}

	return &CreatedWalletInfoDTO{WalletId: walletId}, nil
}

func (s *service) LoadWaller(ctx context.Context, dto *LoadWalletDTO) (*LoadWalletInfoDTO, error) {
	err := s.btcRpcSvc.LoadWallet(ctx, dto.WalletId, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed load wallet: %v", err)
		return nil, errors.WithMessage(ErrFailedLoadWallet, err.Error())
		//return nil, ErrFailedLoadWallet
	}

	return &LoadWalletInfoDTO{
		Message: "successful",
	}, nil
}

func (s *service) ImportAddress(ctx context.Context, dto *ImportAddressDTO) (*ImportAddressInfoDTO, error) {
	err := s.btcRpcSvc.ImportAddress(ctx, dto.Address, dto.WalletId, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed import wallet: %v", err)
		return nil, errors.WithMessage(ErrFailedImportAddress, err.Error())
		//return nil, ErrFailedImportAddress
	}

	return &ImportAddressInfoDTO{
		Message: "successful",
	}, nil
}

func (s *service) RescanWallet(ctx context.Context, dto *RescanWalletDTO) (*RescanWalletInfoDTO, error) {
	err := s.btcRpcSvc.RescanWallet(ctx, dto.WalletId, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed rescan wallet: %v", err)
		return nil, errors.WithMessage(ErrFailedRescanWallet, err.Error())
		//return nil, ErrFailedRescanWallet
	}

	return &RescanWalletInfoDTO{
		Status:  "scanning has been started",
		Message: "if you want to check status of scan, you can use /wallet-info endpoint",
	}, nil
}

func (s *service) ListUnspent(ctx context.Context, dto *ListUnspentDTO) (*ListUnspentInfoDTO, error) {
	list, err := s.btcRpcSvc.ListUnspent(ctx, dto.Address, dto.WalletId, dto.Network)
	if err != nil {
		s.log.WithContext(ctx).Errorf("failed get unspend list: %v", err)
		return nil, errors.WithMessage(ErrFailedGetUnspent, err.Error())
		//return nil, ErrFailedGetUnspent
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
