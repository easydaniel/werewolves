import React from "react"
import _ from "lodash"

import { makeStyles } from "@material-ui/core/styles"
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
})

const BoardInfo = ({ info }) => {
  const { gameID, name, characters } = info
  const numPlayers = _.reduce(characters, (result, val) => result + val, 0)
  const classes = useStyles()
  return (
    <Card className={classes.root}>
      <CardContent>
        <Typography color="textSecondary">{gameID}</Typography>
        <Typography variant="h5" component="h2">
          {`${name} (${numPlayers} 人)`}
        </Typography>
        <List dense>
          {_.map(characters, (count, character) => (
            <ListItem disableGutters key={character}>
              {`${character}: ${count}`}
            </ListItem>
          ))}
        </List>
      </CardContent>
    </Card>
  )
}

export default BoardInfo
