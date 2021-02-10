// @format

import React, { useState } from "react"

import * as Api from "./lib/APIUtils"
import { Snackbar, Button } from "@material-ui/core"

import MuiAlert from "@material-ui/lab/Alert"

import MainPage from "./containers/MainPage"
import GamePage from "./containers/GamePage"
import AuthPage from "./containers/AuthPage"

import { MessageProvider } from "./lib/message/context"

function Alert(props) {
  return <MuiAlert elevation={6} variant="filled" {...props} />
}

const App = () => {
  const [game, setGame] = useState(null)
  const [isGod, setIsGod] = useState(false)
  const [user, setUser] = useState(null)
  const [message, setMessage] = useState({ value: null, severity: null })
  const [open, setOpen] = useState(false)

  const initialize = () => {
    setGame(null)
    setIsGod(false)
  }

  return (
    <MessageProvider
      value={{
        message,
        setMessage: ({ value, severity }) => {
          setMessage({ value, severity })
          setOpen(true)
        },
      }}
    >
      <Button
        onClick={async () => {
          const resp = await Api.getUser()
          console.log(resp)
        }}
      >
        BTN
      </Button>
      {user === null ? (
        <AuthPage setUser={setUser} />
      ) : game === null ? (
        <MainPage setGame={setGame} setIsGod={setIsGod} />
      ) : (
        <GamePage
          game={game}
          setGame={setGame}
          isGod={isGod}
          initialize={initialize}
        />
      )}
      <Snackbar
        open={open}
        autoHideDuration={6000}
        onClose={() => setOpen(false)}
      >
        <Alert severity={message.severity}>{message.value}</Alert>
      </Snackbar>
    </MessageProvider>
  )
}

export default App
