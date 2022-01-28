package bitcoin

import (
	"encoding/json"
	"errors"
)

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

func WalletInfo(client IBtcClient, walletId, network string) (*Info, error) {
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

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	response, err := client.Send(body, walletId, network)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}
