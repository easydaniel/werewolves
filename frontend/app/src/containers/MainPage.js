// @format
import React, { useState } from "react"

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

const boards = ["預女獵白", "狼王守衛"]

const MainPage = () => {
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
              label="板子"
              value={boardIndex}
              onChange={e => setBoardIndex(e.target.value)}
            >
              {boards.map((board, idx) => (
                <MenuItem key={idx} value={idx}>
                  {board}
                </MenuItem>
              ))}
            </TextField>
            <Button variant="contained" color="primary">
              Create
            </Button>
          </FormGroup>
          <FormGroup className={classes.rowContainer} row>
            <TextField
              autoComplete="off"
              color="secondary"
              className={classes.rowInput}
              value={gameID}
              onChange={e => setGameID(e.target.value)}
              id="game-id"
              label="遊戲 ID"
            />
            <Button variant="contained" color="secondary">
              Join
            </Button>
          </FormGroup>
        </FormControl>
      </Paper>
    </Grid>
  )
}

export default MainPage
