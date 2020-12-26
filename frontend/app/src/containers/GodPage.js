// @format
import React from 'react';

import VoteProgress from '../components/VoteProgress';
import NightFlow from '../components/NightFlow';
import BoardInfo from '../components/BoardInfo';

const Board = {
  characters: { Wolf: 4, Villager: 4, God: 4 },
  gameID: 'TA32EB',
  name: '狼王守衛',
};

const Votes = [
  {
    0: [4, 7, 9],
    3: [1, 2, 6],
    7: [3, 5, 8, 10, 11, 12],
  },
  {
    0: [4, 7, 9],
    3: [1, 2, 6],
    7: [3, 5, 8, 10, 11, 12],
  },
];

const Flow = ['第一句', '第二句', '第三句', '第四句'];

const GodPage = () => {
  return (
    <>
      <BoardInfo info={Board} />
      <NightFlow flow={Flow} />
      <VoteProgress votes={Votes} />
    </>
  );
};

export default GodPage;
