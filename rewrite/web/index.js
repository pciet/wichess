import { addLayout, layout, scaleFont } from './layout.js'

import { PageMode, pickMode, modeClick, savePageMode, savedPageMode } from './index/mode.js'
import { landscape } from './index/layouts.js'

export let Mode = savedPageMode()

export function setMode(mode) {
    savePageMode(mode)
    Mode = mode
}

addLayout(100, landscape)

let disableLayout = false

export function layoutPage() {
    // On some web browsers, like Chrome on Android, onresize is caused by the keyboard showing
    // which breaks this and makes the webpage unusable. Layout is disabled as a workaround.
    if (disableLayout === true) {
        return
    }

    layout()
    pickMode(Mode)
    scaleFont()

    if (Mode === PageMode.PUBLIC) {
        const opp = document.querySelector('#opponent')
        opp.onclick = () => { disableLayout = true }
        opp.addEventListener('blur', () => { disableLayout = false })
    }

    document.querySelector('#computer').onclick = modeClick(PageMode.COMPUTER)
    document.querySelector('#public').onclick = modeClick(PageMode.PUBLIC)
}

window.onload = layoutPage
window.onresize = layoutPage
