import { addLayout, layout } from './layout.js'

import { PageMode, pickMode, modeClick } from './index/mode.js'
import { landscape } from './index/layouts.js'

export let Mode = PageMode.PUBLIC

export function setMode(mode) { Mode = mode }

addLayout(100, landscape)

function layoutPage() {
    layout()
    pickMode(Mode)

    document.querySelector('#computer').onclick = modeClick(PageMode.COMPUTER)
    document.querySelector('#public').onclick = modeClick(PageMode.PUBLIC)
}

window.onload = layoutPage
window.onresize = layoutPage
