package bitcoin

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

func CreateWallet(client IBtcClient, network string) (string, error) {
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

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	response, err := client.Send(body, true, network)
	if err != nil {
		return "", errors.New(err.Error())
	}

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", err
	}

	if msg.Result.Warning != "" {
		return "", errors.New(msg.Result.Warning)
	}

	if msg.Error.Message != "" {
		return "", errors.New(msg.Error.Message)
	}

	return msg.Result.Name, nil
}
