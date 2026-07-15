package auth

import (
	"context"

	ssov1 "github.com/BKBTVETH98/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if req.Email == "52" || req.Password == "52" {

	}
	return &ssov1.LoginResponse{Token: "CAP TOKEN 52525252 ALBLAAAAK"}, nil
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {

	return &ssov1.RegisterResponse{UserId: 0}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context, req *ssov1.IsAdminRequest) (
	*ssov1.IsAdminResponse, error) {
	panic("unimplemented")
}
