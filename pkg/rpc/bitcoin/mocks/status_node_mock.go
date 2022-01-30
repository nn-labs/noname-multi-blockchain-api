package mock_bitcoin

import (
	"github.com/stretchr/testify/mock"
	"nn-blockchain-api/pkg/rpc/bitcoin"
)

type MockStatusNode struct {
	mock.Mock
}

func (m *MockStatusNode) Status(client MockIBtcClient, network string) (*bitcoin.StatusNode, error) {
	ret := m.Called(client, network)
	return ret.Get(0).(*bitcoin.StatusNode), ret.Error(1)
}
