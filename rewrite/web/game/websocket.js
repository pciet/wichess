import { updateBoard, parseBoardUpdate } from './board_update.js'
import { fetchMoves } from './fetch_moves.js'
import { showPromotion } from './promotion.js'

export function webSocketPromise(gameID) {
    return new Promise(resolve => {
        resolve(new WebSocket('ws://'+window.location.host+'/alert/'+gameID))
    })
}

export function webSocketOnMessage(event) {
    const alert = JSON.parse(event.data)

    if (alert.d.length !== 0) {
        updateBoard(parseBoardUpdate(alert.d))
    }

    // see docs/promotion.md for an overview of the
    // reverse promotion network communication
    if ('p' in alert) {
        // if the alert indicates a reverse promotion is needed
        // then a /move response with the promotion pick is done, 
        // otherwise the normal move communication continues
        if (alert.p === 'rp') {
            // swapActive()
            showPromotion()
        } else if (alert.p === 'rpd') {
            // if a reverse promotion is done then this webpage
            // is waiting for the opponent to make their next move
        } else {
            throw new Error('got alert with p (promotion) key with invalid value ' + alert.p)
        }
    } else {
        // swapActive()
        // with no reverse promotion the normal action is
        // to continue onto this player's next move choice
        fetchMoves()
    }
}
