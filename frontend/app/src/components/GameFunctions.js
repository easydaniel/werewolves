import React from "react"

import { ButtonGroup, Button } from "@material-ui/core"

const GameFunctions = ({ isGod }) => {
  return (
    <ButtonGroup>
      <Button>發身份</Button>
      <Button>開始投票</Button>
      <Button>統計票型</Button>
    </ButtonGroup>
  )
}

export default GameFunctions
