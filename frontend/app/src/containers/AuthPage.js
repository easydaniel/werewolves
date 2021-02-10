// @format
import React, { useState, useContext } from "react"
import * as Api from "../lib/APIUtils"
import MessageContext from "../lib/message/context"

import {
  Button,
  FormControl,
  FormGroup,
  Grid,
  Paper,
  TextField,
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

  const { setMessage } = useContext(MessageContext)

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
                const [user, error] = await Api.login(username, password)
                if (error) {
                  setMessage({ value: error, severity: "error" })
                } else {
                  // setUser(user)
                  setMessage({ value: "Login success", severity: "success" })
                }
              }}
              variant="outlined"
              color="primary"
            >
              Login
            </Button>
            <Button
              className={classes.button}
              onClick={async () => {
                const [_, err] = await Api.registerUser(username, password)
                if (err) {
                  setMessage({ value: err, severity: "error" })
                } else {
                  const [user, _] = await Api.login(username, password)
                  setUser(user)
                  setMessage({ value: "Login success", severity: "success" })
                }
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
