import { NoKind, Queen, Rook, Bishop, Knight } from '../pieceDefs.js'
import { layoutSelector } from '../layout.js'

import { promotion, players } from './layouts.js'
import { doMove } from './move.js'
import { writePlayersIndicator } from './players.js'

export let promoting = false

export function showPromotion(reverse = false) {
    promoting = true

    layoutSelector('#playernames', promotion)
    
    document.querySelector('#queenprom').addEventListener('click', promClick(Queen, reverse))
    document.querySelector('#rookprom').addEventListener('click', promClick(Rook, reverse))
    document.querySelector('#knightprom').addEventListener('click', promClick(Knight, reverse))
    document.querySelector('#bishopprom').addEventListener('click', promClick(Bishop, reverse))
}

function promClick(kind, reverse) {
    return () => { 
        promoting = false
        // a reverse promotion is followed by another move
        doMove(0, 0, kind, reverse)
        layoutSelector('#playernames', players)
        writePlayersIndicator()
    }
}
