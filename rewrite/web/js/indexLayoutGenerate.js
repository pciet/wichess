import { removeNewlines, elementInteriorDimensions } from './layout.js'
import { modes } from './indexDefinitions.js'
import { armyPicker } from './indexLayoutElements.js'
import { button, selectButton } from './button.js'
import { record } from './index.js'

export function layoutArmyPicker() {
    const parent = document.querySelector('#army')
    const dims = elementInteriorDimensions(parent)

    let squareDim, centeringHorizontal, centeringDivWidth, centeringDivHeight

    if (dims.width > dims.height) {
        centeringHorizontal = true
        squareDim = dims.height / 4
        if ((squareDim*8) > dims.width) { // TODO: this was a hack to make it work, are these calculations optimal?
            squareDim = dims.width / 8
        }
        centeringDivWidth = ((dims.width - (8*squareDim)) / 2) + 'px'
        centeringDivHeight = '100%'
    } else {
        centeringHorizontal = false
        squareDim = dims.width / 8
        if ((squareDim*4) > dims.height) {
            squareDim = dims.height / 4
        }
        centeringDivWidth = '100%'
        centeringDivHeight = ((dims.height - (4*squareDim)) / 2) + 'px'
    }

    const armySpace = document.querySelector('#army')
    armySpace.innerHTML = removeNewlines(armyPicker(centeringHorizontal))

    for (const e of armySpace.querySelectorAll('.armypickercenterer')) {
        e.style.width = centeringDivWidth
        e.style.height = centeringDivHeight
    }

    for (const e of armySpace.querySelectorAll('.pickercell')) {
        e.style.width = squareDim + 'px'
        e.style.height = squareDim + 'px'
    }
}

export function layoutModeOptions(mode) {
    const lmo = (html) => {
        let e = document.querySelector('#modeoptions')
        if (e === null) {
            throw new Error('no mode options element')
        }
        e.innerHTML = removeNewlines(html)
    }

    switch (mode) {
        case modes.COMPUTER:
            selectButton('#computermode')

            lmo(`<div class="computermodehorizontalcentering inline"></div>
                ` + button('modebutton', 'computerbutton', 'PLAY AGAINST COMPUTER', true) + `
                 <div class="computermodehorizontalcentering inline"></div>`)
            break

        case modes.FRIEND:
            selectButton('#friendmode')

            let friendHTML = '<div class="inline">'
            let friendHTML2 = '<div class="inline">'
            for (let i = 0; i < 3; i++) {
                // TODO: insert friend names
                friendHTML += button('modebutton', 'f'+i, '', false)
                friendHTML2 += button('modebutton', 'f'+(i+3), '', false)
            }
            friendHTML += '</div>' + friendHTML2 + '</div>'

            lmo(friendHTML)

            break

        case modes.TIMED:
            selectButton('#timedmode')

            let h = '<div id="record">+' + record.wins + ' -' + record.losses + ' =' + record.draws + '</div>' + `
                <div id="timedbuttons">
            ` + button('modebutton', 'matched15', '15 MINUTE', true) + button('modebutton', 'matched5', '5 MINUTE', true) + `
                </div>
            `

            lmo(h)

            break

        default:
            throw new Error('unknown mode')
    }
}
