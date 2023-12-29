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
    toggleTodoListItemStatus(todo.id, todo.status_name)
  const handleDeleteTodoListItem = () => deleteTodoListItem(todo.id)

  return (
    <>
      {todo.title}
      <button onClick={handleToggleTodoListItemStatus}>
        {todo.status_name === 'Doing' ? 'Done' : 'Doing'}
      </button>
      <button onClick={handleDeleteTodoListItem}>Delete</button>
    </>
  )
}
