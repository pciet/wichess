import { caseMaker } from './layouts.js'
import { initBoard, setBoardAddPieceHandlers, addBoardPieces, board } from './board.js'
import { categoryCase, pickedCategory } from './categories.js'
import { selectValueString } from './builder.js'
import { addMovesAddHandler } from './moves.js'

import { Pieces } from '../wichess/pieceDefs.js'
import { Orientation } from '../wichess/piece.js'
import { boardIndexToAddress, boardIndex } from '../wichess/game/board.js'
import { States } from '../wichess/game/state.js'

export function addCaseLoadButtonHandler() {
    document.querySelector('#loadcase').onclick = () => {
        loadCase(selectValueString('#cases'))
    }
}

export function addNewCaseButtonHandler() {
    document.querySelector('#newcase').onclick = () => {
        const t = document.querySelector('#newcasetext').value
        if (t === '') {
            throw new Error('no input text for title of new case')
        }
        loadCase(t)
    }
}

let testcase

function loadCase(name) {
    testcase = name
    const content = document.querySelector('#content')
    content.innerHTML = content.innerHTML + caseMaker
    addStateOptions()
    addPieceOptions()
    initBoard()
    addPreviousMoveHandler()
    addMovesAddHandler()
    addSaveHandler()

    const c = categoryCase(name)
    addBoardPieces(c.pos)
    setActivePlayer(c.active)
    setState(c.state)
    setPreviousMove(c.prev)
    setMoves(c.moves)
}

// The move sets (each is a from address and set of available moves) for a position are sorted in
// ascending order by the from address index. The moves for a from are also sorted in ascending
// address index order. This is so JSON file changes are minimized.
export let moves = []

function setMoves(m) {
    moves = []
    for (const s of m) {
        const t = []
        for (const tov of s.m) {
            t.push(boardIndex(tov.f, tov.r))
        }
        addMoveSet(boardIndex(s.f.f, s.f.r), t)
    }
}

// addMoveSet will replace a move set if from is the same.
export function addMoveSet(from, tos) {
    let index = 0 // where to insert the set
    let del = 0
    for (let i = 0; i < moves.length; i++) {
        if (moves[i].from < from) {
            index++
            continue
        }
        if (moves[i].from === from) {
            del = 1
        }
        break
    }
    moves.splice(index, del, {
        from: from,
        tos: tos
    })
}

export function toIndicesFrom(addr) {
    for (const set of moves) {
        if (set.from === addr) {
            return set.tos
        }
    }
    return []
}

let active

function setActivePlayer(orientation) {
    active = orientation
    document.querySelector('#orientation').value = orientation.toString()
}

function selectedActivePlayer() {
    const v = parseInt(document.querySelector('#orientation').value)
    return v
}

let state

function setState(s) {
    state = s
    document.querySelector('#state').value = s.toString()
}

function selectedState() {
    return parseInt(document.querySelector('#state').value)
}

let previousMove

function setPreviousMove(move) {
    previousMove = move
    setPreviousAddress('#from', move.f)
    setPreviousAddress('#to', move.t)
}

function setPreviousAddress(selector, addr) {
    document.querySelector(selector).innerHTML = addr.f + '-' + addr.r
}

// TODO: refactor nesting
function addPreviousMoveHandler() {
    const e = document.querySelector('#previous')
    e.onclick = () => {
        const from = document.querySelector('#from')
        from.classList.add('pickmove')
        for (let i = 0; i < 64; i++) {
            document.querySelector('#s'+i).onclick = () => {
                const addr = boardIndexToAddress(i)
                previousMove.f.f = addr.file
                previousMove.f.r = addr.rank
                setPreviousAddress('#from', previousMove.f)
                from.classList.remove('pickmove')

                const to = document.querySelector('#to')
                to.classList.add('pickmove')
                for (let i = 0; i < 64; i++) {
                    document.querySelector('#s'+i).onclick = () => {
                        const addr = boardIndexToAddress(i)
                        previousMove.t.f = addr.file
                        previousMove.t.r = addr.rank
                        setPreviousAddress('#to', previousMove.t)
                        to.classList.remove('pickmove')
                        setBoardAddPieceHandlers()
                    }
                }
            }
        }
    }
}

function addStateOptions() {
    const e = document.querySelector('#state')
    for (let i = 0; i < States.length; i++) {
        const o = document.createElement('option')
        o.value = i
        o.text = States[i].name
        e.add(o)
    }
}

function addPieceOptions() {
    const e = document.querySelector('#pieces')
    for (let i = 0; i < Pieces.length; i++) {
        const o = document.createElement('option')
        o.value = i
        o.text = Pieces[i].name
        e.add(o)
    }
}

function addSaveHandler() {
    document.querySelector('#savecase').onclick = () => {
        const pos = []
        for (let i = 0; i < board.length; i++) {
            if (board[i] === undefined) {
                continue
            }
            const p = board[i]
            const addr = boardIndexToAddress(i)
            pos.push({
                addr: {
                    f: addr.file,
                    r: addr.rank
                },
                k: p.kind,
                o: p.orientation,
                m: p.moved
            })
        }

        const mo = []
        for (const m of moves) {
            if (m.tos.length === 0) {
                continue
            }
            mo.push(m)
        }

        for (let i = 0; i < mo.length; i++) {
            mo[i].tos.sort((a, b) => {
                return a - b
            })
        }

        const movesOut = []
        for (let i = 0; i < mo.length; i++) {
            const m = mo[i]
            const fr = boardIndexToAddress(m.from)
            const ts = []
            for (let j = 0; j < m.tos.length; j++) {
                const addr = boardIndexToAddress(m.tos[j])
                ts[j] = {
                    f: addr.file,
                    r: addr.rank
                }
            }
            movesOut[i] = {
                f: {
                    f: fr.file,
                    r: fr.rank
                },
                m: ts
            }
        }

        const o = {
            case: testcase,
            active: selectedActivePlayer(),
            prev: previousMove,
            state: selectedState(),
            pos: pos,
            moves: movesOut
        }

        const str = JSON.stringify(o)

        fetch('/savecase?cat=' + pickedCategory, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: str
        }).then(() => {
            document.querySelector('#content').innerHTML = ''
        })
    }
}
