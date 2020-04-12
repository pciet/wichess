import { State, Condition, replaceCondition, gameDone } from './game.js'
import { showAcknowledgeButton,
        hideAcknowledgeButton } from './acknowledge.js'
import { showPromotion } from './promotion.js'
import { layoutSelector } from '../layout.js'

export function replaceAndWriteGameCondition(cond) {
    replaceCondition(cond)
    writeGameCondition()
}

export function writeGameCondition() {
    if (gameDone() === true) {
        showAcknowledgeButton()
    } else {
        hideAcknowledgeButton()
    }

    let h
    switch (Condition) {
    case State.NORMAL:
        h = '<div id="title">Wisconsin Chess</div>'
        break
    case State.PROMOTION:
        h = 'PROMOTE'
        showPromotion()
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
    layoutSelector('#condition', `
<div></div>
<div>`+h+`</div>
<div></div>
`)
}
