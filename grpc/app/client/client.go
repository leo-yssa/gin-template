package client

import (
	"context"
	"fmt"
	"gin-api/grpc/pkg/api/auth"
	"gin-api/grpc/pkg/api/hello"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	Connection *grpc.ClientConn
	AuthClient auth.AuthClient
	HelloCient hello.HelloClient
}

func NewClient(host, port string) (*Client, error) {
	connection, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil
	}
	authClient := auth.NewAuthClient(connection)
	helloClient := hello.NewHelloClient(connection)
	return &Client{
		Connection: connection,
		AuthClient: authClient,
		HelloCient: helloClient,
	}, nil
}

func (c *Client) Login(id, email, name string, roleId int32) (*auth.Result, error){
	defer c.Connection.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()
	return c.AuthClient.Login(ctx, &auth.User{
		Id: id,
		Email: email,
		Name: name,
		RoleId: roleId,
	})
}

func (c *Client) Deploy() (*hello.DeployResponse, error) {
	defer c.Connection.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()
	return c.HelloCient.Deploy(ctx, &hello.Empty{})
}

func (c *Client) Say(address string) (*hello.SayResponse, error) {
	defer c.Connection.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()
	return c.HelloCient.Say(ctx, &hello.SayRequest{Address: address})
}