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
}))

const createGame = boardIndex => {
  // createGame with boardIndex
  // response should be board info
  const Board = {
    characters: { Wolf: 4, Villager: 4, God: 4 },
    gameID: "TA32EB",
    name: "狼王守衛",
    hasSheriff: true,
    flow: ["第一句", "第二句", "第三句", "第四句"],
  }
  return Board
}

const joinGame = gameID => {
  // joinGame with gameID
  // response should be board info and game progress info
  const Board = {
    characters: { Wolf: 4, Villager: 4, God: 4 },
    gameID: "TA32EB",
    name: "狼王守衛",
    hasSheriff: true,
    flow: ["第一句", "第二句", "第三句", "第四句"],
  }
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
