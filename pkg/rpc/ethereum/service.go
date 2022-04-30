package ethereum_rpc

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
	"github.com/google/uuid"
	"io/ioutil"
	"math/big"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go
type Service interface {
	Status(ctx context.Context, network string) (*StatusNodeResponse, error)

	PendingNonceAt(ctx context.Context, account string, network string) (*string, error)
	SuggestGasPrice(ctx context.Context, network string) (*string, error)
	EstimateGas(ctx context.Context, fromAddress, toAddress, data string, value, gasPrice *big.Int, network string) (*string, error)

	GetNetworkId(ctx context.Context, network string) (*big.Int, error)
	GetTransactionByHash(ctx context.Context, tx string, network string) (*TransactionByHashResponse, error)

	CreateTransaction(ctx context.Context, fromAddress, toAddress string, amount float64, network string) (*string, *float64, error)
	SignTransaction(ctx context.Context, tx, privateKey string, network string) (*string, error)
	SendTransaction(ctx context.Context, signedTx, network string) (*string, error)
}

type service struct {
	ethClient Client
}

func NewService(ethClient Client) (Service, error) {
	if ethClient == nil {
		return nil, errors.New("invalid ethereum rpc client")
	}

	return &service{ethClient: ethClient}, nil
}

func (s *service) Status(ctx context.Context, network string) (*StatusNodeResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_syncing",
		Params:  []interface{}{},
		Id:      id.String(),
	}

	msg := struct {
		JsonRpc string      `json:"jsonrpc"`
		Id      string      `json:"id"`
		Result  interface{} `json:"result"`
		Error   struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	var result *StatusNodeResponse
	switch msg.Result.(type) {
	case bool:
		result = &StatusNodeResponse{
			SyncMessage: "node has synced",
		}
	default:
		byteData, err := json.Marshal(msg.Result)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(byteData, &result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *service) PendingNonceAt(ctx context.Context, account string, network string) (*string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionCount",
		Params:  []interface{}{account, "pending"},
		Id:      id.String(),
	}

	msg := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
		Error   struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}

func (s *service) SuggestGasPrice(ctx context.Context, network string) (*string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_gasPrice",
		Params:  []interface{}{},
		Id:      id.String(),
	}

	msg := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
		Error   struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}

func (s *service) EstimateGas(ctx context.Context, fromAddress, toAddress, data string, value, gasPrice *big.Int, network string) (*string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	arg := map[string]interface{}{"from": fromAddress, "to": toAddress}
	if len(data) > 0 {
		arg["data"] = hexutil.Bytes(data)
	}
	if value != nil {
		arg["value"] = (*hexutil.Big)(value)
	}
	if gasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(gasPrice)
	}

	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_estimateGas",
		Params:  []interface{}{arg},
		Id:      id.String(),
	}

	msg := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
		Error   struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}

func (s *service) GetNetworkId(ctx context.Context, network string) (*big.Int, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "net_version",
		Params:  []interface{}{},
		Id:      id.String(),
	}

	msg := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
		Error   struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	version := new(big.Int)

	if _, ok := version.SetString(msg.Result, 10); !ok {
		return nil, errors.New(fmt.Sprintf("invalid net_version result %q", msg.Result))
	}

	return version, nil
}

func (s *service) GetTransactionByHash(ctx context.Context, tx string, network string) (*TransactionByHashResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_getTransactionByHash",
		Params:  []interface{}{tx},
		Id:      id.String(),
	}

	msg := struct {
		JsonRpc string                    `json:"jsonrpc"`
		Id      string                    `json:"id"`
		Result  TransactionByHashResponse `json:"result"`
		Error   struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}

func (s *service) CreateTransaction(ctx context.Context, fromAddress, toAddress string, amount float64, network string) (*string, *float64, error) {
	nonce, err := s.PendingNonceAt(ctx, fromAddress, network)
	if err != nil {
		return nil, nil, err
	}

	decodeNonce, err := hexutil.DecodeUint64(*nonce)
	if err != nil {
		return nil, nil, err
	}

	value := big.NewInt(ToWei(amount, 18).Int64()) // in wei (1 ethereum)

	gasPrice, err := s.SuggestGasPrice(ctx, network)
	if err != nil {
		return nil, nil, err
	}

	decodeGasPrice, err := hexutil.DecodeBig(*gasPrice)
	if err != nil {
		return nil, nil, err
	}

	gas, err := s.EstimateGas(ctx, fromAddress, toAddress, "", value, decodeGasPrice, network)
	if err != nil {
		return nil, nil, err
	}

	decodeGas, err := hexutil.DecodeUint64(*gas)
	if err != nil {
		return nil, nil, err
	}

	toEthAddress := common.HexToAddress(toAddress)
	var data []byte

	fee := new(big.Int)
	fee.SetUint64(decodeGasPrice.Uint64() * decodeGas)
	floatFee := ToDecimal(fee, 18).InexactFloat64()

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    decodeNonce,
		GasPrice: decodeGasPrice,
		Gas:      decodeGas,
		To:       &toEthAddress,
		Value:    value,
		Data:     data,
	})

	txBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return nil, nil, err
	}

	rawTxHex := hex.EncodeToString(txBytes)

	return &rawTxHex, &floatFee, nil
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
	if err != nil {
		return nil, err
	}

	err = rlp.DecodeBytes(rawTxBytes, &newTx)
	if err != nil {
		return nil, err
	}

	signedTx, err := types.SignTx(newTx, types.NewLondonSigner(chainID), privateEthKey)
	if err != nil {
		return nil, err
	}

	signedTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, err
	}

	signedTxHex := hex.EncodeToString(signedTxBytes)

	return &signedTxHex, nil
}

func (s *service) SendTransaction(ctx context.Context, signedTx, network string) (*string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_sendRawTransaction",
		Params:  []interface{}{"0x" + signedTx},
		Id:      id.String(),
	}

	msg := struct {
		JsonRpc string `json:"jsonrpc"`
		Id      string `json:"id"`
		Result  string `json:"result"`
		Error   struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
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

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}
