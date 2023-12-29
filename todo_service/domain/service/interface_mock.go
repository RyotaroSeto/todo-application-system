// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=interface_mock.go -package=service
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"
	model "todo_service/domain/model"

	gomock "go.uber.org/mock/gomock"
)

// MockTodo is a mock of Todo interface.
type MockTodo struct {
	ctrl     *gomock.Controller
	recorder *MockTodoMockRecorder
}

// MockTodoMockRecorder is the mock recorder for MockTodo.
type MockTodoMockRecorder struct {
	mock *MockTodo
}

// NewMockTodo creates a new mock instance.
func NewMockTodo(ctrl *gomock.Controller) *MockTodo {
	mock := &MockTodo{ctrl: ctrl}
	mock.recorder = &MockTodoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodo) EXPECT() *MockTodoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockTodo) Add(ctx context.Context, todo *model.Todo) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, todo)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockTodoMockRecorder) Add(ctx, todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTodo)(nil).Add), ctx, todo)
}

// Delete mocks base method.
func (m *MockTodo) Delete(ctx context.Context, id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodo)(nil).Delete), ctx, id)
}

// GetList mocks base method.
func (m *MockTodo) GetList(ctx context.Context) (*model.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx)
	ret0, _ := ret[0].(*model.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockTodoMockRecorder) GetList(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockTodo)(nil).GetList), ctx)
}

// UpdateStatus mocks base method.
func (m *MockTodo) UpdateStatus(ctx context.Context, id uint64, statusID model.TodoStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatus", ctx, id, statusID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatus indicates an expected call of UpdateStatus.
func (mr *MockTodoMockRecorder) UpdateStatus(ctx, id, statusID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockTodo)(nil).UpdateStatus), ctx, id, statusID)
}

// UpdateTitle mocks base method.
func (m *MockTodo) UpdateTitle(ctx context.Context, id uint64, title model.TodoTitle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTitle", ctx, id, title)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTitle indicates an expected call of UpdateTitle.
func (mr *MockTodoMockRecorder) UpdateTitle(ctx, id, title any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTitle", reflect.TypeOf((*MockTodo)(nil).UpdateTitle), ctx, id, title)
}