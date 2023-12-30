package infra

import (
	"context"
	"testing"
	"todo_service/domain/model"
	"todo_service/domain/repository"

	"github.com/stretchr/testify/assert"
)

func TestNewTodoRepository(t *testing.T) {
	tests := []struct {
		name string
		want repository.Database
	}{
		{
			name: "success",
			want: &TodoRepository{
				db: Get(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewTodoRepository())
		})
	}
}

func TestTodoRepository_AddTodo(t *testing.T) {
	type args struct {
		todo *model.Todo
	}
	tests := []struct {
		name      string
		args      args
		want      uint64
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{todo: &model.Todo{
				Title:      "test",
				TodoStatus: model.TodoStatusDoing,
			}},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TodoRepository{
				db: Get(),
			}
			var beforeCount, afterCount int64
			r.db.Model(&model.Todo{}).Count(&beforeCount)
			got, err := r.AddTodo(context.Background(), tt.args.todo)
			tt.assertion(t, err)

			r.db.Model(&model.Todo{}).Count(&afterCount)
			if err == nil {
				assert.Equal(t, beforeCount+1, afterCount)
			} else {
				assert.Equal(t, beforeCount, afterCount)
			}

			deleteTargetTodo(got)
		})
	}
}

func TestTodoRepository_GetTodoList(t *testing.T) {
	todo := insertTodo()
	defer deleteTargetTodo(todo.ID)
	tests := []struct {
		name      string
		want      *model.Todos
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			want: &model.Todos{
				todo,
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TodoRepository{
				db: Get(),
			}
			got, err := r.GetTodoList(context.Background())
			tt.assertion(t, err)

			assert.Equal(t, (*tt.want)[0].ID, (*got)[0].ID)
			assert.Equal(t, (*tt.want)[0].Title, (*got)[0].Title)
			assert.Equal(t, (*tt.want)[0].TodoStatus, (*got)[0].TodoStatus)
		})
	}
}

func TestTodoRepository_UpdateTodoStatus(t *testing.T) {
	todo := insertTodo()
	defer deleteTargetTodo(todo.ID)
	type args struct {
		id       uint64
		statusID model.TodoStatus
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id:       todo.ID,
				statusID: 2,
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TodoRepository{
				db: Get(),
			}
			var beforeStatueID, afterStatueID model.TodoStatus
			r.db.Model(&model.Todo{}).Where(model.Todo{ID: todo.ID}).Select("todo_status").Scan(&beforeStatueID)
			assert.Equal(t, beforeStatueID, model.TodoStatusDoing)

			tt.assertion(t, r.UpdateTodoStatus(context.Background(), tt.args.id, tt.args.statusID))

			r.db.Model(&model.Todo{}).Where(model.Todo{ID: todo.ID}).Select("todo_status").Scan(&afterStatueID)
			assert.NotEqual(t, beforeStatueID, afterStatueID)
			assert.Equal(t, afterStatueID, model.TodoStatusDone)
		})
	}
}

func TestTodoRepository_UpdateTodoTitle(t *testing.T) {
	todo := insertTodo()
	defer deleteTargetTodo(todo.ID)
	type args struct {
		id    uint64
		title model.TodoTitle
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id:    todo.ID,
				title: "changeTitle",
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TodoRepository{
				db: Get(),
			}
			var beforeTitle, afterTitle model.TodoTitle
			r.db.Model(&model.Todo{}).Where(model.Todo{ID: todo.ID}).Select("title").Scan(&beforeTitle)
			assert.Equal(t, beforeTitle, model.TodoTitle("test"))

			tt.assertion(t, r.UpdateTodoTitle(context.Background(), tt.args.id, tt.args.title))

			r.db.Model(&model.Todo{}).Where(model.Todo{ID: todo.ID}).Select("title").Scan(&afterTitle)
			assert.NotEqual(t, beforeTitle, afterTitle)
			assert.Equal(t, afterTitle, model.TodoTitle("changeTitle"))
		})
	}
}

func TestTodoRepository_DeleteTodo(t *testing.T) {
	todo := insertTodo()
	type args struct {
		id uint64
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				id: todo.ID,
			},
			assertion: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &TodoRepository{
				db: Get(),
			}
			tt.assertion(t, r.DeleteTodo(context.Background(), tt.args.id))
		})
	}
}
