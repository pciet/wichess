import { interiorDimensions } from '../layout.js'
import { pieceLookImageName } from '../piece.js'
import { Mode } from '../index.js'

import { PageMode } from './mode.js'
import { armyDefaultAt, randomArmyReplace } from './army.js'

function Picks(left = undefined, right = undefined) {
    this.left = left
    this.right = right
}

const ComputerPicks = new Picks()
const PublicPicks = new Picks()

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

function picksForMode(mode) {
    switch (mode) {
    case PageMode.COMPUTER:
        return ComputerPicks
    case PageMode.PUBLIC:
        return PublicPicks
    }
    throw new Error('unknown mode ' + mode)
}

// selectExistingPicks is used when the interface is reset.
export function selectExistingPicks() {
    const p = picksForMode(Mode)
    if (p.right !== undefined) {
        addPick(p.right, RightPiece, true)
    }
    if (p.left !== undefined) {
        addPick(p.left, LeftPiece, false)
    }
}

function addPick(slot, kind, right) {
    let selector, imgselector
    if (right === true) {
        selector = '#rightpick'
        imgselector = '#rightpickimg'
    } else {
        selector = '#leftpick'
        imgselector = '#leftpickimg'
    }
    pickClick(selector, imgselector, kind, right, slot)()
}

export function addPickClicks() {
    document.querySelector('#leftpick').onclick = 
        pickClick('#leftpick', '#leftpickimg', LeftPiece, false)
    document.querySelector('#rightpick').onclick = 
        pickClick('#rightpick', '#rightpickimg', RightPiece, true)
}

function pickClick(selector, imgselector, kind, right, slot = undefined) {
    return () => {
        const e = document.querySelector(selector)
        e.classList.add('pickedcellpicked')
        document.querySelector(imgselector).classList.add('picked')
        let index
        if (slot === undefined) {
            index = randomArmyReplace(kind)
        } else {
            index = slot
        }
        const p = picksForMode(Mode)
        if (right === true) {
            p.right = index
        } else {
            p.left = index
        }
        const army = document.querySelector('#a'+index)
        army.classList.add('pickedarmycell')

        const click = pickUnclick(selector, imgselector, index, kind, right)
        army.onclick = click
        e.onclick = click
    }
}

function pickUnclick(selector, imgselector, index, kind, right) {
    return () => {
        document.querySelector(imgselector).classList.remove('picked')
        armyDefaultAt(index)
        const e = document.querySelector(selector)
        e.classList.remove('pickedcellpicked')
        e.onclick = pickClick(selector, imgselector, kind, right)
        document.querySelector('#a'+index).onclick = undefined
        const p = picksForMode(Mode)
        if (right === true) {
            p.right = undefined
        } else {
            p.left = undefined
        }
    }
}
