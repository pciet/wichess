export function doMove(fromIndex, toIndex) {
    fetch('/move/'+GameInformation.ID, {
        method: 'POST',
        credentials: 'same-origin',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(new MoveRequest(fromIndex, toIndex))
    }).then(r => r.json()).then(diff => {
        console.log(diff)
    })
}

function MoveRequest(fromIndex, toIndex) {
    this.fromIndex = fromIndex
    this.toIndex = toIndex
    return this
}
