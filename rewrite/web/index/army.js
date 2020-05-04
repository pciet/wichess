import { layoutSelector } from '../layout.js'
import { piecePickImageName, IDPiece } from '../piece.js'
import { Pieces, Pawn, Knight, 
    Bishop, Rook, Queen, King } from '../pieceDefs.js'
import { Mode } from '../index.js'
import { randomInt } from '../random.js'

import { PageMode } from './mode.js'

let DefaultArmy = []

DefaultArmy[0] = new IDPiece(0, Rook)
DefaultArmy[7] = new IDPiece(0, Rook)
DefaultArmy[1] = new IDPiece(0, Knight)
DefaultArmy[6] = new IDPiece(0, Knight)
DefaultArmy[2] = new IDPiece(0, Bishop)
DefaultArmy[5] = new IDPiece(0, Bishop)
DefaultArmy[3] = new IDPiece(0, Queen)
DefaultArmy[4] = new IDPiece(0, King)

for (let i = 8; i < 16; i++) {
    DefaultArmy[i] = new IDPiece(0, Pawn)
}

let ComputerArmy = []
let PublicArmy = []

for (let i = 0; i < 16; i++) {
    ComputerArmy[i] = new IDPiece(DefaultArmy[i].id, DefaultArmy[i].kind)
    PublicArmy[i] = new IDPiece(DefaultArmy[i].id, DefaultArmy[i].kind)
}

export function addArmySelection(mode) {
    let a = ''
    for (let i = 0; i < 2; i++) {
        a += '<div>'
        for (let j = 0; j < 8; j++) {
            a += '<div class="inline armycell" id="a'+((8*(1-i))+j)+'"></div>'
        }
        a += '</div>'
    }
    layoutSelector('#army', a)

    const army = armyForMode(Mode)
    addArmyPictures(army)
}

function addArmyPictures(selection) {
    for (let i = 0; i < 16; i++) {
        addArmyPicture(i, selection[i].kind)
    }
}

function addArmyPicture(index, kind) {
    let t = '<img class="pieceimg" src="/web/img/'
    t += piecePickImageName(kind) + '">'
    document.querySelector('#a'+index).innerHTML = t
}

export function armySelectionJSON() {
    const j = []
    const s = armyForMode(Mode)
    for (let i = 0; i < 16; i++) {
        j[i] = s[i].id
    }
    return JSON.stringify(j)
}

export function armyDefaultAt(index) {
    let k
    switch (Mode) {
    case PageMode.COMPUTER:
        ComputerArmy[index].kind = DefaultArmy[index].kind
        k = ComputerArmy[index].kind
        ComputerArmy[index].id = 0
        break
    case PageMode.PUBLIC:
        PublicArmy[index].kind = DefaultArmy[index].kind
        k = PublicArmy[index].kind
        PublicArmy[index].id = 0
        break
    }
    addArmyPicture(index, k)
    document.querySelector('#a'+index).classList.remove('pickedarmycell')
}

// randomArmyReplace returns the index that was replaced.
// A random basic kind slot for the special kind is replaced.
export function randomArmyReplace(kind) {
    let slot
    switch (Pieces[kind].basicKind) {
    case Pawn:
        slot = randomInt(7) + 8
        if (slotTaken(slot) === true) {
            slot++
            if (slot === 16) {
                slot = 8
            }
        }
        break
    case Rook:
        slot = notTakenSlot(0, 7)
        break
    case Knight:
        slot = notTakenSlot(1, 6)
        break
    case Bishop:
        slot = notTakenSlot(2, 5)
        break
    default:
        throw new Error("can't replace kind " + kind)
    }

    addArmyPicture(slot, kind)

    // TODO: special pieces need to be identified
    armyForMode(Mode)[slot] = new IDPiece(0, kind)

    return slot
}

function armyForMode(m) {
    switch (m) {
    case PageMode.COMPUTER:
        return ComputerArmy
    case PageMode.PUBLIC:
        return PublicArmy
    }
    throw new Error('unknown page mode ' + m)
}

function notTakenSlot(left, right) {
    if (slotTaken(left) === true) {
        return right
    } else if (slotTaken(right) === true) {
        return left
    } else {
        if (randomInt(1) === 0) {
            return left
        } else {
            return right
        }
    }
}

function slotTaken(slot) {
    const a = armyForMode(Mode)
    const kind = a[slot].kind
    if (kind !== Pieces[kind].basicKind) {
        return true
    }
    return false
}
