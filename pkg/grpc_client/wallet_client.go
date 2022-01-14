package grpc_client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "nn-blockchain-api/pkg/grpc_client/proto/wallet"
)

func NewWalletClient(host string) (pb.WalletServiceClient, error) {
	rpcConnection, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		defer rpcConnection.Close()
		return nil, err
	}

	return pb.NewWalletServiceClient(rpcConnection), nil
}
