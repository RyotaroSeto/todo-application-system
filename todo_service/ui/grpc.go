package ui

import (
	"context"
	pb "gen/go/todo"
	"todo_service/domain/model"
	"todo_service/domain/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCService struct {
	pb.UnimplementedTodoApiServer
	todo service.Todo
}

func NewGRPCService(todo service.Todo) pb.TodoApiServer {
	return &GRPCService{todo: todo}
}

func (s *GRPCService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	list, err := s.todo.GetList(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		List: toPbTodoList(list),
	}, nil
}

func (s *GRPCService) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	id, err := s.todo.Add(ctx, toModelTodo(req.Todo))
	if err != nil {
		return nil, err
	}

	return &pb.AddResponse{Id: id}, nil
}

func (s *GRPCService) UpdateStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.UpdateStatusResponse, error) {
	if err := s.todo.UpdateStatus(ctx, req.Id, model.TodoStatus(req.Status)); err != nil {
		return nil, err
	}
	return &pb.UpdateStatusResponse{}, nil
}

func (s *GRPCService) UpdateTitle(ctx context.Context, req *pb.UpdateTitleRequest) (*pb.UpdateTitleResponse, error) {
	if err := s.todo.UpdateTitle(ctx, req.Id, model.TodoTitle(req.Title)); err != nil {
		return nil, err
	}
	return &pb.UpdateTitleResponse{}, nil
}

func (s *GRPCService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	if err := s.todo.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (s *GRPCService) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (s *GRPCService) Watch(req *pb.HealthCheckRequest, stream pb.TodoApi_WatchServer) error {
	return status.Error(codes.Unimplemented, "service watch is not implemented current version.")
}

func toModelTodo(req *pb.Todo) *model.Todo {
	return &model.Todo{
		Title:      model.TodoTitle(req.Title),
		TodoStatus: model.TodoStatus(req.Status),
	}
}

func toPbTodoList(todoList *model.Todos) []*pb.Todo {
	list := make([]*pb.Todo, 0, len(*todoList))
	for _, todo := range *todoList {
		list = append(list, &pb.Todo{
			Id:         todo.ID,
			Title:      string(todo.Title),
			Status:     pb.TodoStatus(pb.TodoStatus_value[todo.String()]),
			StatusName: todo.ShortString(),
		})
	}
	return list
}
