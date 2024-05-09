package main

import (
	"fmt"
	"net"
	"os"

	"gin-api/grpc/app/server"
	"gin-api/grpc/pkg/api/auth"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	auth.RegisterAuthServer(s, &server.Server{})
	if err := s.Serve(listener); err != nil {
		panic(err)
	}
}