package eth

type BaseRequest struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params,array"`
	Id      string   `json:"id"`
}

type BaseRequestWithMapParams struct {
	JsonRpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params,array"`
	Id      string            `json:"id"`
}
