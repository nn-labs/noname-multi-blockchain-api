package eth

type BaseRequest struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      string   `json:"id"`
}

type BaseRequestWithMapParams struct {
	JsonRpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
	Id      string            `json:"id"`
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
	Result []string `json:"result"`
}

type StatusNodeResponse struct {
	StartingBlock string `json:"startingBlock"`
	CurrentBlock  string `json:"currentBlock"`
	HighestBlock  string `json:"highestBlock"`
}
