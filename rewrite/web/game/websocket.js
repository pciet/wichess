import { updateBoard, parseBoardUpdate } from './board_update.js'
import { fetchMoves } from './fetch_moves.js'
import { showPromotion } from './promotion.js'
import { replaceAndWriteGameCondition } from './condition.js'
import { State } from './state.js'
import { unshowMoveablePieces } from './moves.js'
import { moveSound } from './audio.js'
import { ComputerPlayerName } from './players.js'

import { switchActive } from '../game.js'

export function webSocketPromise(gameID) {
    return new Promise(resolve => {
        resolve(new WebSocket('ws://'+window.location.host+'/alert/'+gameID))
    })
}

let paused = false
let queue = []

export function pauseWebSocket() { paused = true }

export function unpauseWebSocket() {
    // assumes this function cannot be interrupted, all messages will be processed before anything
    // else happens
    paused = false
    const l = queue.length
    for (let i = 0; i < l; i++) {
        webSocketOnMessage(queue.shift())
    }
}

export function webSocketOnMessage(event) {
    if (paused === true) {
        queue.push(event)
        return
    }
    const alert = JSON.parse(event.data)
    unshowMoveablePieces()

    if (window.GameInformation.Black.Name !== ComputerPlayerName) {
        moveSound()
    }

    if ((alert.d !== undefined) && (alert.d.length !== 0)) {
        updateBoard(parseBoardUpdate(alert.d))
    }

    if (alert.s === undefined) {
        switchActive()
        fetchMoves()
        return
    }

    // see docs/promotion.md for an overview of the reverse promotion network communication
    switch (alert.s) {
    case 'p':
        throw new Error('unexpected promotion needed alert WebSocket message: ' + alert)
    case 'w':
        // wait, do nothing
        break
    case 'c':
        // continue
        switchActive()
        fetchMoves()
        break
    case 'ch':
        // this player's previous move caused a check
        replaceAndWriteGameCondition(State.CHECK)
        break
    case 'dr':
        // prev move caused draw
        replaceAndWriteGameCondition(State.DRAW)
        break
    case 'chm':
        // prev move caused checkmate
        replaceAndWriteGameCondition(State.CHECKMATE)
        break
    case 'co':
        // opponent conceded
        replaceAndWriteGameCondition(State.CONCEDED)
        break
    default:
        throw new Error('unknown alert state: ' + alert)
    }
}

