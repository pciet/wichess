import { addLayout, layout, scaleFont } from './layout.js'
import { Pieces } from './pieceDefs.js'

import { landscape } from './index/layouts.js'

import { addArmySelection, armySelectionJSON } from './index/army.js'
import { addPieceClicks, FloatingSelection } from './index/click.js'

addLayout(100, landscape)

export function layoutPage() {
    layout()
    scaleFont()

    setAllSquareDimensions('#leftpick', '.pick')
    const armyDim = setAllSquareDimensions('#a0', '.armycell')
    document.querySelector('#army0').style.height = armyDim + 'px'
    setAllSquareDimensions('#c0', '.collectioncell')

    document.querySelector('#details').onclick = () => {
        if (FloatingSelection === undefined) {
            return
        }
        window.open('/details?p=' + Pieces[FloatingSelection.kind])
    }

    addArmySelection()
    addPieceClicks()

    document.querySelector('#match').onclick = () => { 
        window.localStorage.setItem('army', armySelectionJSON())
        window.location = '/match'
    }
}

function setAllSquareDimensions(modelID, elementsClass) {
    const model = document.querySelector(modelID)
    const w = parseFloat(model.style.width)
    const h = parseFloat(model.style.height)

    const setSelectorAllStyleDims = (selector, styleValue) => {
        for (const e of document.querySelectorAll(selector)) {
            e.style.width = styleValue
            e.style.height = styleValue
        }
    }

    if (w > h) {
        setSelectorAllStyleDims(elementsClass, h + 'px')
        return h
    } else if (h > w) {
        setSelectorAllStyleDims(elementsClass, w + 'px')
        return w
    }
    return h
}

let resizeWait

window.onresize = () => {
    clearTimeout(resizeWait)
    resizeWait = setTimeout(layoutPage, 150)
}

window.onload = layoutPage
