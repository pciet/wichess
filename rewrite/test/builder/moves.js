import { addBoardPiecePickHandlers } from './piecePick.js'
import { addMoveSet, board } from './builder.js'

// When the "Add Moves" button is pressed then a series of button and click
// states follows. The result is a move set, a from address and set of
// to addresses.

export function addMovesButtons() {
    const e = document.querySelector('#moves')
    e.innerHTML = movesButton
    e.onclick = movesClick
}

let from
let to = []

const movesButton = '<button id="movesbutton" type="button">Add Moves</button>'

function movesClick() {
    // when the "Add Moves" button is clicked then the from square
    // is picked, after which all available moves for that square are picked.
    addCancelButton()
    addFromClickHandlers()
}

const nextButton = '<button id="nextbutton" type="button">Next</button>'

function addNextFromButton() {
    const e = document.querySelector('#next')
    e.innerHTML = nextButton
    e.onclick = nextClick
}

function removeNextFromButton() {
    document.querySelector('#next').innerHTML = ''
}

function nextClick() {
    addMoveSet(from, to)
    const f = document.querySelector('#s'+from)
    f.classList.remove('fromSelected')
    f.classList.add('fromDefined')
    f.onclick = undefined
    for (const addr of to) {
        const t = document.querySelector('#s'+addr).classList.remove('toSelected')
    }
    from = undefined
    to = []
    removeNextFromButton()
}

const cancelButton = '<button id="cancel" type="button">Cancel</button>'

function addCancelButton() {
    const e = document.querySelector('#moves')
    e.innerHTML = cancelButton
    e.onclick = cancelClick
}

function cancelClick() {
    if (from !== undefined) {
        if (to.length > 0) {
            for (const addr of to) {
                document.querySelector('#s'+addr).classList.remove('toSelected')
                addToClickHandler(addr)
            }
            to = []
        } else {
            document.querySelector('#s'+from).classList.remove('fromSelected')
            from = undefined
        }
    } else {
        addMovesButtons()
        addBoardPiecePickHandlers()
    }
    removeNextFromButton()
}

function addFromClickHandlers() {
    for (let i = 0; i < 64; i++) {
        const s = document.querySelector('#s'+i)
        if (board[i] === undefined) {
            s.onclick = undefined
        } else {
            s.onclick = () => {
                from = i
                document.querySelector('#s'+i).classList.add('fromSelected')
                addToClickHandlers()
            }
        }
    }
}

function addToClickHandler(index) {
    document.querySelector('#s'+index).onclick = () => {
        to.push(index)
        addRemoveToClickHandler(index)
        document.querySelector('#s'+index).classList.add('toSelected')
        addNextFromButton()
    }
}

function addRemoveToClickHandler(index) {
    document.querySelector('#s'+index).onclick = () => {
        document.querySelector('#s'+index).classList.remove('toSelected')
        to.splice(index, 1)
        addToClickHandler(index)
    }
}

function addToClickHandlers() {
    for (let i = 0; i < 64; i++) {
        if (i === from) continue;
        addToClickHandler(i)
    }
}
