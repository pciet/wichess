import { changeBoard, addChangeBoardPieceHandlers } from './board.js'

export function addDeleteChangeHandler() {
    document.querySelector('#deletechange').onclick = () => {
        addDeleteChangeSquareHandlers()
        const e = document.querySelector('#deletechange')
        e.innerText = 'Cancel'
        e.onclick = cancelDelete
    }
}

function addDeleteChangeSquareHandlers() {
    for (let i = 0; i < 64; i++) {
        document.querySelector('#c'+i).onclick = () => {
            changeBoard[i] = undefined
            document.querySelector('#c'+i).innerHTML = ''
            cancelDelete()
        }
    }
}

function cancelDelete() {
    const e = document.querySelector('#deletechange')
    e.innerText = 'Delete Square'
    addChangeBoardPieceHandlers()
    addDeleteChangeHandler()
}
