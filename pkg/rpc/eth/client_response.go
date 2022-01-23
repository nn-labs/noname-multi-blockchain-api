package eth

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

type EthSyncResponse struct {
	BaseResponse
	Result struct {
		StartingBlock string `json:"startingBlock"`
		CurrentBlock  string `json:"currentBlock"`
		HighestBlock  string `json:"highestBlock"`
	}
}
