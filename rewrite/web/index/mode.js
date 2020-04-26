import { interiorDimensions, layoutSelector } from '../layout.js'
import { setMode, Mode } from '../index.js'

import { addPublicMatches } from './public.js'
import { addCollection } from './collection.js'
import { addArmySelection } from './army.js'
import { sizePickCells } from './pick.js'
import { content, playButton } from './layouts.js'
import { addComputerPlayClick } from './computer.js'

export const PageMode = {
    COMPUTER: 0, // play against AI opponent with any pieces in your collection
    PUBLIC: 1 // play against people with up to two pieces from a random pool
}

export function pickMode(mode) {
    document.querySelector('#'+modeButtonName(mode)).classList.add('selected')
    setMode(mode)

    const sq = document.querySelector('#content')
    const dims = interiorDimensions(document.querySelector('#contentdiv'))
    if (dims.width > dims.height) {
        sq.style.width = dims.height + 'px'
        document.querySelector('#contentspacer').style.width = ((dims.width - dims.height)/2) + 'px'
    } else {
        sq.style.height = dims.width + 'px'
        document.querySelector('#contentspacer').style.height = ((dims.height - dims.width)/2) + 'px'
    }

    switch (mode) {
    case PageMode.PUBLIC:
        layoutSelector('#content', content + '<div id="publics"></div>')
        addPublicMatches()
        break
    case PageMode.COMPUTER:
        layoutSelector('#content', playButton + content + '<div id="collection"></div>')
        addCollection()
        addComputerPlayClick()
        break
    }

    addArmySelection(mode)
    sizePickCells()
}

export function modeClick(mode) {
    return () => {
        document.querySelector('#'+modeButtonName(Mode)).classList.remove('selected')
        document.querySelector('#'+modeButtonName(mode)).classList.add('selected')
        pickMode(mode)
    }
}

function modeButtonName(mode) {
    switch (mode) {
    case PageMode.COMPUTER:
        return 'computer'
    case PageMode.PUBLIC:
        return 'public'
    }
}
