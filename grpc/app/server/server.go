package server

import (
	"context"
	"gin-api/grpc/pkg/api/auth"
	"gin-api/grpc/pkg/api/hello"
	"gin-api/grpc/pkg/contract"
	"math/big"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type Server struct {
	auth.UnimplementedAuthServer
	hello.UnimplementedHelloServer
}

func (s *Server) Login(ctx context.Context, in *auth.User) (*auth.Result, error) {
	return nil, status.Errorf(codes.Internal, "dose not implements yet")
}

func (s *Server) Deploy(ctx context.Context, in *hello.Empty) (*hello.DeployResponse, error) {
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

	return &hello.DeployResponse{
		Address: address.Hex(),
		TransactionHash: transaction.Hash().Hex(),
	}, nil
}

func (s *Server) Say(ctx context.Context, in *hello.SayRequest) (*hello.SayResponse, error) {
	client, err := contract.NewClient(contract.HEDERA_TESTNET)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to connect ethereum client")
	}
	address := contract.HexToAddress(in.GetAddress())
	secret, err := contract.HexToECDSA(os.Getenv("OPERATOR_SECRET"))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid privatekey")
	}
	from, err := contract.ToAddress(secret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to extract address")
	}
	message, err := client.Hello.Say(address, client.NewCallOpts(from), client.Ethclient)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to call contract")
	}
	
	return &hello.SayResponse{
		Message: message,
	}, nil
}	