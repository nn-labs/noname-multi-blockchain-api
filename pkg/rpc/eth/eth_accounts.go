package eth

// GetEthAccounts https://eth.wiki/json-rpc/API#eth_accounts
func GetEthAccounts(client IEthClient) (*BaseResponseWithArrayResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_accounts",
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
