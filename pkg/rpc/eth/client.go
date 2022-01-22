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
	Result  string `json:"result"`
}

type IEthClient interface {
	Send(body io.Reader) (*http.Response, error)
	GetWeb3ClientVersion() (*BaseResponse, error)
	EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error)
	DecodeBaseResponse(response *http.Response) (*BaseResponse, error)
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

func (e *ethClient) GetWeb3ClientVersion() (*BaseResponse, error) {
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

	return e.DecodeBaseResponse(resp)
}

func (e *ethClient) EncodeBaseRequest(request BaseRequest) (*bytes.Buffer, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	reqBody := bytes.NewBuffer(data)

	return reqBody, nil
}

func (e *ethClient) DecodeBaseResponse(response *http.Response) (*BaseResponse, error) {
	var baseResponse BaseResponse
	err := json.NewDecoder(response.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &baseResponse, nil
}
