import React from "react"
import _ from "lodash"

import { makeStyles } from "@material-ui/core/styles"
import { red } from "@material-ui/core/colors"
import cx from "classnames"
import {
  Card,
  CardContent,
  List,
  ListItem,
  Typography,
} from "@material-ui/core"

const useStyles = makeStyles({
  root: {
    width: 275,
  },
  teamWolf: {
    color: red[800],
  },
})

const BoardInfo = ({ board, id }) => {
  const { name, characters, has_sheriff: hasSheriff } = board
  const numPlayers = characters.length
  const characterMap = {}
  _.each(characters, ({ name, team }) => {
    if (!_.has(characterMap, name)) {
      characterMap[name] = { team, count: 0 }
    }
    characterMap[name]["count"] += 1
  })
  const classes = useStyles()
  return (
    <Card elevation={3} className={classes.root}>
      <CardContent>
        <Typography color="textSecondary">{id}</Typography>
        <Typography variant="h5" component="h2">
          {`${name} (${numPlayers} 人)`}
        </Typography>
        {hasSheriff && (
          <Typography variant="subtitle2" color="textSecondary">
            有警長
          </Typography>
        )}
        <List dense>
          {_.map(characterMap, ({ team, count }, name) => (
            <ListItem
              className={cx({ [classes.teamWolf]: team === 1 })}
              disableGutters
              key={name}
            >{`${name}: ${count}`}</ListItem>
          ))}
        </List>
      </CardContent>
    </Card>
  )
}

export default BoardInfo
