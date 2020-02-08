export function generateIndexElements() {
    let pawnrowHTML = ''
    let pawnpickHTML = ''
    let royalrowHTML = ''
    let royalpickHTML = ''

    for (let i = 0; i < 8; i++) {
        pawnrowHTML += '<div class="inline armycell" id="pawn'+(i+8)+'"><img class="pieceimg" src="/img/wpawn_64.png"></div>'
        pawnpickHTML += '<div class="inline selectcell" id="pickpawn'+(i+8)+'"></div>'
        royalrowHTML += '<div class="inline armycell" id="royal'+i+'">'
        switch (i) {
            case 0:
            case 7:
                royalrowHTML += '<img class="pieceimg" src="/img/wrook_64.png">'
                break
            case 1:
            case 6:
                royalrowHTML += '<img class="pieceimg" src="/img/wknight_64.png">'
                break
            case 2:
            case 5:
                royalrowHTML += '<img class="pieceimg" src="/img/wbishop_64.png">'
                break
            case 3:
                royalrowHTML += '<img class="pieceimg" src="/img/wqueen_64.png">'
                break
            case 4:
                royalrowHTML += '<img class="pieceimg" src="/img/wking_64.png">'
                break
        }
        royalrowHTML += '</div>'
        royalpickHTML += '<div class="inline selectcell" id="pickroyal'+i+'"></div>'
    }

    let armyHTML = '<div class="armyrow" id="pawnrow">' + pawnrowHTML + '</div>'
    armyHTML +=  '<div class="armypick" id="pawnpick">' + pawnpickHTML + '</div>'
    armyHTML += '<div class="armyrow" id="royalrow">' + royalrowHTML + '</div>'
    armyHTML += '<div class="armypick" id="royalpick">' + royalpickHTML + '</div>'

    document.querySelector('#army').innerHTML = armyHTML

    for (let i = 0; i < 2; i++) {
        for (let j = 0; j < 8; j++) {
            if (i == 0) {
                const l = document.querySelector('#royal'+j).classList
                if (j % 2) {
                    l.add('odd')
                } else {
                    l.add('even')
                }
            } else {
                const l = document.querySelector('#pawn'+(j+8)).classList
                if (j % 2) {
                    l.add('even')
                } else {
                    l.add('odd')
                }
            }
        }
    }

    layoutModeOptions(currentMode)
}

export function layoutModeOptions(mode, wins, losses, draws) {
    const lmo = (html) => {
        let e = document.querySelector('#modeoptions')
        if (e === null) {
            throw new Error('no mode options element')
        }
        e.innerHTML = html
    }

    const sb = (selector) => {
        const e = document.querySelector(selector)
        if (e === null) {
            throw new Error('no element matching selector ' + selector)
        }
        e.classList.add('modebuttonselected')
    }

    switch (mode) {
        case modes.COMPUTER:
            sb('#computermode')
            lmo('<div class="optionbutton" id="computerbutton">Play AI</div>')
            break

        case modes.FRIEND:
            sb('#friendmode')

            let friendHTML = '<div class="inline">'
            let friendHTML2 = '<div class="inline">'
            for (let i = 0; i < 3; i++) {
                friendHTML += '<div id="f'+i+'" class="friendbutton"></div>'
                friendHTML2 += '<div id="f'+(i+3)+'" class="friendbutton"></div>'
            }
            friendHTML += '</div>' + friendHTML2 + '</div>'

            lmo(friendHTML)

            break

        case modes.TIMED:
            sb('#timedmode')

            let h = '<div id="record">+' + wins + ' -' + losses + ' =' + draws + '</div>' + `
                <div id="timedbuttons">
                    <div class="inline timedbutton" id="matched15">15 Minute</div>
                    <div class="inline timedbutton" id="matched5">5 Minute</div>
                </div>
            `

            lmo(h)

            break

        default:
            throw new Error('unknown mode')
    }
}
