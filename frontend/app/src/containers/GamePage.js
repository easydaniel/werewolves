// @format
import React, { useState, useEffect } from "react"
import _ from "lodash"
import { Grid } from "@material-ui/core"

import * as Api from "../lib/APIUtils"
import VoteProgress from "../components/VoteProgress"
import SelectDialog from "../components/SelectDialog"
import GameFunctions from "../components/GameFunctions"
import NightFlow from "../components/NightFlow"
import CharacterCard from "../components/CharacterCard"
import BoardInfo from "../components/BoardInfo"
import PlayerList from "../components/PlayerList"
import { ElectionEnum } from "../components/PlayerList"

import { makeStyles } from "@material-ui/core/styles"
const POLL_FREQUENCY = 1000 // in ms

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

const GamePage = ({ game, setGame, isGod, initialize }) => {
  const classes = useStyles()
  const { id, board, player: players } = game
  const { night_flow } = board

  // Poll game status
  useEffect(() => {
    const gameStatusPoll = setInterval(async () => {
      const [gameStatus, error] = await Api.getGameStatus(id)
      setGame(gameStatus)
    }, POLL_FREQUENCY)
    return () => clearInterval(gameStatusPoll)
  }, [game])

  const leaveRoom = async () => {
    if (isGod) {
    } else {
      const [message, error] = await Api.exitRoom(id)
    }
    initialize()
  }

  const [voteDialogOpen, setVoteDialogOpen] = useState(false)
  const [seatDialogOpen, setSeatDialogOpen] = useState(false)
  const playerFuncs = {
    openVoteDialog: () => setVoteDialogOpen(true),
    openSeatDialog: () => setSeatDialogOpen(true),
    exitSeat: async () => {
      const [message, error] = await Api.exitSeat(id)
      console.log(error)
    },
    leaveRoom,
  }
  const godFuncs = {
    leaveRoom,
  }
  // get available seat
  const availableSeats = _.map(players, (player, idx) => ({
    disabled: player !== null,
    label: player !== null ? player.name : `${idx + 1} 號`,
  }))
  // get available player list
  // const availablePlayers = _.map(
  //   Player,
  //   ({ status: { isAlive, election } }, idx) => ({
  //     label: `${idx + 1} 號玩家`,
  //     disabled: !isAlive,
  //   }),
  // )

  const _submitSeat = async seat => {
    const [message, error] = await Api.setSeat(id, Number(seat))
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
        <BoardInfo id={id} board={board} />
        <div className={classes.spacer} />
        <PlayerList isGod={isGod} players={players} />
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
          {/* <SelectDialog
            title={"選擇投票"}
            availableOptions={availablePlayers}
            open={voteDialogOpen}
            dismissDialog={() => setVoteDialogOpen(false)}
            submit={value => _submitVote(value)}
          /> */}
          <GameFunctions
            isGod={isGod}
            playerFuncs={playerFuncs}
            godFuncs={godFuncs}
          />
        </Grid>
        <VoteProgress votes={Votes} isGod={isGod} />
      </Grid>
      <Grid item lg={2} md={3} sm={12}>
        {isGod ? (
          <NightFlow flow={night_flow} />
        ) : (
          <CharacterCard character={Character} />
        )}
      </Grid>
      <Grid item lg sm md />
    </Grid>
  )
}

export default GamePage
