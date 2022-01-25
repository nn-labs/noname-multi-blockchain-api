package eth

// GetNetPeerCount https://eth.wiki/json-rpc/API#net_peercount
func GetNetPeerCount(client IEthClient) (*BaseResponseWithIntResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "net_peerCount",
		Params:  []string{},
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

	baseResponse, err := client.DecodeBaseResponseWithIntResult(response)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}
