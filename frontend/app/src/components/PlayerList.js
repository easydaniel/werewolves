import React from "react"

import { makeStyles, withStyles } from "@material-ui/core/styles"
import {
  List,
  ListItem,
  ListItemAvatar,
  ListItemIcon,
  Chip,
  Avatar,
  Badge,
} from "@material-ui/core"
import { green } from "@material-ui/core/colors"
import PanToolOutlinedIcon from "@material-ui/icons/PanToolOutlined"
import NotInterestedIcon from "@material-ui/icons/NotInterested"

const useStyles = makeStyles(theme => ({
  avatar: {
    width: theme.spacing(4),
    height: theme.spacing(4),
  },
  livePlayerAvatar: {
    backgroundColor: green[800],
  },
  characterChip: {
    marginRight: theme.spacing(1),
  },
  electionStatus: {
    marginLeft: theme.spacing(1),
  },
}))

const StyledBadge = withStyles(theme => ({
  badge: {
    backgroundColor: "#44b700",
    color: "#44b700",
    boxShadow: `0 0 0 2px ${theme.palette.background.paper}`,
    "&::after": {
      position: "absolute",
      top: 0,
      left: 0,
      width: "100%",
      height: "100%",
      content: '""',
    },
  },
}))(Badge)

const PlayerList = ({ players }) => {
  const classes = useStyles()
  return (
    players && (
      <List dense>
        {players.map(({ name, character }, idx) => (
          <ListItem key={idx}>
            <ListItemAvatar>
              <StyledBadge
                overlap="circle"
                anchorOrigin={{
                  vertical: "bottom",
                  horizontal: "right",
                }}
                variant="dot"
              >
                <Avatar
                  className={{
                    [classes.avatar]: true,
                    [classes.livePlayerAvatar]: true,
                  }}
                >
                  {idx + 1}
                </Avatar>
              </StyledBadge>
            </ListItemAvatar>
            <Chip
              className={classes.characterChip}
              variant="outlined"
              size="small"
              color="primary"
              label={character}
            />
            {name}
            <ListItemIcon className={classes.electionStatus}>
              <PanToolOutlinedIcon />
              <NotInterestedIcon />
            </ListItemIcon>
          </ListItem>
        ))}
      </List>
    )
  )
}

export default PlayerList
