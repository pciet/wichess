import { updateBoard, parseBoardUpdate } from './board_update.js'

export function doMove(fromIndex, toIndex) {
    fetch('/move/'+GameInformation.ID, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({f: fromIndex, t: toIndex})
    }).then(r => r.json()).then(diff => {
        updateBoard(parseBoardUpdate(diff))
    })
}
