package ui

import (
	"context"
	pb "gen/go/todo"
	"testing"
	"todo_service/domain/model"
	"todo_service/domain/service"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewGRPCService(t *testing.T) {
	ctrl := gomock.NewController(t)
	svc := service.NewMockTodo(ctrl)
	type args struct {
		todo service.Todo
	}
	tests := []struct {
		name string
		args args
		want pb.TodoApiServer
	}{
		{
			name: "success",
			args: args{svc},
			want: &GRPCService{todo: svc},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewGRPCService(tt.args.todo))
		})
	}
}

func TestGRPCService_Get(t *testing.T) {
	type fields struct {
		todo *service.MockTodo
	}
	type args struct {
		req *pb.GetRequest
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		want      *pb.GetResponse
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{req: &pb.GetRequest{}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					GetList(argCtx).
					Return(&model.Todos{{
						ID:         1,
						Title:      "title",
						TodoStatus: 1,
					}}, nil)
			},
			want: &pb.GetResponse{
				List: []*pb.Todo{
					{
						Id:         1,
						Title:      "title",
						Status:     1,
						StatusName: "DOING",
					},
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "error",
			args: args{req: &pb.GetRequest{}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					GetList(argCtx).
					Return(nil, assert.AnError)
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.todo = service.NewMockTodo(ctrl)
			if tt.setup != nil {
				tt.setup(&tt.fields)
			}

			s := &GRPCService{
				todo: tt.fields.todo,
			}
			got, err := s.Get(context.Background(), tt.args.req)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGRPCService_Add(t *testing.T) {
	type fields struct {
		todo *service.MockTodo
	}
	type args struct {
		req *pb.AddRequest
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		want      *pb.AddResponse
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{req: &pb.AddRequest{Todo: &pb.Todo{
				Title:      "title",
				Status:     1,
				StatusName: "DOING",
			}}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					Add(argCtx, &model.Todo{
						Title:      "title",
						TodoStatus: 1,
					}).
					Return(uint64(1), nil)
			},
			want:      &pb.AddResponse{Id: 1},
			assertion: assert.NoError,
		},
		{
			name: "error",
			args: args{req: &pb.AddRequest{Todo: &pb.Todo{
				Title:      "title",
				Status:     1,
				StatusName: "DOING",
			}}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					Add(argCtx, &model.Todo{
						Title:      "title",
						TodoStatus: 1,
					}).
					Return(uint64(0), assert.AnError)
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.todo = service.NewMockTodo(ctrl)
			if tt.setup != nil {
				tt.setup(&tt.fields)
			}

			s := &GRPCService{
				todo: tt.fields.todo,
			}
			got, err := s.Add(context.Background(), tt.args.req)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGRPCService_UpdateStatus(t *testing.T) {
	type fields struct {
		todo *service.MockTodo
	}
	type args struct {
		req *pb.UpdateStatusRequest
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		want      *pb.UpdateStatusResponse
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{req: &pb.UpdateStatusRequest{
				Id:     1,
				Status: 1,
			}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					UpdateStatus(argCtx, uint64(1), model.TodoStatus(1)).
					Return(nil)
			},
			want:      &pb.UpdateStatusResponse{},
			assertion: assert.NoError,
		},
		{
			name: "error",
			args: args{req: &pb.UpdateStatusRequest{
				Id:     1,
				Status: 1,
			}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					UpdateStatus(argCtx, uint64(1), model.TodoStatus(1)).
					Return(assert.AnError)
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.todo = service.NewMockTodo(ctrl)
			if tt.setup != nil {
				tt.setup(&tt.fields)
			}

			s := &GRPCService{
				todo: tt.fields.todo,
			}
			got, err := s.UpdateStatus(context.Background(), tt.args.req)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGRPCService_UpdateTitle(t *testing.T) {
	type fields struct {
		todo *service.MockTodo
	}
	type args struct {
		req *pb.UpdateTitleRequest
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		want      *pb.UpdateTitleResponse
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{req: &pb.UpdateTitleRequest{
				Id:    1,
				Title: "title",
			}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					UpdateTitle(argCtx, uint64(1), model.TodoTitle("title")).
					Return(nil)
			},
			want:      &pb.UpdateTitleResponse{},
			assertion: assert.NoError,
		},
		{
			name: "error",
			args: args{req: &pb.UpdateTitleRequest{
				Id:    1,
				Title: "title",
			}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					UpdateTitle(argCtx, uint64(1), model.TodoTitle("title")).
					Return(assert.AnError)
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.todo = service.NewMockTodo(ctrl)
			if tt.setup != nil {
				tt.setup(&tt.fields)
			}

			s := &GRPCService{
				todo: tt.fields.todo,
			}
			got, err := s.UpdateTitle(context.Background(), tt.args.req)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGRPCService_Delete(t *testing.T) {
	type fields struct {
		todo *service.MockTodo
	}
	type args struct {
		req *pb.DeleteRequest
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		want      *pb.DeleteResponse
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{req: &pb.DeleteRequest{
				Id: 1,
			}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					Delete(argCtx, uint64(1)).
					Return(nil)
			},
			want:      &pb.DeleteResponse{},
			assertion: assert.NoError,
		},
		{
			name: "error",
			args: args{req: &pb.DeleteRequest{
				Id: 1,
			}},
			setup: func(f *fields) {
				f.todo.EXPECT().
					Delete(argCtx, uint64(1)).
					Return(assert.AnError)
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.todo = service.NewMockTodo(ctrl)
			if tt.setup != nil {
				tt.setup(&tt.fields)
			}

			s := &GRPCService{
				todo: tt.fields.todo,
			}
			got, err := s.Delete(context.Background(), tt.args.req)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
