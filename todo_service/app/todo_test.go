package app

import (
	"context"
	"testing"
	"todo_service/domain/model"
	"todo_service/domain/repository"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewTodoService(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := repository.NewMockDatabase(ctrl)
	type args struct {
		repo repository.Database
	}
	tests := []struct {
		name string
		args args
		want *TodoService
	}{
		{
			name: "success",
			args: args{
				repo: db,
			},
			want: &TodoService{
				repo: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewTodoService(tt.args.repo))
		})
	}
}

func TestTodoService_Add(t *testing.T) {
	type fields struct {
		repo *repository.MockDatabase
	}
	type args struct {
		todo *model.Todo
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		want      uint64
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				todo: &model.Todo{
					Title:      "test",
					TodoStatus: 1,
				},
			},
			setup: func(f *fields) {
				f.repo.EXPECT().AddTodo(argCtx, &model.Todo{
					Title:      "test",
					TodoStatus: 1,
				}).Return(uint64(1), nil)
			},
			want:      1,
			assertion: assert.NoError,
		},
		{
			name: "failed",
			args: args{
				todo: &model.Todo{
					Title:      "test",
					TodoStatus: 1,
				},
			},
			setup: func(f *fields) {
				f.repo.EXPECT().AddTodo(argCtx, &model.Todo{
					Title:      "test",
					TodoStatus: 1,
				}).Return(uint64(0), assert.AnError)
			},
			want:      0,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.repo = repository.NewMockDatabase(ctrl)

			if tt.setup != nil {
				tt.setup(&tt.fields)
			}
			s := &TodoService{
				repo: tt.fields.repo,
			}
			got, err := s.Add(context.Background(), tt.args.todo)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTodoService_GetList(t *testing.T) {
	type fields struct {
		repo *repository.MockDatabase
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		setup     func(f *fields)
		want      *model.Todos
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			setup: func(f *fields) {
				f.repo.EXPECT().GetTodoList(argCtx).Return(&model.Todos{
					{
						ID:         1,
						Title:      "test",
						TodoStatus: 1,
					},
				}, nil)
			},
			want: &model.Todos{
				{
					ID:         1,
					Title:      "test",
					TodoStatus: 1,
				},
			},
			assertion: assert.NoError,
		},
		{
			name: "failed",
			setup: func(f *fields) {
				f.repo.EXPECT().GetTodoList(argCtx).Return(nil, assert.AnError)
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.repo = repository.NewMockDatabase(ctrl)

			if tt.setup != nil {
				tt.setup(&tt.fields)
			}
			s := &TodoService{
				repo: tt.fields.repo,
			}
			got, err := s.GetList(context.Background())
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTodoService_UpdateStatus(t *testing.T) {
	type fields struct {
		repo *repository.MockDatabase
	}
	type args struct {
		id       uint64
		statusID model.TodoStatus
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id:       1,
				statusID: 1,
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					UpdateTodoStatus(argCtx, uint64(1), model.TodoStatus(1)).
					Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "failed",
			args: args{
				id:       1,
				statusID: 1,
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					UpdateTodoStatus(argCtx, uint64(1), model.TodoStatus(1)).
					Return(assert.AnError)
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.repo = repository.NewMockDatabase(ctrl)
			if tt.setup != nil {
				tt.setup(&tt.fields)
			}
			s := &TodoService{
				repo: tt.fields.repo,
			}
			tt.assertion(t, s.UpdateStatus(context.Background(), tt.args.id, tt.args.statusID))
		})
	}
}

func TestTodoService_UpdateTitle(t *testing.T) {
	type fields struct {
		repo *repository.MockDatabase
	}
	type args struct {
		id    uint64
		title model.TodoTitle
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id:    1,
				title: "test",
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					UpdateTodoTitle(argCtx, uint64(1), model.TodoTitle("test")).
					Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "failed",
			args: args{
				id:    1,
				title: "test",
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					UpdateTodoTitle(argCtx, uint64(1), model.TodoTitle("test")).
					Return(assert.AnError)
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.repo = repository.NewMockDatabase(ctrl)

			if tt.setup != nil {
				tt.setup(&tt.fields)
			}
			s := &TodoService{
				repo: tt.fields.repo,
			}
			tt.assertion(t, s.UpdateTitle(context.Background(), tt.args.id, tt.args.title))
		})
	}
}

func TestTodoService_Delete(t *testing.T) {
	type fields struct {
		repo *repository.MockDatabase
	}
	type args struct {
		id uint64
	}
	argCtx := gomock.AssignableToTypeOf(context.Background())
	tests := []struct {
		name      string
		fields    fields
		args      args
		setup     func(f *fields)
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					DeleteTodo(argCtx, uint64(1)).
					Return(nil)
			},
			assertion: assert.NoError,
		},
		{
			name: "failed",
			args: args{
				id: 1,
			},
			setup: func(f *fields) {
				f.repo.EXPECT().
					DeleteTodo(argCtx, uint64(1)).
					Return(assert.AnError)
			},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			tt.fields.repo = repository.NewMockDatabase(ctrl)

			if tt.setup != nil {
				tt.setup(&tt.fields)
			}

			s := &TodoService{
				repo: tt.fields.repo,
			}
			tt.assertion(t, s.Delete(context.Background(), tt.args.id))
		})
	}
}
