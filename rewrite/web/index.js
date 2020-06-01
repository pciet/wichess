import { addLayout, layout, scaleFont } from './layout.js'

import { PageMode, pickMode, modeClick, savePageMode, savedPageMode } from './index/mode.js'
import { landscape } from './index/layouts.js'

export let Mode = savedPageMode()

export function setMode(mode) {
    savePageMode(mode)
    Mode = mode
}

addLayout(100, landscape)

export function layoutPage() {
    layout()
    pickMode(Mode)
    scaleFont()

    document.querySelector('#computer').onclick = modeClick(PageMode.COMPUTER)
    document.querySelector('#public').onclick = modeClick(PageMode.PUBLIC)
}

window.onload = layoutPage
window.onresize = layoutPage
