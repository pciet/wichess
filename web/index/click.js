import { NoKind, BasicKinds } from '../pieceDefs.js'
import { LeftPick, RightPick, NotInCollection, collectionSelector } from '../collection.js'

import { Army, ArmySlot, deselectArmySlot, replaceArmySlot } from './army.js'

export function addPieceClicks() {
    document.querySelector('#leftpick').onclick = 
        collectionClick('#leftpick', ArmySlot(window.LeftPiece, LeftPick))

    document.querySelector('#rightpick').onclick = 
            collectionClick('#rightpick', ArmySlot(window.RightPiece, RightPick))

    for (let i = 0; i < window.Collection.length; i++) {
        if (window.Collection[i] === NoKind) {
            continue
        }
        const id = '#c'+i
        document.querySelector(id).onclick = 
            collectionClick(id, ArmySlot(window.Collection[i], i+1))
    }

    for (let i = 0; i < 16; i++) {
        document.querySelector('#ac'+i).onclick = armyClick(i)
    }
}

export let FloatingSelection

function armyClick(index) {
    return () => {
        if (FloatingSelection === undefined) {
            if (Army[index].collection !== NotInCollection) {
                deselectArmySlot(index)
            }
            return
        }
        if (BasicKinds[FloatingSelection.kind] !== BasicKinds[Army[index].kind]) {
            return
        }
        replaceArmySlot(index, FloatingSelection)
        const e = document.querySelector(collectionSelector(FloatingSelection.collection))
        e.classList.remove('selected')
        e.classList.add('used')
        e.armySlotIndex = index
        FloatingSelection = undefined
    }
}

function collectionClick(sourceElementID, armySlot) {
    return () => {
        const e = document.querySelector(sourceElementID)

        if (e.armySlotIndex !== undefined) {
            // this collection or pick slot was already clicked and added to the army
            deselectArmySlot(e.armySlotIndex)
            e.armySlotIndex = undefined
            e.classList.remove('selected')
            return
        }

        if (FloatingSelection !== undefined) {
            if (FloatingSelection.collection === armySlot.collection) {
                // this slot was clicked but not added to the army yet
                FloatingSelection = undefined
                e.classList.remove('selected')
                return
            }
            // another piece was floating and is deselected before this one is selected to
            // be floating
            document.querySelector(
                collectionSelector(FloatingSelection.collection)).classList.remove('selected')
        }

        FloatingSelection = armySlot
        e.classList.add('selected')
    }
}
