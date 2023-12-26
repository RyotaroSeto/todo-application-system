package model

type TaskID int64

type TaskTitle string

type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID     TaskID     `json:"id"`
	Title  TaskTitle  `json:"title"`
	Status TaskStatus `json:"status"`
}

type Tasks []Task
