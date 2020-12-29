// @format

import React, { useEffect, useState } from "react"

import MainPage from "./containers/MainPage"
import GamePage from "./containers/GamePage"

// Sync with backend
const Boards = ["預女獵白", "狼王守衛"]

const getBoards = () => {
  // Fetch boards here
  return Boards
}

const App = () => {
  const [boardList, setBoardList] = useState([])
  const [board, setBoard] = useState(null)
  const [isGod, setIsGod] = useState(false)

  useEffect(() => {
    setBoardList(getBoards)
  })
  useEffect(() => {}, [board])

  return board === null ? (
    <MainPage boardList={boardList} setBoard={setBoard} setIsGod={setIsGod} />
  ) : (
    <GamePage board={board} isGod={isGod} />
  )
}

export default App
