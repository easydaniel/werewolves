// @format
import React from 'react'

import VoteProgress from '../components/VoteProgress'
import NightFlow from '../components/NightFlow'

const BoardInfo = {
    flow: ['1', '2', '3'],
    characters: { Wolf: 4, Villager: 4, God: 4},
    gameID: 'TA32EB'
}

const Votes = [
    {
        0: [4,7,9],
        3: [1,2,6],
        7: [3,5,8,10,11,12]
    },
    {
        0: [4,7,9],
        3: [1,2,6],
        7: [3,5,8,10,11,12]
    }
]

const Flow = [
  '第一句',
  '第二句',
  '第三句',
  '第四句',
]

const GodPage = () => {
    return (
        <NightFlow flow={Flow} />
    )
}

export default GodPage
