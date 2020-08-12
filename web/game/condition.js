import { PlayerOrientation, ActiveOrientation, 
    Condition, replaceCondition, gameDone } from '../game.js'
import { layoutSelector } from '../layout.js'

import { State, States } from './state.js'
import { showPromotion } from './promotion.js'
import { removeAllSquareEventHandlers } from './board_click.js'
import { navigationLayout, addNavigationClickHandlers } from './layouts_navigation.js'

export function replaceAndWriteGameCondition(cond) {
    replaceCondition(cond)
    writeGameCondition()
}

export function writeGameCondition() {
    if (gameDone() === true) {
        removeAllSquareEventHandlers()
        const ok = layoutSelector('#navigation', navigationLayout())
        if (ok === null) {
            layoutSelector('#portraitnavigation', navigationLayout(true))
        }
        addNavigationClickHandlers()
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
