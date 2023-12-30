import { Todo } from '../model/Todo'

const todoDataUrl = 'http://localhost:2000/todo'

export const getAllTodosData = async () => {
  const response = await fetch(todoDataUrl, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  })
  console.log(1)
  console.log(response.json())
  //   return await fetch(todoDataUrl, {
  //     method: 'GET',
  //     headers: {
  //       'Content-Type': 'application/json',
  //     },
  //   }).then((response) => response.json())
  return [
    {
      id: '1',
      title: 'title1',
      status_id: 1,
      status_name: 'Doing',
    },
    {
      id: '2',
      title: 'title2',
      status_id: 2,
      status_name: 'Done',
    },
    {
      id: '3',
      title: 'title3',
      status_id: 2,
      status_name: 'Done',
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
  return {
    id: '4',
    title: 'title4',
    status_id: 1,
    status_name: 'Doing',
  }
}

export const deleteTodoData = async (id: string) => {
  // const response = await fetch(`${todoDataUrl}/${id}`, {
  //   method: 'DELETE',
  //   headers: {
  //     'Content-Type': 'application/json',
  //   },
  // })
  return id
}

export const updateTodoData = async (id: string, todo: Todo) => {
  //   const response = await axios.put(`${todoDataUrl}/${id}`, todo)
  //   return response.data
  return todo
}
