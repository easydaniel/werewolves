import React from 'react';
import VoteList from './VoteList';

import {
  Timeline,
  TimelineItem,
  TimelineSeparator,
  TimelineDot,
  TimelineConnector,
  TimelineContent,
} from '@material-ui/lab';

import {Typography, Paper} from '@material-ui/core';
import {makeStyles} from '@material-ui/core/styles';

const useStyles = makeStyles(theme => ({
  paper: {
    padding: '6px 16px',
  },
}));

const VoteProgress = ({votes}) => {
  const classes = useStyles();
  return (
    <Timeline>
      {votes &&
        votes.map((vote, idx) => (
          <TimelineItem key={idx}>
            <TimelineSeparator>
              <TimelineDot />
              {idx !== votes.length - 1 && <TimelineConnector />}
            </TimelineSeparator>
            <TimelineContent>
              <Paper elevation={3} className={classes.paper}>
                <Typography variant="h6" component="h1">
                  第 {idx + 1} 天投票
                </Typography>
                <VoteList vote={vote} />
              </Paper>
            </TimelineContent>
          </TimelineItem>
        ))}
    </Timeline>
  );
};

export default VoteProgress;
