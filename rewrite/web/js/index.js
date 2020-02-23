import { addLayout, layout, layoutElement } from './layout.js'
import { indexLandscape, indexPortrait, indexSkinnyPortrait } from './indexLayouts.js'
import { layoutArmyPicker, layoutModeOptions } from './indexLayoutGenerate.js'
import { modes } from './indexDefinitions.js'

addLayout(1, indexSkinnyPortrait)
addLayout(1.8, indexPortrait)
addLayout(100, indexLandscape)

window.onload = layoutPage
window.onresize = layoutPage

function layoutPage() {
    layout()
    layoutArmyPicker()

    const modeClickFunc = (modeArgs) => {
        return () => {
            layoutModeOptions(modeArgs)
            layoutElement(document.querySelector('#modeoptions'))
        }
    }

    // TODO: read current mode from cookie
    modeClickFunc({mode: modes.COMPUTER})()

    document.querySelector('#computermode').onclick = modeClickFunc({mode: modes.COMPUTER})
    document.querySelector('#friendmode').onclick = modeClickFunc({mode: modes.FRIEND})
    document.querySelector('#timedmode').onclick = modeClickFunc({
        mode: modes.TIMED,
        wins: 50, // TODO: get win/loss/draw record from server
        losses: 22,
        draws: 90
    })
}


