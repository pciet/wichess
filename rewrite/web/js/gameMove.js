import { updateBoard } from './gameUpdate.js'
import { Board } from './game.js'

export function doMove(fromIndex, toIndex) {
    fetch('/move/'+GameInformation.ID, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(new MoveRequest(fromIndex, toIndex))
    }).then(r => r.json()).then(diff => {
        updateBoard(Board, diff)
    })
}

function MoveRequest(fromIndex, toIndex) {
    this.f = fromIndex
    this.t = toIndex
    return this
}
