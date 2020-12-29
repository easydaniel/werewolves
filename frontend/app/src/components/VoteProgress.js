import React from "react"
import VoteList from "./VoteList"

import {
  Timeline,
  TimelineItem,
  TimelineSeparator,
  TimelineDot,
  TimelineConnector,
  TimelineContent,
  TimelineOppositeContent,
} from "@material-ui/lab"

import { Typography, Paper, TextField } from "@material-ui/core"
import { makeStyles } from "@material-ui/core/styles"

const useStyles = makeStyles(theme => ({
  paper: {
    padding: "6px 16px",
  },
}))

const VoteProgress = ({ votes, isGod }) => {
  const classes = useStyles()
  return (
    <Timeline align={isGod ? "left" : "alternate"}>
      <>
        {votes &&
          votes.map((vote, idx) => (
            <TimelineItem key={idx}>
              {isGod && (
                <TimelineOppositeContent>
                  <TextField
                    autoComplete={"off"}
                    variant={"outlined"}
                    fullWidth
                    label={`第 ${idx + 1} 晚`}
                    placeholder={`Ex:\t狼刀 X\n\t女巫開藥...`}
                    multiline
                    rows={Object.keys(votes).length + 4}
                    rowsMax={Object.keys(votes).length + 4}
                  />
                </TimelineOppositeContent>
              )}
              <TimelineSeparator>
                <TimelineDot />
                {idx !== votes.length - 1 && <TimelineConnector />}
              </TimelineSeparator>
              <TimelineContent>
                <Paper elevation={3} className={classes.paper}>
                  <Typography variant="h6" component="h1">
                    第 {idx + 1} 天放逐投票
                  </Typography>
                  <VoteList vote={vote} />
                </Paper>
              </TimelineContent>
            </TimelineItem>
          ))}
      </>
    </Timeline>
  )
}

export default VoteProgress
