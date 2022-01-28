package bitcoin

import (
	"encoding/json"
	"errors"
)

func SignTx(client IBtcClient, tx, privateKey string, utxos []map[string]interface{}, network string) (string, error) {
	privateKeyArray := []string{privateKey}
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "signrawtransactionwithkey",
		Params:  []interface{}{tx, privateKeyArray, utxos},
	}

	msg := struct {
		Result struct {
			Hex      string `json:"hex"`
			Complete bool   `json:"complete"`
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

	response, err := client.Send(body, false, network)
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

	if !msg.Result.Complete {
		return "", errors.New("signing transaction not complete. Please try again")
	}

	return msg.Result.Hex, nil
}
