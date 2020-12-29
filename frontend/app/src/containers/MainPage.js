// @format
import React, { useState, useEffect } from "react"

import {
  MenuItem,
  FormControl,
  FormGroup,
  Paper,
  Button,
  TextField,
  Grid,
} from "@material-ui/core"

import { makeStyles } from "@material-ui/core/styles"

const useStyles = makeStyles(theme => ({
  root: {
    width: "100vw",
    height: "100vh",
  },
  container: {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    width: 350,
    height: 150,
    padding: theme.spacing(2),
  },
  rowContainer: {
    display: "flex",
    justifyContent: "space-between",
    margin: theme.spacing(1),
    width: 250,
  },
  rowInput: {
    width: 140,
  },
  button: {
    width: 80,
  },
}))

const Board = {
  characters: { Wolf: 4, Villager: 4, God: 4 },
  gameID: "TA32EB",
  name: "狼王守衛",
  hasSheriff: true,
  flow: [
    "狼人請睜眼。請選擇你們要擊殺的對象，確定是＿嗎？狼人請閉眼。",
    "預言家(通靈師)請睜眼。請選擇你要查驗的對象。他的身份是＿。預言家(通靈師)請閉眼。",
    "女巫請睜眼。今晚＿倒牌，請問你要使用解藥嗎？請問你要使用毒藥嗎？解藥給手勢，毒藥給數字。女巫請閉眼。",
    "獵人請睜眼。你的開槍狀態為＿？獵人請閉眼。",
    "守衛請睜眼。請選擇你要守護的對象，確定是＿嗎？守衛請閉眼。",
  ],
}

const createGame = boardIndex => {
  // createGame with boardIndex
  // response should be board info
  return Board
}

const joinGame = gameID => {
  // joinGame with gameID
  // response should be board info and game progress info
  return Board
}

const MainPage = ({ boardList, setBoard, setIsGod }) => {
  const classes = useStyles()
  const [boardIndex, setBoardIndex] = useState()
  const [gameID, setGameID] = useState()

  return (
    <Grid
      className={classes.root}
      container
      direction="row"
      justify="center"
      alignItems="center"
    >
      <Paper elevation={3} className={classes.container}>
        <FormControl>
          <FormGroup className={classes.rowContainer} row>
            <TextField
              className={classes.rowInput}
              id="board-select"
              select
              variant="outlined"
              label="板子"
              value={boardIndex}
              onChange={e => setBoardIndex(e.target.value)}
            >
              {boardList.map((board, idx) => (
                <MenuItem key={idx} value={idx}>
                  {board}
                </MenuItem>
              ))}
            </TextField>
            <Button
              className={classes.button}
              onClick={() => {
                setIsGod(true)
                const board = createGame(boardIndex)
                setBoard(board)
              }}
              variant="outlined"
              color="primary"
            >
              Create
            </Button>
          </FormGroup>
          <FormGroup className={classes.rowContainer} row>
            <TextField
              autoComplete="off"
              variant="outlined"
              color="secondary"
              className={classes.rowInput}
              value={gameID}
              onChange={e => setGameID(e.target.value)}
              id="game-id"
              label="遊戲 ID"
            />
            <Button
              className={classes.button}
              onClick={() => {
                const board = joinGame(gameID)
                setBoard(board)
              }}
              variant="outlined"
              color="secondary"
            >
              Join
            </Button>
          </FormGroup>
        </FormControl>
      </Paper>
    </Grid>
  )
}

export default MainPage
