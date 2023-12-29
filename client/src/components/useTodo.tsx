import { useState, useEffect } from 'react'

import * as todoData from '../hooks/todos'
import { Todo } from '../model/Todo'

export const useTodo = () => {
  const [todoList, setTodoList] = useState<Todo[]>([])

  useEffect(() => {
    todoData.getAllTodosData().then((todo) => {
      console.log(...todo)
      setTodoList([...todo].reverse())
    })
  }, [])

  const toggleTodoListItemStatus = (id: string, done: boolean) => {
    // const todoItem = todoList.find((item: Todo) => item.id === id)
    // const newTodoItem: Todo = { ...todoItem!, done: !done }
    // todoData.updateTodoData(id, newTodoItem).then((updatedTodo) => {
    //   const newTodoList = todoList.map((item) =>
    //     item.id !== updatedTodo.id ? item : updatedTodo
    //   )
    //   setTodoList(newTodoList)
    // })
  }

  const addTodoListItem = (todoTitle: string) => {
    const newTodoItem = {
      id: '',
      title: todoTitle,
      status_id: 1,
      status_name: 'Doing',
    }

    todoData.addTodoData(newTodoItem).then((addTodo) => {
      setTodoList([addTodo, ...todoList])
    })
  }

  const deleteTodoListItem = (id: string) => {
    // todoData.deleteTodoData(id).then((deletedid) => {
    //   const newTodoList = todoList.filter((item) => item.id !== deletedid)
    //   setTodoList(newTodoList)
    // })
  }

  // 作成した関数を返す
  return {
    todoList,
    toggleTodoListItemStatus,
    addTodoListItem,
    deleteTodoListItem,
  }
}
