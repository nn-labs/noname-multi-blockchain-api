package eth

// GetEthCompilers https://eth.wiki/json-rpc/API#eth_getcode
func GetEthCompilers(client IEthClient, params []string) (*BaseResponseWithArrayResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_getCompilers",
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

	baseResponse, err := client.DecodeBaseResponseWithArrayResult(response)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}
