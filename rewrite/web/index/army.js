import { layoutSelector } from '../layout.js'
import { pieceLookImageName, IDPiece } from '../piece.js'
import { Pawn, Knight, Bishop, Rook, Queen, King } from '../pieceDefs.js'
import { Mode } from '../index.js'

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

    switch (mode) {
    case PageMode.COMPUTER:
        addArmyPictures(ComputerArmy)
        break
    case PageMode.PUBLIC:
        addArmyPictures(PublicArmy)
        break
    }
}

function addArmyPictures(selection) {
    for (let i = 0; i < 16; i++) {
        let t = '<img class="pieceimg" src="/web/img/'
        t += pieceLookImageName(selection[i].kind) + '">'
        document.querySelector('#a'+i).innerHTML = t
    }
}

export function armySelectionJSON() {
    const j = []
    let s
    switch (Mode) {
    case PageMode.COMPUTER:
        s = ComputerArmy
        break
    case PageMode.PUBLIC:
        s = PublicArmy
        break
    }
    for (let i = 0; i < 16; i++) {
        j[i] = s[i].id
    }
    return JSON.stringify(j)
}
