package ui

import (
	"context"
	pb "gen/go/user"
	"user_service/domain/service"
)

type GRPCService struct {
	pb.UnimplementedUserApiServer
	svc service.User
}

func NewGRPCService(svc service.User) pb.UserApiServer {
	return &GRPCService{svc: svc}
}

func (s *GRPCService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// id, err := s.svc.Register(ctx, toModelUser(req.User))
	// if err != nil {
	// 	return nil, err
	// }

	// return &pb.RegisterResponse{Id: id}, nil
	return &pb.RegisterResponse{Id: 0}, nil
}

func (s *GRPCService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{}, nil
}
