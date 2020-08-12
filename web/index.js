import { addLayout, layout, scaleFont, setAllSquareDimensions } from './layout.js'
import { Pieces } from './pieceDefs.js'

import { landscape, square, unsupported } from './index/layouts.js'

import { addArmySelection, armySelectionJSON } from './index/army.js'
import { addPieceClicks, DetailsKind } from './index/click.js'

addLayout(1000, unsupported)
addLayout(2.5, landscape)
addLayout(1.29, square)
addLayout(0.4, unsupported)

export function layoutPage() {
    layout()
    scaleFont()

    if (document.querySelector('#ac0') === null) {
        // unsupported aspect ratio
        document.body.classList.add('visible')    
        return
    }

    const armyDim = setAllSquareDimensions('#ac0', '.armycell')
    document.querySelector('#army0').style.height = armyDim + 'px'
    setAllSquareDimensions('#leftpick', '.pick')
    setAllSquareDimensions('#c0', '.collectioncell')

    // these buttons have varying selectors for CSS styling

    let de = document.querySelector('#details')
    if (de === null) {
        de = document.querySelector('#squaredetails')
    }
    de.onclick = () => { window.open('/details?p=' + Pieces[DetailsKind]) }

    let re = document.querySelector('#rules')
    if (re === null) {
        re = document.querySelector('#squarerules')
    }
    re.onclick = () => { window.open('/rules') }

    addArmySelection()
    addPieceClicks()

    let me = document.querySelector('#match')
    if (me === null) {
        me = document.querySelector('#squarematch')
    }
    me.onclick = () => {
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
