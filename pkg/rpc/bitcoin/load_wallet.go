package bitcoin

import (
	"encoding/json"
	"errors"
)

func LoadWallet(client IBtcClient, network string) error {
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
		Params:  []interface{}{"development"},
	}

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return errors.New(err.Error())
	}

	response, err := client.Send(body, false, network)
	if err != nil {
		return errors.New(err.Error())
	}

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return err
	}

	if msg.Result.Warning != "" {
		return errors.New(msg.Result.Warning)
	}

	if msg.Error.Message != "" {
		return errors.New(msg.Error.Message)
	}

	return nil
}
