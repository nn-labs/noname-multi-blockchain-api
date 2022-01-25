package eth

// GetEthBlockNumber https://eth.wiki/json-rpc/API#eth_blocknumber
func GetEthBlockNumber(client IEthClient) (*BaseResponseWithStringResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []string{},
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
