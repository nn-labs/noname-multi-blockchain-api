package eth

// GetEthUncleByBlockHashAndIndex https://eth.wiki/json-rpc/API#eth_getunclebyblockhashandindex
func GetEthUncleByBlockHashAndIndex(client IEthClient, params []string) (*BaseResponseWithStringResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_getUncleByBlockHashAndIndex",
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
