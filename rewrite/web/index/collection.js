import { layoutSelector, interiorDimensions } from '../layout.js'

export function addCollection() {
    let a = ''
    for (let i = 0; i < 3; i++) {
        a += '<div>'
        for (let j = 0; j < 6; j++) {
            a += '<div class="inline collectioncell" id="c'+((6*(2-i))+j)+'"></div>'
        }
        a += '</div>'
    }
    layoutSelector('#collection', a)

    const dim = interiorDimensions(document.querySelector('#c0')).height + 'px'
    for (const e of document.querySelectorAll('.collectioncell')) {
        e.style.width = dim
    }
}
