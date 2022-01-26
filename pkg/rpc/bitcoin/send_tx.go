package bitcoin

import (
	"encoding/json"
	"errors"
)

func SendTx(client IBtcClient, signedTx string) (string, error) {
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

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	response, err := client.Send(body)
	if err != nil {
		return "", errors.New(err.Error())
	}

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Error.Message != "" {
		return "", errors.New(msg.Error.Message)
	}

	return msg.Result, nil
}
