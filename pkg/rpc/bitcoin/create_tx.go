package bitcoin

import (
	"encoding/json"
	"errors"
)

func CreateTransaction(client IBtcClient, inputs []map[string]interface{}, outputs []map[string]string, network string) (string, error) {
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "createrawtransaction",
		Params:  []interface{}{inputs, outputs},
	}

	msg := struct {
		Result string `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	response, err := client.Send(body, false, network)
	if err != nil {
		return "", errors.New(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Error.Message != "" {
		return "", errors.New(msg.Error.Message)
	}

	return msg.Result, nil
}
