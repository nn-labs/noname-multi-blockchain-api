package eth

import (
	"context"
	"encoding/json"
	"errors"
)

type Service interface {
	Status(ctx context.Context, network string) (*StatusNode, error)
}

type service struct {
	ethClient Client
}

func NewService(ethClient Client) (Service, error) {
	if ethClient == nil {
		return nil, errors.New("invalid ethereum client")
	}

	return &service{ethClient: ethClient}, nil
}

func (s *service) Status(ctx context.Context, network string) (*StatusNode, error) {
	req := BaseRequest{
		JsonRpc: "2.0",
		Method:  "eth_syncing",
		Params:  []string{},
	}

	msg := struct {
		JsonRpc string     `json:"jsonrpc"`
		Id      string     `json:"id"`
		Result  StatusNode `json:"result"`
	}{}

	body, err := s.ethClient.EncodeBaseRequest(req)
	if err != nil {
		return nil, err
	}

	response, err := s.ethClient.Send(ctx, body, network)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&msg)
	if err != nil {
		return nil, err
	}

	return &msg.Result, nil
}
