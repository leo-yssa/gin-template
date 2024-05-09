package client

import (
	"context"
	"fmt"
	"gin-api/grpc/pkg/api/auth"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Connection *grpc.ClientConn
	AuthClient auth.AuthClient
}

func NewClient(host, port string) (*Client, error) {
	connection, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil
	}
	client := auth.NewAuthClient(connection)
	return &Client{
		Connection: connection,
		AuthClient: client,
	}, nil
}

func (c *Client) Login(id, email, name string, roleId int32) (*auth.Result, error){
	defer c.Connection.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.AuthClient.Login(ctx, &auth.User{
		Id: id,
		Email: email,
		Name: name,
		RoleId: roleId,
	})
}