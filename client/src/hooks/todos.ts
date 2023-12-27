// import axios from 'axios'
import { Todo } from '../model/Todo'

const todoDataUrl = 'http://localhost:3100/todos'

// 全TODOリスト取得
export const getAllTodosData = async () => {
  //   const response = await axios.get(todoDataUrl)
  //   return response.data
  return [
    {
      id: '1',
      content: 'content1',
      done: false,
    },
    {
      id: '2',
      content: 'content2',
      done: true,
    },
    {
      id: '3',
      content: 'content3',
      done: false,
    },
  ]
}

export const addTodoData = async (todo: Todo) => {
  //   const response = await axios.post(todoDataUrl, todo)
  //   return response.data
  return todo
}

// 1件のTODOを削除する
export const deleteTodoData = async (id: string) => {
  //   await axios.delete(`${todoDataUrl}/${id}`)
  return id
}

// 1件のTODOを更新する
export const updateTodoData = async (id: string, todo: Todo) => {
  //   const response = await axios.put(`${todoDataUrl}/${id}`, todo)
  //   return response.data
  return todo
}
