package bitcoin

import (
	"encoding/json"
	"errors"
)

func ImportAddress(client IBtcClient, address, walletId, network string) error {
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

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return errors.New(err.Error())
	}

	response, err := client.Send(body, walletId, network)
	if err != nil {
		return errors.New(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return err
	}

	if msg.Error.Message != "" {
		return errors.New(msg.Error.Message)
	}

	return nil
}
