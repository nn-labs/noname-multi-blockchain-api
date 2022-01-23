package eth

import (
	"encoding/json"
	"net/http"
)

func GetEthSync(client IEthClient) (*BaseResponseWithBoolResult, *EthSyncResponse, error) {
	request := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_syncing",
		Params:  []string{},
		Id:      "1",
	}

	baseRequest, err := client.EncodeBaseRequest(request)
	if err != nil {
		return nil, nil, err
	}

	response, err := client.Send(baseRequest)
	if err != nil {
		return nil, nil, err
	}

	baseResponse, err := client.DecodeBaseResponseWithBoolResult(response)
	if err != nil {
		return nil, nil, err
	}
	if baseResponse != nil {
		return baseResponse, nil, nil
	}

	decodedResponse, err := decodeEthSyncResponse(response)
	return nil, decodedResponse, err
}

func decodeEthSyncResponse(response *http.Response) (*EthSyncResponse, error) {
	var ethSyncResponse EthSyncResponse
	err := json.NewDecoder(response.Body).Decode(&ethSyncResponse)
	if err != nil {
		return nil, err
	}

	return &ethSyncResponse, nil
}
