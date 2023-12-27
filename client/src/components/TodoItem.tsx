import { Todo } from '../model/Todo'

export const TodoItem = ({
  todo,
  toggleTodoListItemStatus,
  deleteTodoListItem,
}: {
  todo: Todo
  toggleTodoListItemStatus: any
  deleteTodoListItem: any
}) => {
  const handleToggleTodoListItemStatus = () =>
    toggleTodoListItemStatus(todo.id, todo.done)
  const handleDeleteTodoListItem = () => deleteTodoListItem(todo.id)

  return (
    <>
      {todo.content}
      <button onClick={handleToggleTodoListItemStatus}>
        {todo.done ? 'back to Progress' : 'Done'}
      </button>
      <button onClick={handleDeleteTodoListItem}>Delete</button>
    </>
  )
}
