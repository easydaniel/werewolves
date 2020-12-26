import React from "react"
import _ from "lodash"

import { List, ListItem } from "@material-ui/core"

import ArrowRightAltIcon from "@material-ui/icons/ArrowRightAlt"

const VoteList = ({ vote }) => {
  return (
    <List dense>
      {_.map(vote, (voters, votee) => (
        <ListItem key={votee}>
          {voters.join(",")}
          <ArrowRightAltIcon />
          {votee === "0" ? "棄票" : votee}
        </ListItem>
      ))}
    </List>
  )
}

export default VoteList
