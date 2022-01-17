package bitcoin

import (
	"errors"
)

func StatusNode() (*StatusNodeDTO, error) {
	msg := struct {
		Result StatusNodeDTO `json:"result"`
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
