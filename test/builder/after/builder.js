// The after move test builder JavaScript was copied from the moves test builder.
// Ideally anything that could be shared between these two tests would be instead of copied.

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
