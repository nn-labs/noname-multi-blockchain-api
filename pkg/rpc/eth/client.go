package eth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Client interface {
	Send(ctx context.Context, body io.Reader, network string) (*http.Response, error)
	EncodeBaseRequest(request interface{}) (*bytes.Buffer, error)
	//DecodeBaseResponseWithIntResult(response *http.Response) (*BaseResponseWithIntResult, error)
	//DecodeBaseResponseWithStringResult(response *http.Response) (*BaseResponseWithStringResult, error)
	//DecodeBaseResponseWithBoolResult(response *http.Response) (*BaseResponseWithBoolResult, error)
	//DecodeBaseResponseWithArrayResult(response *http.Response) (*BaseResponseWithArrayResult, error)
}

type client struct {
	ethRpcEndpointTestNet string
	ethRpcEndpointMainNet string
}

func NewClient(ethRpcEndpointTestNet string, ethRpcEndpointMainNet string) (Client, error) {
	if ethRpcEndpointTestNet == "" {
		return nil, errors.New("invalid ethereum testnet endpoint")
	}
	if ethRpcEndpointMainNet == "" {
		return nil, errors.New("invalid ethereum mainnet endpoint")
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

//func (c *client) DecodeBaseResponseWithIntResult(response *http.Response) (*BaseResponseWithIntResult, error) {
//	var baseResponse BaseResponseWithIntResult
//	err := json.NewDecoder(response.Body).Decode(&baseResponse)
//	if err != nil {
//		return nil, err
//	}
//
//	return &baseResponse, nil
//}
//
//func (c *client) DecodeBaseResponseWithStringResult(response *http.Response) (*BaseResponseWithStringResult, error) {
//	var baseResponse BaseResponseWithStringResult
//	err := json.NewDecoder(response.Body).Decode(&baseResponse)
//	if err != nil {
//		return nil, err
//	}
//
//	return &baseResponse, nil
//}
//
//func (c *client) DecodeBaseResponseWithBoolResult(response *http.Response) (*BaseResponseWithBoolResult, error) {
//	var baseResponse BaseResponseWithBoolResult
//	err := json.NewDecoder(response.Body).Decode(&baseResponse)
//	if err != nil {
//		return nil, err
//	}
//
//	return &baseResponse, nil
//}
//
//func (c *client) DecodeBaseResponseWithArrayResult(response *http.Response) (*BaseResponseWithArrayResult, error) {
//	var baseResponse BaseResponseWithArrayResult
//	err := json.NewDecoder(response.Body).Decode(&baseResponse)
//	if err != nil {
//		return nil, err
//	}
//
//	return &baseResponse, nil
//}
