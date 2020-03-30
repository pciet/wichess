import { updateBoard, parseBoardUpdate } from './board_update.js'
import { fetchMoves } from './fetch_moves.js'

export function webSocketPromise(gameID) {
    return new Promise(resolve => {
        resolve(new WebSocket('ws://'+window.location.host+'/alert/'+gameID))
    })
}

export function webSocketOnMessage(event) {
    const d = JSON.parse(event.data)
    // length 0 indicates no changes and a completed game, which will
    // be shown by the following call to /moves
    if (d.length !== 0) {
        updateBoard(parseBoardUpdate(d))
    }
    fetchMoves()
}
