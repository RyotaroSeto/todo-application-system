'use client'
import React, { useRef } from 'react'
import { useTodo } from '../components/useTodo'
import { Todo } from '../model/Todo'
import { TodoAdd } from '../components/TodoAdd'
import { TodoList } from '../components/TodoList'
import { TodoTitle } from '../components/TodoTitle'

export default function Home() {
  const {
    todoList,
    toggleTodoListItemStatus,
    addTodoListItem,
    deleteTodoListItem,
  } = useTodo()
  const inputEl = useRef<HTMLTextAreaElement>(null)

  const handleAddTodoListItem = () => {
    if (inputEl.current?.value === '') {
      return
    }
    addTodoListItem(inputEl.current!.value)
    inputEl.current!.value = ''
  }

  // 未完了リスト
  const incompletedList = todoList.filter((todo: Todo) => !todo.done)
  // 完了リスト
  const completedList = todoList.filter((todo: Todo) => todo.done)

  return (
    <>
      <TodoTitle title="TODO" as="h1" />
      <TodoAdd
        buttonText="ADD"
        inputEl={inputEl}
        handleAddTodoListItem={handleAddTodoListItem}
      />
      <TodoList
        todoList={incompletedList}
        toggleTodoListItemStatus={toggleTodoListItemStatus}
        deleteTodoListItem={deleteTodoListItem}
        title="Doing"
        as="h2"
      />
      <TodoList
        todoList={completedList}
        toggleTodoListItemStatus={toggleTodoListItemStatus}
        deleteTodoListItem={deleteTodoListItem}
        title="Done"
        as="h2"
      />
    </>
  )
}
