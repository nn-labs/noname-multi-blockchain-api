package bitcoin

import (
	"errors"
)

type StatusNode struct {
	Chain                string      `json:"chain"`
	Blocks               interface{} `json:"blocks"`
	Headers              interface{} `json:"headers"`
	Verificationprogress interface{} `json:"verificationprogress"`
	Softforks            struct {
		Bip34 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip34"`
		Bip66 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip66"`
		Bip65 struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"bip65"`
		Csv struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"csv"`
		Segwit struct {
			Type   string      `json:"type"`
			Active bool        `json:"active"`
			Height interface{} `json:"height"`
		} `json:"segwit"`
		Taproot struct {
			Type string `json:"type"`
			Bip9 struct {
				Status              string      `json:"status"`
				StartTime           interface{} `json:"start_time"`
				Timeout             interface{} `json:"timeout"`
				Since               interface{} `json:"since"`
				MinActivationHeight int         `json:"min_activation_height"`
			} `json:"bip9"`

			Active bool `json:"active"`
		} `json:"taproot"`
	} `json:"softforks"`

	Warnings string `json:"warnings"`
}

func Status() (*StatusNode, error) {
	msg := struct {
		Result StatusNode `json:"result"`
		Error  struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	req := struct {
		JsonRPC string   `json:"json_rpc"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
	}{
		JsonRPC: "2.0",
		Method:  "getblockchaininfo",
		Params:  []string{},
	}

	err := Client(req, &msg)
	if err != nil {
		return nil, errors.New(msg.Error.Message)
	}

	if msg.Error.Message != "" {
		return nil, errors.New(msg.Error.Message)
	}

	return &msg.Result, nil
}
