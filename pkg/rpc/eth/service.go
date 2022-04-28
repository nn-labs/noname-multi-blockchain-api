package eth

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type Service interface {
	Status(ctx context.Context, network string) (*StatusNodeResponse, error)

	PendingNonceAt(ctx context.Context, account string, network string) (*uint64, error)
	SuggestGasPrice(ctx context.Context, network string) (*big.Int, error)
	EstimateGas(ctx context.Context, network string) (*uint64, error)

	GetNetworkId(ctx context.Context, network string) (*big.Int, error)
	GetTransactionByHash(ctx context.Context, tx string, network string) (*TransactionByHashResponse, error)

	CreateTransaction(ctx context.Context, fromAddress, toAddress string, amount int64, network string) (*string, error)
	SignTransaction(ctx context.Context, tx, privateKey string, network string) (*string, error)
	SendTransaction(ctx context.Context, signedTx, network string) (*string, error)
}

type service struct {
	ethClient Client
}

func NewService(ethClient Client) (Service, error) {
	if ethClient == nil {
		return nil, errors.New("invalid ethereum client")
	}

	return &service{ethClient: ethClient}, nil
}

func (s *service) Status(ctx context.Context, network string) (*StatusNodeResponse, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_syncing",
		Params:  []string{},
	}

	msg := struct {
		JsonRpc string             `json:"jsonrpc"`
		Id      string             `json:"id"`
		Result  StatusNodeResponse `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	return &msg.Result, nil
}

func (s *service) PendingNonceAt(ctx context.Context, account string, network string) (*uint64, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionCount",
		Params:  []string{account, "pending"},
	}

	msg := struct {
		JsonRpc string         `json:"jsonrpc"`
		Id      string         `json:"id"`
		Result  hexutil.Uint64 `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	return (*uint64)(&msg.Result), nil
}

func (s *service) SuggestGasPrice(ctx context.Context, network string) (*big.Int, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_gasPrice",
		Params:  []string{},
	}

	msg := struct {
		JsonRpc string      `json:"jsonrpc"`
		Id      string      `json:"id"`
		Result  hexutil.Big `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	return (*big.Int)(&msg.Result), nil
}

func (s *service) EstimateGas(ctx context.Context, network string) (*uint64, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_estimateGas",
		Params:  []string{},
	}

	msg := struct {
		JsonRpc string         `json:"jsonrpc"`
		Id      string         `json:"id"`
		Result  hexutil.Uint64 `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	return (*uint64)(&msg.Result), nil
}

func (s *service) GetNetworkId(ctx context.Context, network string) (*big.Int, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "net_version",
		Params:  []string{},
	}

	msg := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	version := new(big.Int)

	if _, ok := version.SetString(msg.Result, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", msg.Result)
	}

	return version, nil
}

func (s *service) GetTransactionByHash(ctx context.Context, tx string, network string) (*TransactionByHashResponse, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionByHash",
		Params:  []string{tx},
	}

	msg := struct {
		JsonRpc string                    `json:"jsonrpc"`
		Id      string                    `json:"id"`
		Result  TransactionByHashResponse `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	return &msg.Result, nil
}

func (s *service) CreateTransaction(ctx context.Context, fromAddress, toAddress string, amount int64, network string) (*string, error) {
	nonce, err := s.PendingNonceAt(ctx, fromAddress, network)
	if err != nil {
		return nil, err
	}

	value := big.NewInt(ToWei(amount, 18).Int64()) // in wei (1 eth)

	gasLimit, err := s.EstimateGas(ctx, network)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.SuggestGasPrice(ctx, network)
	if err != nil {
		return nil, err
	}

	toEthAddress := common.HexToAddress(toAddress)
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    *nonce,
		GasPrice: gasPrice,
		Gas:      *gasLimit,
		To:       &toEthAddress,
		Value:    value,
		Data:     data,
	})

	txHash := tx.Hash().String()
	return &txHash, nil
}

func (s *service) SignTransaction(ctx context.Context, tx, privateKey string, network string) (*string, error) {
	chainID, err := s.GetNetworkId(ctx, network)
	if err != nil {
		return nil, err
	}

	privateEthKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	newTx := new(types.Transaction)
	rawTxBytes, err := hex.DecodeString(tx)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(newTx, types.NewLondonSigner(chainID), privateEthKey)
	if err != nil {
		return nil, err
	}

	signedTxHash := signedTx.Hash().String()
	return &signedTxHash, nil
}

func (s *service) SendTransaction(ctx context.Context, signedTx, network string) (*string, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_sendRawTransaction",
		Params:  []string{signedTx},
	}

	msg := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	return &msg.Result, nil
}
