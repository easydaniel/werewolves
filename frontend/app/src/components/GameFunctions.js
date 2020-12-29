import React from "react"

import { ButtonGroup, Button } from "@material-ui/core"

const GodFunctions = () => {
  return (
    <ButtonGroup>
      <Button>發身份</Button>
      <Button>玩家自爆</Button>
      <Button>開啟投票</Button>
      <Button>統計票型</Button>
    </ButtonGroup>
  )
}

const PlayerFunctions = ({ openVoteDialog, openSeatDialog, getCharacter }) => {
  return (
    <ButtonGroup>
      <Button onClick={() => openSeatDialog()}>選取號碼</Button>
      <Button onClick={() => getCharacter()}>獲取身份</Button>
      <Button onClick={() => openVoteDialog()}>投票</Button>
      <Button>上警</Button>
      <Button>退水</Button>
    </ButtonGroup>
  )
}

const GameFunctions = ({ isGod, playerFuncs, godFuncs }) => {
  return isGod ? (
    <GodFunctions {...godFuncs} />
  ) : (
    <PlayerFunctions {...playerFuncs} />
  )
}

export default GameFunctions
