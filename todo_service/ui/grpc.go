package ui

import (
	"context"
	pb "todo_service/proto"
)

type GRPCService struct {
	pb.UnimplementedTodoServiceServer
}

func NewGRPCService() pb.TodoServiceServer {
	return &GRPCService{}
}

func (s *GRPCService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{Todo: nil}, nil
}

func (s *GRPCService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *GRPCService) UpdateStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	return &pb.UpdateStatusResponse{}, nil
}

func (s *GRPCService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}
