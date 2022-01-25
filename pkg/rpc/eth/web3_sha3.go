package eth

// GetWeb3Sha3 https://eth.wiki/json-rpc/API#web3_sha3
func GetWeb3Sha3(client IEthClient, params []string) (*BaseResponseWithStringResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "web3_sha3",
		Params:  params,
		Id:      "64",
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
