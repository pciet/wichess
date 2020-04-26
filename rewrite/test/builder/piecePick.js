import { pieceImageName } from '../wichess/piece.js'
import { addPiece } from './builder.js'

export function addBoardPiecePickHandlers() {
    for (let i = 0; i < 64; i++) {
        document.querySelector('#s'+i).onclick = () => {
            const pc = document.querySelector('#piecechoice')
            const pi = pc.options[pc.selectedIndex].value
            addPiece(i, pi)
            let img
            if (pi === '0') {
                img = ''
            } else {
                const orc = document.querySelector('#orientationchoice')
                const or = orc.options[orc.selectedIndex].value
                img = '<img src="img/'+pieceImageName(i, pi, or)+'">'
            }
            document.querySelector('#s'+i).innerHTML = img
        }
    }
}
