// @format
import React from "react"

import { Grid } from "@material-ui/core"

import VoteProgress from "../components/VoteProgress"
import GameFunctions from "../components/GameFunctions"
import NightFlow from "../components/NightFlow"
import CharacterCard from "../components/CharacterCard"
import BoardInfo from "../components/BoardInfo"
import PlayerList from "../components/PlayerList"
import { ElectionEnum } from "../components/PlayerList"

import { makeStyles } from "@material-ui/core/styles"

const Board = {
  characters: { Wolf: 4, Villager: 4, God: 4 },
  gameID: "TA32EB",
  name: "狼王守衛",
  hasSheriff: true,
  flow: ["第一句", "第二句", "第三句", "第四句"],
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

const Players = [
  {
    name: "John",
    character: "Wolf",
    status: {
      isVacant: false,
      isConnected: true,
      isAlive: true,
      election: ElectionEnum.INVOLVED,
    },
  },
  {
    name: "Eason",
    character: "Seer",
    status: {
      isVacant: false,
      isConnected: true,
      isAlive: false,
      election: null,
    },
  },
  {
    name: "Daniel",
    character: "Hunter",
    status: {
      isVacant: true,
      isConnected: false,
      isAlive: true,
      election: ElectionEnum.CANCELED,
    },
  },
]

const Character = {
  name: "預言家",
  imgSrc:
    "https://www.werewolfonline.game/pc/gw/20191028140957/img/role/1_pic_67fc4ec.png",
}

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

const GamePage = ({ board, isGod }) => {
  const classes = useStyles()
  const { flow, ...info } = board
  return (
    <Grid className={classes.container} container direction="row">
      <Grid item lg sm md />
      <Grid item lg={2} md={3} sm={12}>
        <BoardInfo info={info} />
        <div className={classes.spacer} />
        <PlayerList isGod={isGod} players={Players} />
      </Grid>
      <Grid item lg={4} md={6} sm={12}>
        <Grid container justify="center">
          <GameFunctions isGod={isGod} />
        </Grid>
        <VoteProgress votes={Votes} isGod={isGod} />
      </Grid>
      <Grid item lg={2} md={3} sm={12}>
        {isGod ? (
          <NightFlow flow={flow} />
        ) : (
          <CharacterCard character={Character} />
        )}
      </Grid>
      <Grid item lg sm md />
    </Grid>
  )
}

export default GamePage
