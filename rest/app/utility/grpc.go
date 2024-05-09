package utility

import (
	"gin-api/grpc/app/client"
	"os"
)

func NewClient() (*client.Client, error) {
	return client.NewClient(os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT"))
}