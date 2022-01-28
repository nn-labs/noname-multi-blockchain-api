package bitcoin

import (
	"encoding/json"
	"errors"
)

func FundForRawTransaction(client IBtcClient, createdTx, changeAddress, network string) (string, *float64, error) {
	subtractFeeFromOutputs := []int64{0}

	params := map[string]interface{}{
		"changeAddress":          changeAddress,
		"subtractFeeFromOutputs": subtractFeeFromOutputs,
	}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "fundrawtransaction",
		Params:  []interface{}{createdTx, params},
	}

	msg := struct {
		Result struct {
			Hex string  `json:"hex"`
			Fee float64 `json:"fee"`
		} `json:"result"`
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return "", nil, errors.New(err.Error())
	}

	response, err := client.Send(body, "", network)
	if err != nil {
		return "", nil, errors.New(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return "", nil, err
	}

	if msg.Error.Message != "" {
		return "", nil, errors.New(msg.Error.Message)
	}

	return msg.Result.Hex, &msg.Result.Fee, nil
}
