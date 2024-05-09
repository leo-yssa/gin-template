package server

import (
	"context"
	"gin-api/grpc/pkg/api/auth"
	"gin-api/grpc/pkg/contract"
	"math/big"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type Server struct {
	auth.UnimplementedAuthServer
}

func (s *Server) Login(ctx context.Context, in *auth.User) (*auth.Result, error) {
	
	client, err := contract.NewClient(contract.HEDERA_TESTNET)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to connect ethereum client")
	}
	secret, err := contract.HexToECDSA(os.Getenv("OPERATOR_SECRET"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid privatekey")
	}
	from, err := contract.ToAddress(secret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to extract address")
	}
	transactor, err := client.NewTransactor(from, big.NewInt(0), secret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to make transactor")
	}
	address, transaction, _, err := client.Hello.DeployContract(transactor, client.Ethclient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to deploy contract")
	}

	return &auth.Result{
		Address: address.Hex(),
		TransactionHash: transaction.Hash().Hex(),
	}, nil

}