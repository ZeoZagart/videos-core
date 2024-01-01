package server

import (
	"context"
	"fmt"

	DB "go.videos.core/db"

	pb "go.videos.core/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func registerAuthServer(server *grpc.Server, authDB DB.IAuthTable) {
	pb.RegisterAuthServiceServer(server, &authServer{authDB})
	reflection.Register(server)
}

type authServer struct {
	authTable DB.IAuthTable
}

func (s *authServer) CreateAccount(ctxt context.Context, request *pb.SignUpRequest) (*pb.LoginResponse, error) {
	user, err := s.authTable.CreateAccount(request)
	if err != nil {
		fmt.Printf("Error creating account: %v\n", err)
		return nil, err
	}
	return &pb.LoginResponse{Email: user.Email, Token: user.PasswordHash}, nil
}

func (s *authServer) Login(ctxt context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := s.authTable.VerifyUser(ctxt, request)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{Email: resp.Email, Token: resp.PasswordHash}, nil
}
