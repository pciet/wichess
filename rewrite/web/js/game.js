import { addLayout, layout, layoutElement, scaleFont } from './layout.js'
import { landscape, landscapeFloating, landscapeWideFloating, landscapeVeryWideFloating, square, portrait, unsupportedWindowDimension } from './gameLayouts.js'
import { boardGET } from './gameBoard.js'
import { movesGET } from './gameMoves.js'

boardGET(GameHeader.ID)
movesGET(GameHeader.ID, 1)

const lowerSquareRatio = 0.8
const upperSquareRatio = 1.5

addLayout(lowerSquareRatio, portrait)
addLayout(upperSquareRatio, square)
addLayout(1.8, landscape)
addLayout(2, landscapeFloating)
addLayout(3, landscapeWideFloating)
addLayout(3.4, landscapeVeryWideFloating)
addLayout(1000, unsupportedWindowDimension)

window.onload = layoutPage
window.onresize = layoutPage

function layoutPage() {
    setBoardDimension()
    layout()
    scaleFont()

    document.querySelector('#back').onclick = () => {
        window.location = '/'
    }
}

// TODO: add style instead of returning false
// TODO: allow setting multiple properties
function addCSSRuleProperty(cssSelector, property, value) {
    for (const s of document.styleSheets) {
        for (const r of s.cssRules) {
            if (r.selectorText !== cssSelector) {
                continue
            }
            r.style.setProperty(property, value)
            return true
        }
    }
    return false
}

function setBoardDimension() {
    let d
    if (window.innerWidth < window.innerHeight) {
        d = window.innerWidth
    } else {
        d = window.innerHeight
    }

    const r = window.innerWidth / window.innerHeight
    if ((r <= upperSquareRatio) && (r > lowerSquareRatio)) {
        d = d * 0.75
        let ok = addCSSRuleProperty('#boardrow', 'height', ((d / window.innerHeight) * 100) + '%')
        if (ok === false) {
            throw new Error('no #boardrow CSS rule')
        }
        ok = addCSSRuleProperty('#board', 'height', '100%')
        if (ok === false) {
            throw new Error('no #board CSS rule')
        }
        addCSSRuleProperty('#board', 'width', ((d / window.innerWidth) * 100) + '%')
        return
    }

    // if not the square layout then this:
    let ok = addCSSRuleProperty('#board', 'width', ((d / window.innerWidth) * 100) + '%')
    if (ok === false) {
        throw new Error('no #board CSS rule')
    }
    ok = addCSSRuleProperty('#board', 'height', ((d / window.innerHeight) * 100) + '%')
    if (ok === false) {
        throw new Error('no #board CSS rule')
    }
}
