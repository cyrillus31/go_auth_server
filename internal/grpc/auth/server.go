package auth

import (
	ssov1 "github.com/cyrillus31/go_auth_server_proto/gen/go/sso"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type serverAPI struct {
	// Unimplemented Auth is needed so that we can launch our server wihtout
	// concrete endpoints implementations. But we will still create
	// our own placeholder methods below.
	ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	panic("Not implemented")
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	panic("Not implemented")
}

func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	panic("Not implemented")
}

// func (s *serverAPI) Login(ctx context.Context, *LoginRequest) (*LoginResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
// }
