import React from "react"
import cx from "classnames"
import { makeStyles, withStyles } from "@material-ui/core/styles"
import {
  List,
  ListItem,
  ListItemAvatar,
  ListItemIcon,
  ListSubheader,
  Chip,
  Avatar,
  Badge,
} from "@material-ui/core"
import { green, red } from "@material-ui/core/colors"
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
  deadPlayerAvatar: {
    backgroundColor: red[800],
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

export const ElectionEnum = Object.freeze({
  INVOLVED: 1,
  CANCELED: 2,
})

const ElectionStatusIcon = ({ election }) => {
  switch (election) {
    case ElectionEnum.INVOLVED:
      return <PanToolOutlinedIcon />
    case ElectionEnum.CANCELED:
      return <NotInterestedIcon />
    default:
      return null
  }
}

const PlayerList = ({ players, isGod }) => {
  const classes = useStyles()
  return (
    players && (
      <List dense subheader={<ListSubheader>玩家列表</ListSubheader>}>
        {players.map(
          (
            {
              name,
              character,
              status: { isVacant, isConnected, isAlive, election },
            },
            idx,
          ) => (
            <ListItem key={idx}>
              <ListItemAvatar>
                <StyledBadge
                  overlap="circle"
                  anchorOrigin={{
                    vertical: "bottom",
                    horizontal: "right",
                  }}
                  invisible={!isConnected}
                  variant={"dot"}
                >
                  <Avatar
                    className={cx({
                      [classes.avatar]: true,
                      [classes.livePlayerAvatar]: !isVacant && isAlive,
                      [classes.deadPlayerAvatar]: !isVacant && !isAlive,
                    })}
                  >
                    {idx + 1}
                  </Avatar>
                </StyledBadge>
              </ListItemAvatar>
              {isGod && (
                <Chip
                  className={classes.characterChip}
                  variant="outlined"
                  size="small"
                  color="primary"
                  label={character}
                />
              )}
              {name || `${idx + 1} 號玩家`}
              <ListItemIcon className={classes.electionStatus}>
                <ElectionStatusIcon election={election} />
              </ListItemIcon>
            </ListItem>
          ),
        )}
      </List>
    )
  )
}

export default PlayerList
