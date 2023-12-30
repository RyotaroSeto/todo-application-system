package infra

import (
	"context"
	"testing"
	"todo_service/domain/model"
)

func TestMain(m *testing.M) {
	setUpDB()
	m.Run()
	cleanUpDB()
}

func setUpDB() {
	cfg, _ := LoadConfig(context.Background())
	DBConnect(cfg.DNS(), cfg.DBMaxConn, cfg.DBMaxIdle)
}

func cleanUpDB() {
	Get().Exec("DELETE FROM todo.todos").
		Exec("DELETE FROM todo.status")
}

var todoID uint64 = 0

func insertTodo() model.Todo {
	todoID++
	todo := model.Todo{
		ID:         todoID,
		Title:      "test",
		TodoStatus: model.TodoStatusDoing,
	}
	Get().Create(&todo)
	return todo
}

func deleteTargetTodo(id uint64) {
	Get().Exec("DELETE FROM todo.todos WHERE id = ?", id)
}
