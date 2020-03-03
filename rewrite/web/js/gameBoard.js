const board = [64]

export function boardGET(gameID) {
    const r = new XMLHttpRequest()
    r.onload = readBoardGETResponse
    r.open('GET', '/games/' + gameID)
    r.send()
}

function readBoardGETResponse(r) {
    if (r.target.status !== 200) {
        throw new Error('/games/[id] GET failed with status code ' + r.target.status)
    }

    const b = JSON.parse(r.target.response)

    for (let i = 0; i < 64; i++) {
        board[i] = {
            piece: b.board[i],
            id: 0
        }
    }

    for (let p of b.pieceIdentifiers) {
        board[squareIndex(p.address)].id = p.id
    }

    console.log(board)
}

function squareIndex(address) {
    return address.file + (address.rank * 8)
}
