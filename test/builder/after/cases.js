import { caseMaker } from './layouts.js'
import { initBoard, setBoardAddPieceHandlers, 
    addBoardPieces, addChangeBoardPieces, board, changeBoard } from './board.js'
import { categoryCase, pickedCategory } from './categories.js'
import { selectValueString } from './builder.js'
import { addStartHandler } from './start.js'
import { addDeleteChangeHandler } from './delete.js'

import { Pieces } from '../../wichess/pieceDefs.js'
import { Orientation } from '../../wichess/piece.js'
import { boardIndexToAddress, boardIndex } from '../../wichess/game/board.js'
import { States } from '../../wichess/game/state.js'

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
    addPieceOptions()
    initBoard()
    addMoveHandler()
    addPreviousMoveHandler()
    addStartHandler()
    addDeleteChangeHandler()
    addSaveHandler()

    const c = categoryCase(name)
    addBoardPieces(c.pos)
    addChangeBoardPieces(c.cha)
    setMove(c.mov)
    setPreviousMove(c.prev)
}

let move = {
    f: {},
    t: {}
}

function setMove(m) {
    if (m === undefined) {
        return
    }
    move = m
    setMoveAddress('#from', m.f)
    setMoveAddress('#to', m.t)
}

function setMoveAddress(selector, addr) {
    document.querySelector(selector).innerHTML = addr.f + '-' + addr.r
}

// TODO: refactor nesting
function addMoveHandler() {
    document.querySelector('#move').onclick = () => {
        const from = document.querySelector('#from')
        from.classList.add('pickmove')
        for (let i = 0; i < 64; i++) {
            document.querySelector('#s'+i).onclick = () => {
                const addr = boardIndexToAddress(i)
                move.f.f = addr.file
                move.f.r = addr.rank
                setMoveAddress('#from', move.f)
                from.classList.remove('pickmove')

                const to = document.querySelector('#to')
                to.classList.add('pickmove')
                for (let i = 0; i < 64; i++) {
                    document.querySelector('#s'+i).onclick = () => {
                        const addr = boardIndexToAddress(i)
                        move.t.f = addr.file
                        move.t.r = addr.rank
                        setMoveAddress('#to', move.t)
                        to.classList.remove('pickmove')
                        setBoardAddPieceHandlers()
                    }
                }
            }
        }
    }
}

let previousMove = {
    f: {},
    t: {}
}

function setPreviousMove(m) {
    if (m === undefined) {
        return
    }
    previousMove = m
    setMoveAddress('#prevfrom', m.f)
    setMoveAddress('#prevto', m.t)
}

// TODO: refactor nesting and duplicated handlers for previous move and move
function addPreviousMoveHandler() {
    document.querySelector('#prevmove').onclick = () => {
        const from = document.querySelector('#prevfrom')
        from.classList.add('pickmove')
        for (let i = 0; i < 64; i++) {
            document.querySelector('#s'+i).onclick = () => {
                const addr = boardIndexToAddress(i)
                previousMove.f.f = addr.file
                previousMove.f.r = addr.rank
                setMoveAddress('#prevfrom', move.f)
                from.classList.remove('pickmove')

                const to = document.querySelector('#prevto')
                to.classList.add('pickmove')
                for (let i = 0; i < 64; i++) {
                    document.querySelector('#s'+i).onclick = () => {
                        const addr = boardIndexToAddress(i)
                        previousMove.t.f = addr.file
                        previousMove.t.r = addr.rank
                        setMoveAddress('#prevto', move.t)
                        to.classList.remove('pickmove')
                        setBoardAddPieceHandlers()
                    }
                }
            }
        }
    }
}

function addPieceOptions() {
    const e = document.querySelector('#pieces')
    const ec = document.querySelector('#changepieces')
    const create = (index) => {
        const o = document.createElement('option')
        o.value = index
        o.text = Pieces[index]    
        return o
    }
    for (let i = 0; i < Pieces.length; i++) {
        e.add(create(i))
        ec.add(create(i))
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
                s: {
                    f: p.start.file,
                    r: p.start.rank
                }
            })
        }

        const cha = []
        for (let i = 0; i < changeBoard.length; i++) {
            const c = changeBoard[i]
            if (c === undefined) {
                continue
            }
            const addr = boardIndexToAddress(i)
            cha.push({
                a: {
                    f: addr.file,
                    r: addr.rank
                },
                p: {
                    k: c.kind,
                    o: c.orientation
                }
            })
        }

        const o = {
            case: testcase,
            mov: move,
            prev: previousMove,
            pos: pos,
            cha: cha
        }

        fetch('/after/save?cat=' + pickedCategory, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(o)
        }).then(() => {
            document.querySelector('#content').innerHTML = ''
        })
    }
}
