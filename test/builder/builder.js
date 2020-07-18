// The layout of the position builder for the available moves test are added as the interface
// progresses.
// Initially only a select list of categories (named matching a tag used in the associated 
// filename) and a load button are shown.
// When loaded a select list of the cases in the file is shown with another load button, and 
// next to that is a text input and a new button for adding a new case instead of inspecting 
// an existing one.
// When load or new is pressed then the interface is shown with selection of the active 
// orientation, selection of the expected game state, selection of the previous move, a board 
// with all of the pieces shown (each with a kind, orientation, and if it's moved), followed
// by controls for adding or changing pieces, a button to start picking the previous move, and
// a button to start picking moves for a piece.
// At the bottom is a save button that writes out the category file and removes the position
// interface so that just the case list and new field are shown again.

import { writeCategorySelectList, addCategoryLoadButtonHandler } from './categories.js'

writeCategorySelectList()
addCategoryLoadButtonHandler()

export function selectValue(selector) {
    const s = document.querySelector(selector)
    return parseInt(s.options[s.selectedIndex].value)
}

export function selectValueString(selector) {
    const s = document.querySelector(selector)
    return s.options[s.selectedIndex].value
}
