import { Todo } from '../model/Todo'

const todoDataUrl = 'http://localhost:8282/todo'

export const getAllTodosData = async () => {
  const response = await fetch(todoDataUrl, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  })
  //   return await fetch(todoDataUrl, {
  //     method: 'GET',
  //     headers: {
  //       'Content-Type': 'application/json',
  //     },
  //   }).then((response) => response.json())
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
  const response = await fetch(todoDataUrl, {
    method: 'POST',
    body: JSON.stringify(todo),
    headers: {
      'Content-Type': 'application/json',
    },
  })
  return todo
}

export const deleteTodoData = async (id: string) => {
  const response = await fetch(`${todoDataUrl}/${id}`, {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
  })
  return id
}

// export const updateTodoData = async (id: string, todo: Todo) => {
//   //   const response = await axios.put(`${todoDataUrl}/${id}`, todo)
//   //   return response.data
//   return todo
// }
