import React from "react"
import { Paper, Typography } from "@material-ui/core"
import { makeStyles } from "@material-ui/core/styles"
import { Skeleton } from "@material-ui/lab"
const useStyles = makeStyles(theme => ({
  container: {
    padding: theme.spacing(2),
  },
  media: {
    height: "480px",
    width: "100%",
    objectFit: "scale-down",
    marginTop: -theme.spacing(12),
    marginBottom: -theme.spacing(8),
  },
}))

const CharacterCard = ({ character }) => {
  const classes = useStyles()
  const { name, imgSrc } = character
  return (
    <Paper className={classes.container} elevation={3}>
      <Typography align="center" variant="h4">
        身份
      </Typography>
      {imgSrc ? (
        <img className={classes.media} src={imgSrc} />
      ) : (
        <Skeleton className={classes.media} />
      )}
      {name && (
        <Typography align="center" variant="h6">
          {name}
        </Typography>
      )}
    </Paper>
  )
}

export default CharacterCard
