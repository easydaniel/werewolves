// @format
import React, { useState } from "react"
import * as Api from "../lib/APIUtils"

import {
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
    width: 280,
    height: 240,
    padding: theme.spacing(2),
  },
  rowContainer: {
    display: "flex",
    justifyContent: "space-around",
    margin: theme.spacing(1),
    width: 200,
  },
  rowInput: {
    width: 180,
  },
  button: {
    width: 80,
  },
}))

const AuthPage = ({ setUser }) => {
  const classes = useStyles()
  const [username, setUsername] = useState()
  const [password, setPassword] = useState()

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
              autoComplete="off"
              variant="outlined"
              color="primary"
              className={classes.rowInput}
              value={username}
              onChange={e => setUsername(e.target.value)}
              id="username"
              label="Username"
            />
          </FormGroup>
          <FormGroup className={classes.rowContainer} row>
            <TextField
              autoComplete="off"
              variant="outlined"
              color="primary"
              className={classes.rowInput}
              value={password}
              onChange={e => setPassword(e.target.value)}
              id="password"
              label="Password"
              type="password"
            />
          </FormGroup>
          <FormGroup className={classes.rowContainer} row>
            <Button
              className={classes.button}
              onClick={async () => {
                const user = await Api.login(username, password)
                console.log(user)
              }}
              variant="outlined"
              color="primary"
            >
              Login
            </Button>
            <Button
              className={classes.button}
              onClick={async () => {
                await Api.logout()
              }}
              variant="outlined"
              color="primary"
            >
              Logout
            </Button>
            <Button
              className={classes.button}
              onClick={async () => {
                const resp = await Api.registerUser(username, password)

                console.log(resp)
              }}
              variant="outlined"
              color="secondary"
            >
              Register
            </Button>
          </FormGroup>
        </FormControl>
      </Paper>
    </Grid>
  )
}

export default AuthPage
