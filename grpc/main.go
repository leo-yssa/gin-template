package main

import (
	"fmt"
	"net"
	"os"

	"gin-api/grpc/app/server"
	"gin-api/grpc/pkg/api/auth"
	"gin-api/grpc/pkg/api/hello"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),
			grpc_logrus.UnaryServerInterceptor(logrusEntry),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)
	auth.RegisterAuthServer(s, &server.Server{})
	hello.RegisterHelloServer(s, &server.Server{})
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}