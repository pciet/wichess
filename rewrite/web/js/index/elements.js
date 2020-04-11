export function armyPicker(centeringHorizontal) {
    let centeringDiv = '<div class="armypickercenterer'
    if (centeringHorizontal === true) {
        centeringDiv += ' inline'
    }
    centeringDiv += '"></div>'

    let pawnrow = '<div class="armyrow" id="pawnrow">'
    let royalrow = '<div class="armyrow" id="royalrow">'
    let pawnpick = '<div class="armypick" id="pawnpick">'
    let royalpick = '<div class="armypick" id="royalpick">'

    for (let i = 0; i < 8; i++) {
        pawnrow += '<div class="inline pickercell" id="pawn'+(i+8)+'"><img class="pieceimg" src="/img/look_pawn.png"></div>'

        royalrow += '<div class="inline pickercell" id="royal'+i+'">'
        switch (i) {
        case 0:
        case 7:
            royalrow += '<img class="pieceimg" src="/img/look_rook.png">'
            break
        case 1:
        case 6:
            royalrow += '<img class="pieceimg" src="/img/look_knight.png">'
            break
        case 2:
        case 5:
            royalrow += '<img class="pieceimg" src="/img/look_bishop.png">'
            break
        case 3:
            royalrow += '<img class="pieceimg" src="/img/look_queen.png">'
            break
        case 4:
            royalrow += '<img class="pieceimg" src="/img/look_king.png">'
            break
        }
        royalrow += '</div>'

        pawnpick += '<div class="inline pickercell" id="pickpawn'+(i+8)+'"></div>'
        royalpick += '<div class="inline pickercell" id="pickroyal'+i+'"></div>'
    }

    pawnrow += '</div>'
    royalrow += '</div>'
    pawnpick += '</div>'
    royalpick += '</div>'

    let picker = centeringDiv + '<div id="armypicker"'
    if (centeringHorizontal === true) {
        picker += ' class="inline"'
    }
    picker += '>' + pawnrow + royalrow + pawnpick + royalpick + '</div>' + centeringDiv

    return picker
}
