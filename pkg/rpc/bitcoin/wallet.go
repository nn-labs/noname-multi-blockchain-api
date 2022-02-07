package bitcoin

import (
	"encoding/json"
	"github.com/google/uuid"
	"nn-blockchain-api/pkg/errors"
	"time"
)

//go:generate mockgen -source=wallet.go -destination=mocks/wallet_mock.go

type Info struct {
	Walletname            string      `json:"walletname"`
	Walletversion         int         `json:"walletversion"`
	Format                string      `json:"format"`
	Balance               float64     `json:"balance"`
	UnconfirmedBalance    float64     `json:"unconfirmed_balance"`
	ImmatureBalance       float64     `json:"immature_balance"`
	Txcount               int         `json:"txcount"`
	Keypoololdest         int         `json:"keypoololdest"`
	Keypoolsize           int         `json:"keypoolsize"`
	Hdseedid              string      `json:"hdseedid"`
	KeypoolsizeHdInternal int         `json:"keypoolsize_hd_internal"`
	Paytxfee              float64     `json:"paytxfee"`
	PrivateKeysEnabled    bool        `json:"private_keys_enabled"`
	AvoidReuse            bool        `json:"avoid_reuse"`
	Scanning              interface{} `json:"scanning"`
	Descriptors           bool        `json:"descriptors"`
}

type Unspent struct {
	Txid          string  `json:"txid"`
	Vout          int     `json:"vout"`
	Address       string  `json:"address"`
	Label         string  `json:"label"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Confirmations int     `json:"confirmations"`
	Spendable     bool    `json:"spendable"`
	Solvable      bool    `json:"solvable"`
	Safe          bool    `json:"safe"`
}

type WalletService interface {
	WalletInfo(walletId, network string) (*Info, error)
	CreateWallet(network string) (string, error)
	LoadWallet(walletId, network string) error
	ImportAddress(address, walletId, network string) error
	RescanWallet(walletId, network string) error
	ListUnspent(address, walletId, network string) ([]*Unspent, error)
}

type walletService struct {
	btcClient IBtcClient
}

func NewWalletService(btcClient IBtcClient) (WalletService, error) {
	if btcClient == nil {
		return nil, errors.NewInternal("invalid btc client")
	}
	return &walletService{btcClient: btcClient}, nil
}

func (svc *walletService) WalletInfo(walletId, network string) (*Info, error) {
	msg := struct {
		Result Info `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "getwalletinfo",
		Params:  []interface{}{},
	}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := svc.btcClient.Send(body, walletId, network)
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

	return &msg.Result, nil
}

func (svc *walletService) CreateWallet(network string) (string, error) {
	walletId, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "createwallet",
		Params:  []interface{}{walletId},
	}

	msg := struct {
		Result struct {
			Name    string `json:"name"`
			Warning string `json:"warning"`
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

	if msg.Result.Warning != "" {
		return "", errors.NewInternal(msg.Result.Warning)
	}

	if msg.Error.Message != "" {
		return "", errors.NewInternal(msg.Error.Message)
	}

	return msg.Result.Name, nil
}

func (svc *walletService) LoadWallet(walletId, network string) error {
	msg := struct {
		Result struct {
			Name    string `json:"name"`
			Warning string `json:"warning"`
		} `json:"result"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "loadwallet",
		Params:  []interface{}{walletId},
	}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return err
	}

	response, err := svc.btcClient.Send(body, "", network)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return err
	}

	if msg.Result.Warning != "" {
		return errors.NewInternal(msg.Result.Warning)
	}

	if msg.Error.Message != "" {
		return errors.NewInternal(msg.Error.Message)
	}

	return nil
}

func (svc *walletService) ImportAddress(address, walletId, network string) error {
	msg := struct {
		Result string `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "importaddress",
		Params:  []interface{}{address, "", false},
	}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return err
	}

	response, err := svc.btcClient.Send(body, walletId, network)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return err
	}

	if msg.Error.Message != "" {
		return errors.NewInternal(msg.Error.Message)
	}

	return nil
}

func (svc *walletService) RescanWallet(walletId, network string) error {
	//msg := struct {
	//	Result struct {
	//		StartHeight int64 `json:"start_height"`
	//		StopHeight  int64 `json:"stop_height"`
	//	} `json:"result"`
	//	Error struct {
	//		Message string `json:"message"`
	//	} `json:"error"`
	//}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "rescanblockchain",
		Params:  []interface{}{},
	}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return err
	}

	errs := make(chan error, 1)
	go func() {
		response, err := svc.btcClient.Send(body, walletId, network)
		if err != nil {
			errs <- err
		}

		//err = json.NewDecoder(response.Body).Decode(&msg)
		//if err != nil {
		//	errs <- errors.New(err.Error())
		//}
		//
		//if msg.Error.Message != "" {
		//	errs <- errors.New(msg.Error.Message)
		//}

		defer response.Body.Close()
		close(errs)
	}()

	select {
	case <-time.After(10 * time.Second):
		return nil
	case err := <-errs:
		if err != nil {
			return errors.NewInternal(err.Error())
		}
	}

	return nil
}

func (svc *walletService) ListUnspent(address, walletId, network string) ([]*Unspent, error) {
	msg := struct {
		Result []*Unspent `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "listunspent",
		Params:  []interface{}{1, 99999999, []string{address}},
	}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := svc.btcClient.Send(body, walletId, network)
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

	return msg.Result, nil
}
