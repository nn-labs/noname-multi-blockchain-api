package bitcoin_rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

//go:generate mockgen -source=client.go -destination=mocks/client_mock.go
type Client interface {
	Send(ctx context.Context, body io.Reader, walletId string, network string) (*http.Response, error)
	EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error)
	//DecodeBaseResponse(response *http.Response, msg interface{}) (*BaseResponse, error)
}

type client struct {
	btcRpcEndpointTestNet string
	btcRpcEndpointMainNet string
	btcUser               string
	btcPassword           string
}

func NewClient(btcRpcEndpointTestNet, btcRpcEndpointMainNet, btcUser, btcPassword string) (Client, error) {
	if btcRpcEndpointTestNet == "" {
		return nil, errors.New("invalid bitcoin rpc testnet endpoint")
	}
	if btcRpcEndpointMainNet == "" {
		return nil, errors.New("invalid bitcoin rpc mainnet endpoint")
	}
	if btcUser == "" {
		return nil, errors.New("invalid bitcoin rpc user")
	}
	if btcPassword == "" {
		return nil, errors.New("invalid bitcoin rpc password")
	}

	return &client{
		btcRpcEndpointTestNet: btcRpcEndpointTestNet,
		btcRpcEndpointMainNet: btcRpcEndpointMainNet,
		btcUser:               btcUser,
		btcPassword:           btcPassword,
	}, nil
}

func (c *client) Send(ctx context.Context, body io.Reader, walletId string, network string) (*http.Response, error) {
	var endPoint string

	if network == "main" {
		endPoint = c.btcRpcEndpointMainNet
	} else {
		endPoint = c.btcRpcEndpointTestNet
	}

	if walletId != "" {
		endPoint = endPoint + "/wallet/" + walletId
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", endPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(c.btcUser, c.btcPassword)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//defer resp.Body.Close()

	return resp, nil
}

func (c *client) EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(data)

	return reqBody, nil
}

//func (btc *btcClient) DecodeBaseResponse(response *http.Response, msg interface{}) (*BaseResponse, error) {
//	err := json.NewDecoder(response.Body).Decode(&response)
//	if err != nil {
//		return nil, err
//	}
//
//	return &msg, nil
//}
