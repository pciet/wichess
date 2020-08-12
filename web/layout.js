import { layoutElement } from './layoutElement.js'
export { addLayout, removeLayout, removeAllLayouts, layout, layoutSelector, setAllSquareDimensions,
    removeNewlines, interiorDimensions, layoutElement, scaleFont, addCSSRuleProperty }

let layouts = []

function addLayout(maxAspectRatio, htmlString) {
    layouts.push({
        maxRatio: maxAspectRatio,
        html: removeNewlines(htmlString)
    })
}

function removeLayout(maxAspectRatio) {
    for (let i = 0; i < layouts.length; i++) {
        if (layouts[i].maxRatio !== maxAspectRatio) {
            continue
        }
        layouts.splice(i, 1)
        return
    }
    throw new Error('no layout with maxAspectRatio ' + maxAspectRatio)
}

function removeAllLayouts() { layouts = [] }

function layout() {
    document.body.innerHTML = pickLayout().html
    layoutElement(document.body)
}

// The layoutSelector function adds the string into the selector's element's innerHTML after 
// removing all newlines, then calls layoutElement on the element.
// If the element isn't found then null is returned.
function layoutSelector(s, withString) {
    const e = document.querySelector(s)
    if (e === null) {
        return null
    }
    e.innerHTML = removeNewlines(withString)
    layoutElement(e)
}

function setAllSquareDimensions(modelID, elementsClass) {
    const model = document.querySelector(modelID)
    const w = parseFloat(model.style.width)
    const h = parseFloat(model.style.height)

    const setSelectorAllStyleDims = (selector, styleValue) => {
        for (const e of document.querySelectorAll(selector)) {
            e.style.width = styleValue
            e.style.height = styleValue
        }
    }

    if (w > h) {
        setSelectorAllStyleDims(elementsClass, h + 'px')
        return h
    } else if (h > w) {
        setSelectorAllStyleDims(elementsClass, w + 'px')
        return w
    }
    return h
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

// If your dynamically added inline elements are spilling over to the next line then verify
// that you are removing newlines. You'll see an added space between them otherwise.

// stackoverflow.com/questions/10805125/how-to-remove-all-line-breaks-from-a-string
// stackoverflow.com/questions/12014441/remove-every-white-space-between-tags-using-javascript
function removeNewlines(html) {
    return html.replace(/(<(pre|script|style|textarea)[^]+?<\/\1)|(^|>)\s+|\s+(?=<|$)/g, '$1$3')
}

function interiorDimensions(e) {
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

    const paddingWidth = propValue('padding-left') + propValue('padding-right')
    const paddingHeight = propValue('padding-top') + propValue('padding-bottom')
    
    const borderWidth = propValue('border-left-width') + propValue('border-right-width')
    const borderHeight = propValue('border-top-width') + propValue('border-bottom-width')

    return {
        width: rect.width - (paddingWidth + borderWidth),
        height: rect.height - (paddingHeight + borderHeight)
    }
}

function scaleFont() {
    const pixelCount = window.innerWidth * window.innerHeight
    const scale = (0.0000012 * pixelCount) + 0.3
    for (const s of document.styleSheets) {
        for (const r of s.cssRules) {
            if (r.selectorText !== 'html') {
                continue
            }
            const fontSizeValue = r.style.getPropertyValue('font-size')
            if (fontSizeValue === '') {
                continue
            }
            if (fontSizeValue.includes('pt') === false) {
                throw new Error('CSS html font-size not pt')
            }
            document.querySelector('html').style.fontSize = 
                (parseFloat(fontSizeValue) * scale) + 'pt'
            return
        }
    }
    throw new Error('no CSS html rule that defines font-size')
}

// TODO: add style instead of throwing an error
// TODO: allow setting multiple properties
function addCSSRuleProperty(cssSelector, property, value) {
    for (const s of document.styleSheets) {
        for (const r of s.cssRules) {
            if (r.selectorText !== cssSelector) {
                continue
            }
            r.style.setProperty(property, value)
            return
        }
    }
    throw new Error('no CSS rule for selector ' + cssSelector)
}
