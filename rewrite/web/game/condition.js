import { Condition, replaceCondition, gameDone } from '../game.js'
import { layoutSelector } from '../layout.js'

import { State, States } from './state.js'
import { showAcknowledgeButton,
        hideAcknowledgeButton } from './acknowledge.js'
import { showPromotion } from './promotion.js'

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
    if (Condition === State.NORMAL) {
        h = '<div id="title">Wisconsin Chess</div>'
    } else {
        h = States[Condition].name
    }

    if (Condition === State.PROMOTION) {
        showPromotion()
    }

    layoutSelector('#condition', `
<div></div>
<div>`+h+`</div>
<div></div>
`)
}
