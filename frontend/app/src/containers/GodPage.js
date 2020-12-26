// @format
import React from 'react'

import VoteProgress from '../components/VoteProgress'

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

const GodPage = () => {
    return (
        <VoteProgress votes={Votes} />
    )
}

export default GodPage
