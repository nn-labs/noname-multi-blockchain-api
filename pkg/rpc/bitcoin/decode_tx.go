package bitcoin

import (
	"encoding/json"
	"errors"
)

type DecodedTx struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`

		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm     string `json:"asm"`
			Hex     string `json:"hex"`
			Address string `json:"address"`
			Type    string `json:"type"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
}

func DecodeTx(client IBtcClient, tx string, network string) (*DecodedTx, error) {
	msg := struct {
		Result DecodedTx `json:"result"`
		Error  struct {
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "decoderawtransaction",
		Params:  []interface{}{tx},
	}

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

	return &DecodedTx{
		Txid:     msg.Result.Txid,
		Hash:     msg.Result.Hash,
		Version:  msg.Result.Version,
		Size:     msg.Result.Size,
		Vsize:    msg.Result.Vsize,
		Weight:   msg.Result.Weight,
		Locktime: msg.Result.Locktime,
		Vin:      msg.Result.Vin,
		Vout:     msg.Result.Vout,
	}, nil
}
