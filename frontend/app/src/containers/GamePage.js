// @format
import React, { useState } from "react"
import _, { set } from "lodash"
import { Grid } from "@material-ui/core"

import VoteProgress from "../components/VoteProgress"
import SelectDialog from "../components/SelectDialog"
import GameFunctions from "../components/GameFunctions"
import NightFlow from "../components/NightFlow"
import CharacterCard from "../components/CharacterCard"
import BoardInfo from "../components/BoardInfo"
import PlayerList from "../components/PlayerList"
import { ElectionEnum } from "../components/PlayerList"

import { makeStyles } from "@material-ui/core/styles"

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

const GamePage = ({ game, isGod }) => {
  const classes = useStyles()
  const {
    gameID,
    board: { flow, ...info },
  } = game
  const [voteDialogOpen, setVoteDialogOpen] = useState(false)
  const [seatDialogOpen, setSeatDialogOpen] = useState(false)

  const playerFuncs = {
    openVoteDialog: () => setVoteDialogOpen(true),
    openSeatDialog: () => setSeatDialogOpen(true),
    getCharacter: () => {
      console.log("getcharacter")
    },
  }
  // get available seat
  const availableSeats = _.map(_.range(0, 5), (_, idx) => ({
    disabled: idx & 1,
    label: `${idx} 號`,
  }))
  // get available player list
  const availablePlayers = _.map(
    Players,
    ({ status: { isAlive, election } }, idx) => ({
      label: `${idx + 1} 號玩家`,
      disabled: !isAlive,
    }),
  )

  const _submitSeat = value => {
    console.log(value)
    setSeatDialogOpen(false)
  }

  const _submitVote = value => {
    console.log(value)
    setVoteDialogOpen(false)
  }

  return (
    <Grid className={classes.container} container direction="row">
      <Grid item lg sm md />
      <Grid item lg={2} md={3} sm={12}>
        <BoardInfo gameID={gameID} info={info} />
        <div className={classes.spacer} />
        <PlayerList isGod={isGod} players={Players} />
      </Grid>
      <Grid item lg={4} md={6} sm={12}>
        <Grid container justify="center">
          <SelectDialog
            title={"選擇號碼"}
            availableOptions={availableSeats}
            open={seatDialogOpen}
            dismissDialog={() => setSeatDialogOpen(false)}
            submit={value => _submitSeat(value)}
          />
          <SelectDialog
            title={"選擇投票"}
            availableOptions={availablePlayers}
            open={voteDialogOpen}
            dismissDialog={() => setVoteDialogOpen(false)}
            submit={value => _submitVote(value)}
          />
          <GameFunctions isGod={isGod} playerFuncs={playerFuncs} />
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
