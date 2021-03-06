import { Moves, Board } from '../game.js'

import { squareElement, boardAddressToIndex } from './board.js'
import { doMove } from './move.js'
import { moveablePiecesShown, unshowMoveablePieces, 
    previousMoveShown, unshowPreviousMove } from './moves.js'
import { hideOptions, optionControlsShown } from './layouts_controls.js'

const hoverClass = 'hover'
const moveClass = 'move'
const selectClass = 'select'
const selectMoveClass = 'selectmove'

let clicked = undefined

export function writeSquareClick(fromIndex, toIndices) {
    const s = squareElement(fromIndex)
    s.moves = toIndices

    s.mouseEnterFunc = event => {
        s.classList.add(hoverClass)
        for (const m of toIndices) {
            squareElement(m).classList.add(moveClass)
        }
    }
    s.addEventListener('mouseenter', s.mouseEnterFunc)

    s.mouseLeaveFunc = event => {
        s.classList.remove(hoverClass)
        for (const m of toIndices) {
            squareElement(m).classList.remove(moveClass)
        }
    }
    s.addEventListener('mouseleave', s.mouseLeaveFunc)

    s.moveClickFunc = () => {
        if (optionControlsShown === true) {
            hideOptions()
        }
        const removeMoveSelect = (fromSquare, toList) => {
            fromSquare.classList.remove(selectClass)
            for (const m of toList) {
                const ms = squareElement(m)
                ms.moveClickFunc = undefined
                ms.classList.remove(selectMoveClass)
            }
            s.mouseLeaveFunc()
        }

        // if this square is clicked then just deselect it
        if (clicked === fromIndex) {
            removeMoveSelect(squareElement(fromIndex), toIndices)
            clicked = undefined
            return
        }

        // if another square is already clicked then deselect it first
        if (clicked !== undefined) {
            const os = squareElement(clicked)
            removeMoveSelect(os, os.moves)
        }

        if (moveablePiecesShown === true) {
            unshowMoveablePieces()
        }

        if (previousMoveShown === true) {
            unshowPreviousMove()
        }

        clicked = fromIndex
        s.classList.add(selectClass)
        for (const m of toIndices) {
            const ms = squareElement(m)
            ms.classList.add(selectMoveClass)
            ms.moveClickFunc = () => {
                removeMoveSelect(squareElement(fromIndex), toIndices)
                clicked = undefined
                // handlers added back in doMove with new moves
                removeAllSquareEventHandlers()
                doMove(fromIndex, m)
            }
        }
    }
}

export function removeAllSquareEventHandlers() {
    for (let i = 0; i < Moves.length; i++) {
        const s = squareElement(boardAddressToIndex(Moves[i].from))
        s.removeEventListener('mouseenter', s.mouseEnterFunc)
        s.removeEventListener('mouseleave', s.mouseLeaveFunc)
        s.moveClickFunc = undefined
    }
}
