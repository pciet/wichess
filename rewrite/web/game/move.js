import { updateBoard, parseBoardUpdate } from './board_update.js'
import { fetchMoves } from './fetch_moves.js'
import { State } from './state.js'
import { replaceAndWriteGameCondition } from './condition.js'
import { moveSound } from './audio.js'
import { pauseWebSocket, unpauseWebSocket } from './websocket.js'

import { switchActive } from '../game.js'

export function doMove(fromIndex, toIndex, promotion = undefined, reversePromotion = false) {
    let body
    if (promotion === undefined) {
        body = JSON.stringify({f: fromIndex, t: toIndex})
    } else {
        body = JSON.stringify({p: promotion})
    }

    moveSound()
    replaceAndWriteGameCondition(State.NORMAL)

    // if the opponent move alert is received before the move response then the board will
    // be in an incorrect state
    pauseWebSocket()

    fetch('/move/'+GameInformation.ID, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: body
    }).then(r => r.json()).then(moveResponse => {
        updateBoard(parseBoardUpdate(moveResponse.d))
        if ('s' in moveResponse) {
            switch (moveResponse.s) {
            case 'p':
                // promotion needed
                replaceAndWriteGameCondition(State.PROMOTION)
                break
            case 'c':
                // this player takes an extra move
                fetchMoves()
                break
            default:
                throw new Error('unexpected /move response update state ' + moveResponse.s)
            }
        } else {
            switchActive()
        }
        unpauseWebSocket()
    })
}
