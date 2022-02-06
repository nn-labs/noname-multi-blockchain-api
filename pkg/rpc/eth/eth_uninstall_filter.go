package eth

// GetEthUninstallFilter https://eth.wiki/json-rpc/API#eth_uninstallfilter
func GetEthUninstallFilter(client IEthClient, params []string) (*BaseResponseWithBoolResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_uninstallFilter",
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

	baseResponse, err := client.DecodeBaseResponseWithBoolResult(response)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}
