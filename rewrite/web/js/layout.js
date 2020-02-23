import { layoutElement } from './layoutElement.js'
export { addLayout, layout, removeNewlines, elementInteriorDimensions, layoutElement }

const layouts = []

function addLayout(maxAspectRatio, htmlString) {
    layouts.push({
        maxRatio: maxAspectRatio,
        html: removeNewlines(htmlString)
    })
}

function layout() {
    document.body.innerHTML = pickLayout().html
    layoutElement(document.body)
}

function pickLayout() {
    if (layouts.length === 0) {
        throw new Error('no layouts')
    }

    const w = window.innerWidth
    const h = window.innerHeight

    document.body.style.width = w + 'px'
    document.body.style.height = h + 'px'

    const ratio = w / h
    let use = undefined

    for (const l of layouts) {
        if (l.maxRatio < ratio) {
            continue
        }
        if (use === undefined) {
            use = l
            continue
        }
        if (use.maxRatio === l.maxRatio) {
            throw new Error('duplicate maxAspectRatio ' + use.maxRatio)
        }
        if (l.maxRatio > use.maxRatio) {
            continue
        }
        use = l
    }
    if (use === undefined) {
        throw new Error('no layout for aspect ratio (width/height) ' + ratio)
    }

    return use
}

// stackoverflow.com/questions/10805125/how-to-remove-all-line-breaks-from-a-string
// stackoverflow.com/questions/12014441/remove-every-white-space-between-tags-using-javascript
function removeNewlines(html) {
    return html.replace(/(<(pre|script|style|textarea)[^]+?<\/\1)|(^|>)\s+|\s+(?=<|$)/g, '$1$3')
}

function elementInteriorDimensions(e) {
    const rect = e.getBoundingClientRect()

    // with 'box-sizing: border-box' getBoundingClientRect includes padding and border
    const s = window.getComputedStyle(e)

    const propValue = (name) => {
        const str = s.getPropertyValue(name)
        if (str === '') {
            return 0
        } else {
            return parseFloat(str)
        }
    }

    const p = propValue('padding')
    const b = propValue('border')

    return {
        width: rect.width - (2*(p+b)),
        height: rect.height - (2*(p+b))
    }
}
