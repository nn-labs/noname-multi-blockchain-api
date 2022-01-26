package eth

// GetEthCall https://eth.wiki/json-rpc/API#eth_call
func GetEthCall(client IEthClient, params map[string]string) (*BaseResponseWithStringResult, error) {
	request := BaseRequestWithMapParams{
		JsonRpc: "2.0",
		Method:  "eth_call",
		Params:  params,
		Id:      "1",
	}

	baseRequest, err := client.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := client.Send(baseRequest)
	if err != nil {
		return nil, err
	}

	baseResponse, err := client.DecodeBaseResponseWithStringResult(response)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}
