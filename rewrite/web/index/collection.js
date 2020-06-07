import { layoutSelector, interiorDimensions } from '../layout.js'
import { NoKind } from '../pieceDefs.js'
import { pieceLookImageName } from '../piece.js'

const CollectionCount = 21

export function addCollection() {
    let a = ''
    for (let i = 0; i < 3; i++) {
        a += '<div>'
        for (let j = 0; j < 7; j++) {
            a += '<div class="inline collectioncell noselect" id="c'+((7*i)+j)+'"></div>'
        }
        a += '</div>'
    }
    layoutSelector('#collection', a)

    const dim = interiorDimensions(document.querySelector('#c0')).height + 'px'
    for (const e of document.querySelectorAll('.collectioncell')) {
        e.style.width = dim
    }

    for (let i = 0; i < CollectionCount; i++) {
        if (Collection[i] == NoKind) {
            continue
        }
        document.querySelector('#c'+i).innerHTML = '<img class="pieceimg noselect" src="/web/img/' +
            pieceLookImageName(Collection[i]) + '">'
    }
}
