import { State, Condition, gameDone } from './game.js'
import { showAcknowledgeButton,
        hideAcknowledgeButton } from './acknowledge.js'

export function writeGameCondition() {
    if (gameDone() === true) {
        showAcknowledgeButton()
    } else {
        hideAcknowledgeButton()
    }

    let h
    switch (Condition) {
    case State.NORMAL:
        h = 'WISCONSIN CHESS'
        break
    case State.PROMOTION:
        h = 'PROMOTE'
        break
    case State.CHECK:
        h = 'CHECK'
        break
    case State.CHECKMATE:
        h = 'CHECKMATE'
        break
    case State.DRAW:
        h = 'DRAW'
        break
    case State.CONCEDED:
        h = 'CONCEDED'
        break
    case State.TIME_OVER:
        h = 'TIME OVER'
        break
    default:
        throw new Error('unknown game state ' + Condition)
    }
    document.querySelector('#condition').innerHTML = h
}
