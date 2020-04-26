import { addBoardPiecePickHandlers } from './piecePick.js'

export function prevFromClick() {
    const b = document.querySelector('#prevfrom')
    b.innerHTML = 'Pick Prev From'
    b.onclick = () => {
        b.innerHTML = 'Prev From'
        b.onclick = prevFromClick
        b.classList.remove('pickingbutton')
        addBoardPiecePickHandlers()
    }
    b.classList.add('pickingbutton')
    addPrevFromSquareClicks()
}

export function prevToClick() {

}
