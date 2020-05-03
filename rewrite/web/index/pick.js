import { interiorDimensions } from '../layout.js'
import { pieceLookImageName } from '../piece.js'

import { armyDefaultAt, randomArmyReplace } from './army.js'

export function sizePickCells() {
    const dims = interiorDimensions(document.querySelector('#picks'))
    document.querySelector('#leftpick').style.width = dims.height + 'px'
    document.querySelector('#rightpick').style.width = dims.height + 'px'
}

export function addPickImages() {
    let t = '<img class="pieceimg" id="leftpickimg" src="/web/img/'
    t += pieceLookImageName(LeftPiece) + '">'
    document.querySelector('#leftpick').innerHTML = t

    t = '<img class="pieceimg" id="rightpickimg" src="web/img/'
    t += pieceLookImageName(RightPiece) + '">'
    document.querySelector('#rightpick').innerHTML = t
}

export function addPickClicks() {
    document.querySelector('#leftpick').onclick = 
        pickClick('#leftpick', '#leftpickimg', LeftPiece)
    document.querySelector('#rightpick').onclick = 
        pickClick('#rightpick', '#rightpickimg', RightPiece)
}

function pickClick(selector, imgselector, kind) {
    return () => {
        document.querySelector(imgselector).classList.add('picked')
        document.querySelector(selector).onclick =
            pickUnclick(selector, imgselector, 
                randomArmyReplace(kind), kind)
    }
}

function pickUnclick(selector, imgselector, index, kind) {
    return () => {
        document.querySelector(imgselector).classList.remove('picked')
        armyDefaultAt(index)
        document.querySelector(selector).onclick =
            pickClick(selector, imgselector, kind)
    }
}
