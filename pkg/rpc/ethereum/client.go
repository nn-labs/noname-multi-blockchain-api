package ethereum_rpc

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
	Send(ctx context.Context, body io.Reader, network string) (*http.Response, error)
	EncodeBaseRequest(request interface{}) (*bytes.Buffer, error)
}

type client struct {
	ethRpcEndpointTestNet string
	ethRpcEndpointMainNet string
}

func NewClient(ethRpcEndpointTestNet string, ethRpcEndpointMainNet string) (Client, error) {
	if ethRpcEndpointTestNet == "" {
		return nil, errors.New("invalid ethereum rpc testnet endpoint")
	}
	if ethRpcEndpointMainNet == "" {
		return nil, errors.New("invalid ethereum rpc mainnet endpoint")
	}

	return &client{
		ethRpcEndpointTestNet: ethRpcEndpointTestNet,
		ethRpcEndpointMainNet: ethRpcEndpointMainNet,
	}, nil
}

func (c *client) Send(ctx context.Context, body io.Reader, network string) (*http.Response, error) {
	var endPoint string

	if network == "main" {
		endPoint = c.ethRpcEndpointMainNet
	} else {
		endPoint = c.ethRpcEndpointTestNet
	}

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "POST", endPoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	//defer resp.Body.Close()

	return resp, nil
}

func (c *client) EncodeBaseRequest(request interface{}) (*bytes.Buffer, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(data)

	return reqBody, nil
}
