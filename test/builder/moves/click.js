import { board, setBoardAddPieceHandlers } from './board.js'
import { addMoveSet, moves, toIndicesFrom } from './cases.js'

export function addMovesAddHandler() {
    const e = document.querySelector('#moves')
    e.innerHTML = movesButton
    e.onclick = movesClick
}

let from
let to = []

const movesButton = '<button id="movesbutton" type="button">Add Moves</button>'

function movesClick() {
    // when the "Add Moves" button is clicked then the from square is picked, after which all 
    // available moves for that square are picked.
    addCancelButton()
    addFromClickHandlers()
    addMovesDefined()
}

const cancelButton = '<button id="cancel" type="button">Cancel</button>'

function addCancelButton() {
    const e = document.querySelector('#moves')
    e.innerHTML = cancelButton
    e.onclick = cancelClick
}

function addMovesDefined() {
    for (const s of moves) {
        document.querySelector('#s'+s.from).classList.add('fromDefined')
    }
}

function addFromClickHandlers() {
    for (let i = 0; i < 64; i++) {
        const s = document.querySelector('#s'+i)
        if (board[i] === undefined) {
            s.onclick = undefined
        } else {
            s.onclick = () => {
                from = i
                to = toIndicesFrom(i)

                document.querySelector('#s'+i).classList.add('fromSelected')

                addToClickHandlers()

                if (to.length !== 0) {
                    addNextFromButton()
                }
                // to onclick is replaced with the remove to onclick for already selected moves
                for (const addr of to) {
                    addRemoveToClickHandler(addr)
                    document.querySelector('#s'+addr).classList.add('toSelected')
                }
            }
        }
    }
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
        addMovesAddHandler()
        setBoardAddPieceHandlers()
    }
    removeNextFromButton()
}

function addToClickHandler(index) {
    document.querySelector('#s'+index).onclick = () => {
        to.push(index)
        addRemoveToClickHandler(index)
        document.querySelector('#s'+index).classList.add('toSelected')
        addNextFromButton()
    }
}

function addToClickHandlers() {
    for (let i = 0; i < 64; i++) {
        if (i === from) continue;
        addToClickHandler(i)
    }
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
    if (to.length > 0) {
        f.classList.add('fromDefined')
    } else {
        f.classList.remove('fromDefined')
    }
    f.onclick = undefined
    for (const addr of to) {
        const t = document.querySelector('#s'+addr).classList.remove('toSelected')
    }
    from = undefined
    to = []
    addFromClickHandlers()
    removeNextFromButton()
}

function addRemoveToClickHandler(index) {
    document.querySelector('#s'+index).onclick = () => {
        document.querySelector('#s'+index).classList.remove('toSelected')
        let spliceIndex
        for (let i = 0; i < to.length; i++) {
            if (to[i] !== index) {
                continue
            }
            spliceIndex = i
            break
        }
        if (spliceIndex === undefined) {
            throw new Error('to ' + index + ' not found in to array ' + to)
        }
        to.splice(spliceIndex, 1)
        addToClickHandler(index)
    }
}
