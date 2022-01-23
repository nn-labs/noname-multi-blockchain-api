package eth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type ethClient struct {
	endpoint string
}

type BaseRequest struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params,array"`
	Id      string   `json:"id"`
}

type BaseResponse struct {
	Id      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
}

type BaseResponseWithIntResult struct {
	BaseResponse
	Result int `json:"result"`
}

type BaseResponseWithStringResult struct {
	BaseResponse
	Result string `json:"result"`
}

type BaseResponseWithBoolResult struct {
	BaseResponse
	Result bool `json:"result"`
}

type BaseResponseWithArrayResult struct {
	BaseResponse
	Result []string `json:"result,array"`
}

type IEthClient interface {
	Send(body io.Reader) (*http.Response, error)
	GetWeb3ClientVersion() (*BaseResponseWithStringResult, error)
	EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error)
	DecodeBaseResponseWithIntResult(response *http.Response) (*BaseResponseWithIntResult, error)
	DecodeBaseResponseWithStringResult(response *http.Response) (*BaseResponseWithStringResult, error)
	DecodeBaseResponseWithBoolResult(response *http.Response) (*BaseResponseWithBoolResult, error)
	DecodeBaseResponseWithArrayResult(response *http.Response) (*BaseResponseWithArrayResult, error)
}

func NewEthClient(endpoint string) IEthClient {
	return &ethClient{
		endpoint: endpoint,
	}
}

func (e *ethClient) Send(body io.Reader) (*http.Response, error) {
	resp, err := http.Post(e.endpoint, "application/json", body)
	return resp, err
}

func (e *ethClient) GetWeb3ClientVersion() (*BaseResponseWithStringResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "web3_clientVersion",
		Params:  []string{},
		Id:      "67",
	}

	reqBody, err := e.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	resp, err := e.Send(reqBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return e.DecodeBaseResponseWithStringResult(resp)
}

func (e *ethClient) EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(data)

	return reqBody, nil
}

func (e *ethClient) DecodeBaseResponseWithIntResult(response *http.Response) (*BaseResponseWithIntResult, error) {
	var baseResponse BaseResponseWithIntResult
	err := json.NewDecoder(response.Body).Decode(&baseResponse)
	if err != nil {
		return nil, err
	}

	return &baseResponse, nil
}

func (e *ethClient) DecodeBaseResponseWithStringResult(response *http.Response) (*BaseResponseWithStringResult, error) {
	var baseResponse BaseResponseWithStringResult
	err := json.NewDecoder(response.Body).Decode(&baseResponse)
	if err != nil {
		return nil, err
	}

	return &baseResponse, nil
}

func (e *ethClient) DecodeBaseResponseWithBoolResult(response *http.Response) (*BaseResponseWithBoolResult, error) {
	var baseResponse BaseResponseWithBoolResult
	err := json.NewDecoder(response.Body).Decode(&baseResponse)
	if err != nil {
		return nil, err
	}

	return &baseResponse, nil
}

func (e *ethClient) DecodeBaseResponseWithArrayResult(response *http.Response) (*BaseResponseWithArrayResult, error) {
	var baseResponse BaseResponseWithArrayResult
	err := json.NewDecoder(response.Body).Decode(&baseResponse)
	if err != nil {
		return nil, err
	}

	return &baseResponse, nil
}
