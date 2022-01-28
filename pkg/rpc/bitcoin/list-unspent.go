package bitcoin

import (
	"encoding/json"
	"errors"
)

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

func ListUnspent(client IBtcClient, address, network string) ([]*Unspent, error) {
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

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	response, err := client.Send(body, true, network)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return msg.Result, nil
}
