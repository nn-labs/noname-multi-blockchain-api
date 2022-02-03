package bitcoin

import (
	"encoding/json"
	"errors"
	"math/big"
)

func GetCurrentFee(client IBtcClient, network string) (*float64, error) {
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "estimatesmartfee",
		Params:  []interface{}{2},
	}

	msg := struct {
		Result struct {
			Feerate float64 `json:"feerate"`
			Blocks  int64   `json:"blocks"`
		}
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := client.EncodeBaseRequest(req)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	response, err := client.Send(body, "", network)
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

	var fee float64
	fee = msg.Result.Feerate
	// sanity check
	if fee > 0.05 {
		fee = 0.1
	} else if fee < 0 {
		fee = 0
	}
	//fmt.Printf("fee: %f\n", fee)

	if fee == 0 {
		return &fee, errors.New("could not get fees")
	}

	return &fee, nil
}

func GetCurrentFeeRate(client IBtcClient, network string) (*big.Int, error) {
	fee, err := GetCurrentFee(client, network)
	if err != nil {
		return nil, err
	}

	// convert to satoshis to bytes
	// feeRate := big.NewInt(int64(msg.Result * 1.0E8))
	// convert to satoshis to kb
	feeRate := big.NewInt(int64(*fee * 1.0e5))

	//fmt.Printf("fee rate: %s\n", feeRate)

	return feeRate, nil
}
