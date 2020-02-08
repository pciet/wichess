export { layoutPage }

import { addLayout, layout } from './layout.js'
import { indexLandscape, indexPortrait } from './layouts.js'
import { generateIndexElements } from './indexLayout.js'

addLayout(1, indexLandscape, generateIndexElements)
addLayout(100, indexPortrait, generateIndexElements)

window.onload = layoutPage
window.onresize = layoutPage

// TODO: request the record from an HTTP endpoint instead of getting it with templating
// TODO: these need to be in the HTML file to be templated in
const wins = 50
const losses = 22 
const draws = 90

function layoutPage() {
    layout()
    // TODO: setup click handlers
}


