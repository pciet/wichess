import { Pieces, Pawn, Knight, Bishop, Rook, Queen, King } from '../pieceDefs.js'
import { NotInCollection, LeftPick, RightPick, collectionSelector } from '../collection.js'

export function ArmySlot(kind, collection = NotInCollection) {
    return {
        kind: kind,
        collection: collection
    }
}

export const Army = []

Army[8] = ArmySlot(Rook)
Army[15] = ArmySlot(Rook)
Army[9] = ArmySlot(Knight)
Army[14] = ArmySlot(Knight)
Army[10] = ArmySlot(Bishop)
Army[13] = ArmySlot(Bishop)
Army[11] = ArmySlot(Queen)
Army[12] = ArmySlot(King)

for (let i = 0; i < 8; i++) { Army[i] = ArmySlot(Pawn) }

export const DefaultArmy = []

for (let i = 0; i < 16; i++) { DefaultArmy[i] = Army[i].kind }

export function armyImage(index) { return '/web/img/pick_'+Pieces[Army[index].kind]+'.png' }

export function addArmySelection() {
    for (let i = 0; i < 16; i++) {
        const s = Army[i]

        if (s.collection === NotInCollection) {
            // then this army slot is not selected from the picks or collection
            continue
        } 

        const e = document.querySelector(collectionSelector(s.collection))
        e.armySlotIndex = i
        e.classList.add('used')

        const ae = document.querySelector('#a'+i)
        ae.classList.add('replaced')

        // the layout text isn't replaced when a piece is added, so the image has to be swapped here
        ae.src = armyImage(i)
        ae.classList.remove('piecehidden')
    }
}

export function deselectArmySlot(index) { 
    const ce = document.querySelector(collectionSelector(Army[index].collection))
    ce.classList.remove('used')
    ce.armySlotIndex = undefined

    Army[index] = ArmySlot(DefaultArmy[index]) 

    const e = document.querySelector('#a'+index)
    e.classList.add('piecehidden')
    e.classList.remove('replaced')
}

export function replaceArmySlot(index, armySlot) {
    if (Army[index].collection !== NotInCollection) {
        const ce = document.querySelector(collectionSelector(Army[index].collection))
        ce.classList.remove('used')
        ce.armySlotIndex = undefined
    }
    Army[index] = armySlot
    const e = document.querySelector('#a'+index)
    e.src = armyImage(index)
    e.classList.add('replaced')
    e.classList.remove('piecehidden')
}

export function armySelectionJSON() {
    const j = []
    for (let i = 0; i < 16; i++) { j[i] = Army[i].collection }
    return JSON.stringify(j)
}
