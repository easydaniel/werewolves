// @format

import React, { useEffect, useState } from "react"

import MainPage from "./containers/MainPage"
import GamePage from "./containers/GamePage"

// Sync with backend
const Boards = ["預女獵白", "狼王守衛"]

const getBoards = () => {
  // Fetch boards here. can use API
  return Boards
}

const App = () => {
  const [boardList, setBoardList] = useState([])
  const [game, setGame] = useState(null)
  const [isGod, setIsGod] = useState(false)

  useEffect(() => {
    setBoardList(getBoards)
  })

  const initialize = () => {
    setGame(null)
    setIsGod(false)
  }

  return game === null ? (
    <MainPage boardList={boardList} setGame={setGame} setIsGod={setIsGod} />
  ) : (
    <GamePage game={game} isGod={isGod} leaveGame={() => initialize()} />
  )
}

export default App
