import { PlayerOrientation, ActiveOrientation, 
    Condition, replaceCondition, gameDone } from '../game.js'
import { layoutSelector } from '../layout.js'

import { State, States } from './state.js'
import { showAcknowledgeButton, hideAcknowledgeButton } from './acknowledge.js'
import { showPromotion } from './promotion.js'
import { removeAllSquareEventHandlers } from './board_click.js'

export function replaceAndWriteGameCondition(cond) {
    replaceCondition(cond)
    writeGameCondition()
}

export function writeGameCondition() {
    if (gameDone() === true) {
        showAcknowledgeButton()
        removeAllSquareEventHandlers()
        if (Condition === State.CONCEDED) {
            document.querySelector('#backconcede').hidden = true
        }
        document.querySelector('#controls').hidden = true
    } else {
        hideAcknowledgeButton()
    }

    const s = document.querySelector('#statustext')
    if (Condition === State.NORMAL) {
        s.innerHTML = 'Wisconsin Chess'
        s.classList.remove('statusindicating')
    } else {
        s.innerHTML = States[Condition].name
        s.classList.add('statusindicating')
    }

    if ((Condition === State.PROMOTION) && (PlayerOrientation === ActiveOrientation)) {
        showPromotion()
    }
}
