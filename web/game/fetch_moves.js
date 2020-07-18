import { BoardAddress } from './board.js'
import { fetchedMoves } from './moves.js'

function MovesResponse(moveSets, state) {
    this.moveSets = moveSets
    this.state = state
}

function MoveSet(from, toAddresses) {
    this.from = from
    this.tos = toAddresses
}

export function fetchMoves() {
    fetchMovesPromise(GameInformation.ID).then(m => { fetchedMoves(m) })
}

export function fetchMovesPromise(gameID, turnNumber) {
    return fetch('/moves/'+gameID+'?turn='+turnNumber).then(r => r.json()).then(r => {
        if (r.m === undefined) {
            return new MovesResponse(undefined, r.s)
        }
        const moveSets = []
        for (const ms of r.m) {
            const tos = []
            for (const to of ms.m) {
                if ((to.f < 0) || (to.f > 7) || (to.r < 0) || (to.r > 7)) {
                    throw new Error('host calculated bad move, from ' + ms.f.f + '-' + ms.f.r +
                        ' to ' + to.f + '-' + to.r)
                }
                tos.push(new BoardAddress(to.f, to.r))
            }
            moveSets.push(new MoveSet(new BoardAddress(ms.f.f, ms.f.r), tos))
        }
        return new MovesResponse(moveSets, r.s)
    })
}
