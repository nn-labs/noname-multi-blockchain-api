package eth

func GetEthMining(client IEthClient) (*BaseResponseWithBoolResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_mining",
		Params:  []string{},
		Id:      "71",
	}

	baseRequest, err := client.EncodeBaseRequest(request)
	if err != nil {
		return nil, err
	}

	response, err := client.Send(baseRequest)
	if err != nil {
		return nil, err
	}

	baseResponse, err := client.DecodeBaseResponseWithBoolResult(response)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}
