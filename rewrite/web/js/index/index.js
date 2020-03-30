import { addLayout, layout, layoutElement } from '../layout.js'
import { indexLandscape, indexPortrait, indexSkinnyPortrait } from './layouts.js'
import { layoutArmyPicker, layoutModeOptions } from './generate.js'

// These player record values are templated into a script of constants above the index.js script tag.
export const record = {
    wins: Wins,
    losses: Losses,
    draws: Draws
}

export const Modes = {
    COMPUTER: 0,
    FRIEND: 1,
    TIMED: 2
}

addLayout(1, indexSkinnyPortrait)
addLayout(1.8, indexPortrait)
addLayout(100, indexLandscape)

window.onload = layoutPage
window.onresize = layoutPage

const armySelection = []
for (let i = 0; i < 16; i++) {
    armySelection[i] = 0
}

function layoutPage() {
    layout()
    layoutArmyPicker()

    const modeClickFunc = (mode) => {
        return () => {
            layoutModeOptions(mode)
            layoutElement(document.querySelector('#modeoptions'))
            setModeClickHandlers(mode)
        }
    }

    // TODO: read current mode from cookie
    modeClickFunc(Modes.COMPUTER)()

    document.querySelector('#computermode').onclick = modeClickFunc(Modes.COMPUTER)
    document.querySelector('#friendmode').onclick = modeClickFunc(Modes.FRIEND)
    document.querySelector('#timedmode').onclick = modeClickFunc(Modes.TIMED)
}

function setModeClickHandlers(mode) {
    switch (mode) {
        case Modes.COMPUTER:
            document.querySelector('#computerbutton').onclick = () => {
                const req = new XMLHttpRequest()
                req.onload = (event) => {
                    if (event.target.status !== 200) {
                        throw new Error('server responded to POST /computer with status code ' + event.target.status)
                    }
                    window.location = '/computer'
                }
                req.open('POST', '/computer')
                req.setRequestHeader('Content-Type', 'application/json;charset=UTF-8')
                req.send(JSON.stringify(armySelection))
            }
            break
        case Modes.FRIEND:

            break

        case Modes.TIMED:

            break
        default:
            throw new Error('unknown mode ' + mode)
    }
}
