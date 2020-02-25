export function startBoardGET(gameID) {
    const r = new XMLHttpRequest()
    r.onload = readBoardGETResponse
    r.open('GET', '/games/' + gameID)
    r.send()
}

function readBoardGETResponse(r) {
    if (r.status !== 200) {
        throw new Error('/games/ GET failed with status code ' + r.status)
    }
    console.log(r)
}
