package rpc_bitcoin

import (
	"context"
	"encoding/json"
	"nn-blockchain-api/pkg/errors"
)

//go:generate mockgen -source=health.go -destination=mocks/health_mock.go

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

type HealthService interface {
	Status(ctx context.Context, network string) (*StatusNode, error)
}

type healthService struct {
	btcClient IBtcClient
}

func NewHealthService(btcClient IBtcClient) (HealthService, error) {
	if btcClient == nil {
		return nil, errors.NewInternal("invalid btc client")
	}
	return &healthService{btcClient: btcClient}, nil
}

func (svc *healthService) Status(ctx context.Context, network string) (*StatusNode, error) {
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "getblockchaininfo",
		Params:  []interface{}{},
	}

	msg := struct {
		Result StatusNode `json:"result"`
		Error  struct {
			Code    int64  `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}{}

	body, err := svc.btcClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := svc.btcClient.Send(ctx, body, "", network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	if msg.Error.Message != "" {
		return nil, errors.NewInternal(msg.Error.Message)
	}

	return &msg.Result, nil
}
