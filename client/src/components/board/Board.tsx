import React, { memo } from 'react'

export type LayoutProps = {
  children: React.ReactNode
}

const Board: React.FC<LayoutProps> = memo((props) => {
  const { children } = props

  return (
    <>
      <h1>Board</h1>
      <p>with Tailwind CSS</p>
      {children}
    </>
  )
})

Board.displayName = 'Board'

export default Board
