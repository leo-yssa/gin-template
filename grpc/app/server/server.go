package server

import (
	"context"
	"gin-api/grpc/pkg/api/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)
type Server struct {
	auth.UnimplementedAuthServer
}

func (s *Server) Login(ctx context.Context, in *auth.User) (*auth.Result, error) {
	return nil, status.Errorf(codes.NotFound, "not implemented yet")
}