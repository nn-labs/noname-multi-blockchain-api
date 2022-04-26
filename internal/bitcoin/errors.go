package bitcoin

import (
	"nn-blockchain-api/pkg/codes"
	"nn-blockchain-api/pkg/errors"
)

const (
	StatusInvalidRequest      errors.Status = "invalid_request"
	StatusFailedGetStatusNode errors.Status = "failed_get_status_node"
	StatusFailedCreateTx      errors.Status = "failed_create_tx"
	StatusFailedDecodeTx      errors.Status = "failed_decode_tx"
	StatusFailedFundForTx     errors.Status = "failed_fund_for_tx"
	StatusFailedSignTx        errors.Status = "failed_sign_tx"
	StatusFailedSendTx        errors.Status = "failed_send_tx"
	StatusFailedGetWalletInfo errors.Status = "failed_get_wallet_info"
	StatusFailedCreateWallet  errors.Status = "failed_create_wallet"
	StatusFailedLoadWallet    errors.Status = "failed_load_wallet"
	StatusFailedImportAddress errors.Status = "failed_import_address"
	StatusFailedRescanWallet  errors.Status = "failed_rescan_wallet"
	StatusFailedGetUnspent    errors.Status = "failed_get_unspent"
)

var (
	ErrInvalidRequest      = errors.New(codes.BadRequest, StatusInvalidRequest)
	ErrFailedGetStatusNode = errors.New(codes.InternalError, StatusFailedGetStatusNode)
	ErrFailedCreateTx      = errors.New(codes.InternalError, StatusFailedCreateTx)
	ErrFailedDecodeTx      = errors.New(codes.InternalError, StatusFailedDecodeTx)
	ErrFailedFundForTx     = errors.New(codes.InternalError, StatusFailedFundForTx)
	ErrFailedSignTx        = errors.New(codes.InternalError, StatusFailedSignTx)
	ErrFailedSendTx        = errors.New(codes.InternalError, StatusFailedSendTx)
	ErrFailedGetWalletInfo = errors.New(codes.InternalError, StatusFailedGetWalletInfo)
	ErrFailedCreateWallet  = errors.New(codes.InternalError, StatusFailedCreateWallet)
	ErrFailedLoadWallet    = errors.New(codes.InternalError, StatusFailedLoadWallet)
	ErrFailedImportAddress = errors.New(codes.InternalError, StatusFailedImportAddress)
	ErrFailedRescanWallet  = errors.New(codes.InternalError, StatusFailedRescanWallet)
	ErrFailedGetUnspent    = errors.New(codes.InternalError, StatusFailedGetUnspent)
)
