import { NoKind, BasicKinds, King } from '../pieceDefs.js'
import { LeftPick, RightPick, NotInCollection, collectionSelector } from '../collection.js'

import { DefaultArmy, Army, ArmySlot, deselectArmySlot, replaceArmySlot } from './army.js'

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

let FloatingSelection
export let DetailsKind = King

function armyClick(index) {
    return () => {
        if (FloatingSelection === undefined) {
            DetailsKind = Army[index].kind
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
        undarkenCollection()
        e.classList.remove('selected')
        e.classList.add('used')
        e.armySlotIndex = index
        FloatingSelection = undefined
        DetailsKind = Army[index].kind
    }
}

function collectionClick(sourceElementID, armySlot) {
    return () => {
        DetailsKind = armySlot.kind
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
                undarkenCollection()
                return
            }
            // another piece was floating and is deselected before this one is selected to
            // be floating
            document.querySelector(
                collectionSelector(FloatingSelection.collection)).classList.remove('selected')
        }

        FloatingSelection = armySlot
        e.classList.add('selected')
        darkenCollection(e, armySlot.kind)
    }
}

function darkenCollection(excludedElement, selectedKind) {
    const basic = BasicKinds[selectedKind]
    for (let i = 0; i < 16; i++) {
        if (DefaultArmy[i] === basic) {
            continue
        }
        document.querySelector('#ab'+i).classList.add('darkened')
        document.querySelector('#a'+i).classList.add('darkened')
    }
    for (let i = 0; i < 21; i++) {
        document.querySelector('#c'+i).classList.add('darkened')
    }
    document.querySelector('#leftpick').classList.add('darkened')
    document.querySelector('#rightpick').classList.add('darkened')
    excludedElement.classList.remove('darkened')
}

function undarkenCollection() {
    for (let i = 0; i < 16; i++) {
        document.querySelector('#ab'+i).classList.remove('darkened')
        document.querySelector('#a'+i).classList.remove('darkened')
    }
    for (let i = 0; i < 21; i++) {
        document.querySelector('#c'+i).classList.remove('darkened')
    }
    document.querySelector('#leftpick').classList.remove('darkened')
    document.querySelector('#rightpick').classList.remove('darkened')
}
