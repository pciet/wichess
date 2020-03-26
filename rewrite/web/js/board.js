export function BoardAddress(file, rank) {
    this.file = file
    this.rank = rank
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
