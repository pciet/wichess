import { layoutSelector } from '../layout.js'
import { piecePickImageName, CollectionPiece, NoSlot } from '../piece.js'
import { BasicKinds, Pawn, Knight,  Bishop, Rook, Queen, King } from '../pieceDefs.js'
import { Mode } from '../index.js'
import { randomInt } from '../random.js'

import { PageMode } from './mode.js'

let DefaultArmy = []

DefaultArmy[8] = new CollectionPiece(NoSlot, Rook)
DefaultArmy[15] = new CollectionPiece(NoSlot, Rook)
DefaultArmy[9] = new CollectionPiece(NoSlot, Knight)
DefaultArmy[14] = new CollectionPiece(NoSlot, Knight)
DefaultArmy[10] = new CollectionPiece(NoSlot, Bishop)
DefaultArmy[13] = new CollectionPiece(NoSlot, Bishop)
DefaultArmy[11] = new CollectionPiece(NoSlot, Queen)
DefaultArmy[12] = new CollectionPiece(NoSlot, King)

for (let i = 0; i < 8; i++) {
    DefaultArmy[i] = new CollectionPiece(NoSlot, Pawn)
}

let ComputerArmy = []
let PublicArmy = []

for (let i = 0; i < 16; i++) {
    ComputerArmy[i] = new CollectionPiece(DefaultArmy[i].slot, DefaultArmy[i].kind)
    PublicArmy[i] = new CollectionPiece(DefaultArmy[i].slot, DefaultArmy[i].kind)
}

export function addArmySelection(mode) {
    let a = ''
    for (let i = 0; i < 2; i++) {
        a += '<div>'
        for (let j = 0; j < 8; j++) {
            a += '<div class="inline armycell noselect" id="a'+((8*i)+j)+'"></div>'
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
    let t = '<img class="pieceimg noselect" src="/web/img/'
    t += piecePickImageName(kind) + '">'
    document.querySelector('#a'+index).innerHTML = t
}

// The army selection is only made of collection slot integers.
export function armySelectionJSON() {
    const j = []
    const s = armyForMode(Mode)
    for (let i = 0; i < 16; i++) {
        j[i] = s[i].slot
    }
    return JSON.stringify(j)
}

export function armyDefaultAt(index) {
    let k
    switch (Mode) {
    case PageMode.COMPUTER:
        ComputerArmy[index].kind = DefaultArmy[index].kind
        k = ComputerArmy[index].kind
        ComputerArmy[index].slot = NoSlot
        break
    case PageMode.PUBLIC:
        PublicArmy[index].kind = DefaultArmy[index].kind
        k = PublicArmy[index].kind
        PublicArmy[index].slot = NoSlot
        break
    }
    addArmyPicture(index, k)
    document.querySelector('#a'+index).classList.remove('pickedarmycell')
}

// A random basic kind for the special kind is replaced.
// randomArmyReplace returns the army index that was replaced.
export function randomArmyReplace(kind, collectionSlot) {
    // This slot is the army slot, not the collection slot.
    let slot
    switch (BasicKinds[kind]) {
    case Pawn:
        slot = randomInt(7)
        if (slotTaken(slot) === true) {
            slot++
            if (slot === 8) {
                slot = 0
            }
        }
        break
    case Rook:
        slot = notTakenSlot(8, 15)
        break
    case Knight:
        slot = notTakenSlot(9, 14)
        break
    case Bishop:
        slot = notTakenSlot(10, 13)
        break
    default:
        throw new Error("can't replace kind " + kind)
    }

    addArmyPicture(slot, kind)

    // The two random picks aren't in the collection. This means the army configuration can 
    // have special pieces that don't have a slot, which is a case the host looks for.
    armyForMode(Mode)[slot] = new CollectionPiece(collectionSlot, kind)

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
    if (kind !== BasicKinds[kind]) {
        return true
    }
    return false
}
