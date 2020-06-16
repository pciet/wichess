import { gameDone } from '../game.js'

import { ct } from './layouts_ct.js'
import { optionControlsShown } from './layouts_controls.js'

export function addNavigationClickHandlers() {
    if (gameDone() === true) {
        const ack = document.querySelector('#ack')
        document.querySelector('#ack').onclick = () => {
            fetch('/acknowledge/' + GameInformation.ID).then(() => { window.location = '/' })
        }
        return
    }
    if (optionControlsShown === true) {
        document.querySelector('#concede').onclick = () => {
            fetch('/concede/' + GameInformation.ID).then(() => { window.location = '/' })
        }
    }
}

const concedeButton = ct('concede', '', true, true, '&#x2717;')
const ackButton = ct('ack', '', true, true, '&#x2713;') 

export function navigationLayout() {
    return '<div class="inline"></div>' + navigationButton() + '<div class="inline"></div>'
}

function navigationButton() {
    if (gameDone() === true) {
        return ackButton
    }
    if (optionControlsShown === true) {
        return concedeButton
    }
    return ''
}
