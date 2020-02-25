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
