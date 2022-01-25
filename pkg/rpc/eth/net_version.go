package eth

var (
	EthMainnet     string = "1"
	MordenTestnet  string = "2"
	RopstenTestnet string = "3"
	RinkebyTestnet string = "4"
	KovanTestnet   string = "42"
)

// GetNetVersion https://eth.wiki/json-rpc/API#net_version
func GetNetVersion(client IEthClient) (*BaseResponseWithStringResult, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "net_version",
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

	baseResponse, err := client.DecodeBaseResponseWithStringResult(response)
	if err != nil {
		return nil, err
	}

	return baseResponse, nil
}
