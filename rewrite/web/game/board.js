export function BoardAddress(file, rank) {
    this.file = file
    this.rank = rank
}

export function boardIndex(file, rank) {
    return boardAddressToIndex(new BoardAddress(file, rank))
}

export function boardIndexToAddress(index) {
    return new BoardAddress(index%8, Math.floor(index/8))
}

export function boardAddressToIndex(address) {
    return address.file + (address.rank*8)
}

export function squareElement(atIndex) {
    return document.querySelector('#s'+atIndex.toString())
}

export function AddressedPiece(address, piece) {
    this.address = address
    this.piece = piece
}

export function chessBoard() {
    let b = ''
    for (let i = 0; i < 8; i++) {
        b += '<div class="row">'
        for (let j = 0; j < 8; j++) {
            b += '<div class="inline chesssquare" id="s'+(((7-i)*8)+j)+'"></div>'
        }
        b += '</div>'
    }
    return b
}
