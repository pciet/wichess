import { interiorDimensions, layoutSelector } from '../layout.js'
import { setMode, Mode } from '../index.js'

import { addPublicMatches } from './public.js'
import { addCollection } from './collection.js'
import { addArmySelection } from './army.js'
import { sizePickCells, addPickImages,
    addPickClicks, selectExistingPicks } from './pick.js'
import { content, playButton } from './layouts.js'
import { addComputerPlayClick } from './computer.js'

export const PageMode = {
    // play against AI opponent with any pieces in your collection
    COMPUTER: 0, 
    // play against people with up to two pieces from a random pool
    PUBLIC: 1
}

export function pickMode(mode) {
    document.querySelector('#'+modeButtonName(mode)).classList.add('selected')
    setMode(mode)

    const sq = document.querySelector('#content')
    const dims = interiorDimensions(document.querySelector('#contentdiv'))
    if (dims.width > dims.height) {
        sq.style.width = dims.height + 'px'
        document.querySelector('#contentspacer').style.width = 
            ((dims.width - dims.height)/2) + 'px'
    } else {
        sq.style.height = dims.width + 'px'
        document.querySelector('#contentspacer').style.height = 
            ((dims.height - dims.width)/2) + 'px'
    }

    switch (mode) {
    case PageMode.PUBLIC:
        layoutSelector('#content', content + '<div id="publics"></div>')
        addPublicMatches()
        break
    case PageMode.COMPUTER:
        // TODO: collection feature
        //layoutSelector('#content', playButton + content + '<div id="collection"></div>')
        //addCollection()
        layoutSelector('#content', playButton + content + '<div></div>')
        addComputerPlayClick()
        break
    }

    // TODO: these picks and army things should be better organized

    addArmySelection(mode)
    sizePickCells()
    addPickImages()
    addPickClicks()
    selectExistingPicks()
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

export function savedPageMode() {
    const m = window.localStorage.getItem('mode')
    if (m === null) {
        savePageMode(PageMode.PUBLIC)
        return PageMode.PUBLIC
    }
    return parseInt(m)
}

export function savePageMode(m) {
    window.localStorage.setItem('mode', m)
}
