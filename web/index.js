import { addLayout, layout, scaleFont, setAllSquareDimensions } from './layout.js'
import { Pieces } from './pieceDefs.js'

import { landscape } from './index/layouts.js'

import { addArmySelection, armySelectionJSON } from './index/army.js'
import { addPieceClicks, DetailsKind } from './index/click.js'

addLayout(100, landscape)

export function layoutPage() {
    layout()
    scaleFont()

    const armyDim = setAllSquareDimensions('#ac0', '.armycell')
    document.querySelector('#army0').style.height = armyDim + 'px'
    setAllSquareDimensions('#leftpick', '.pick')
    setAllSquareDimensions('#c0', '.collectioncell')

    document.querySelector('#details').onclick = () => {
        window.open('/details?p=' + Pieces[DetailsKind])
    }

    document.querySelector('#rules').onclick = () => {
        window.open('/rules')
    }

    addArmySelection()
    addPieceClicks()

    document.querySelector('#match').onclick = () => { 
        window.localStorage.setItem('army', armySelectionJSON())
        window.location = '/match'
    }

    document.body.classList.add('visible')
}

let resizeWait

window.onresize = () => {
    clearTimeout(resizeWait)
    resizeWait = setTimeout(layoutPage, 150)
}

window.onload = layoutPage
