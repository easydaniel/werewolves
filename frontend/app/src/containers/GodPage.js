// @format
import React from "react"

import { Grid } from "@material-ui/core"

import VoteProgress from "../components/VoteProgress"
import GameFunctions from "../components/GameFunctions"
import NightFlow from "../components/NightFlow"
import BoardInfo from "../components/BoardInfo"
import PlayerList from "../components/PlayerList"
import { ElectionEnum } from "../components/PlayerList"

import { makeStyles } from "@material-ui/core/styles"

const Board = {
  characters: { Wolf: 4, Villager: 4, God: 4 },
  gameID: "TA32EB",
  name: "狼王守衛",
  hasSheriff: true,
}

const Votes = [
  {
    0: [4, 7, 9],
    3: [1, 2, 6],
    7: [3, 5, 8, 10, 11, 12],
  },
  {
    0: [4, 7, 9],
    3: [1, 2, 6],
    7: [3, 5, 8, 10, 11, 12],
  },
]

const Flow = ["第一句", "第二句", "第三句", "第四句"]
const Players = [
  {
    name: "John",
    character: "Wolf",
    status: {
      isConnected: true,
      isAlive: true,
      election: ElectionEnum.INVOLVED,
    },
  },
  {
    name: "Eason",
    character: "Seer",
    status: {
      isConnected: true,
      isAlive: false,
      election: null,
    },
  },
  {
    name: "Daniel",
    character: "Hunter",
    status: {
      isConnected: false,
      isAlive: true,
      election: ElectionEnum.CANCELED,
    },
  },
]
const useStyles = makeStyles(theme => ({
  container: {
    flexGrow: 1,
    paddingTop: 120,
    height: "100vh",
  },
  spacer: {
    height: theme.spacing(2),
  },
}))

const GodPage = () => {
  const classes = useStyles()
  return (
    <Grid className={classes.container} container direction="row">
      <Grid xs />
      <Grid xs={2}>
        <BoardInfo info={Board} />
        <div className={classes.spacer} />
        <PlayerList players={Players} />
      </Grid>
      <Grid xs={4}>
        <Grid container justify="center">
          <GameFunctions />
        </Grid>
        <VoteProgress votes={Votes} />
      </Grid>
      <Grid xs={2}>
        <NightFlow flow={Flow} />
      </Grid>
      <Grid xs />
    </Grid>
  )
}

export default GodPage
