// @format
import React, { useState, useEffect } from "react"

import * as Api from "../lib/APIUtils"
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

const MainPage = ({ setGame, setIsGod }) => {
  const classes = useStyles()
  const [boardList, setBoardList] = useState([])
  const [boardIndex, setBoardIndex] = useState(null)
  const [gameID, setGameID] = useState()

  useEffect(async () => {
    const list = await Api.getBoardList()
    setBoardList(list)
  }, [])

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
              onClick={async () => {
                setIsGod(true)
                const game = await Api.createGame(boardList[boardIndex])
                setGame(game)
              }}
              disabled={boardIndex === null}
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
              onClick={async () => {
                const game = await Api.getGameStatus(gameID)
                setGame(game)
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
