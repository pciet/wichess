export function movesGET(gameID, turn) {
    const r = new XMLHttpRequest()
    r.onload = readMovesGETResponse
    r.open('GET', '/moves/' + gameID + '?turn=' + turn)
    r.send()
}

function readMovesGETResponse(r) {
    if (r.target.status !== 200) {
        throw new Error('/moves/[id] GET failed with status code ' + r.target.status)
    }

    const m = JSON.parse(r.target.response)

    console.log(m)

    console.log(board)
}
