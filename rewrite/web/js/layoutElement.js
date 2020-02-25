import { elementInteriorDimensions } from './layout.js'

// Recursively travels the document rooted at element e to calculate and write dimensions.
export function layoutElement(e) {
    // an element with no children has already had its dimensions written
    if (e.children.length === 0) {
        return
    }

    // border-box means padding and border don't need to be in these calculations
    if (window.getComputedStyle(e).getPropertyValue('box-sizing') !== 'border-box') {
        throw new Error(e.tagName + ' box-sizing should be border-box')
    }

    // read dimensions defined as CSS percentage strings
    const dims = childrenDefinedDimensions(e)

    // calculates and defines percentages that were not in CSS
    const allDims = allChildrenDimensions(e, dims)

    // 'inline' CSS means these dimension numbers are written into the HTML tags
    writeInlineCSSDimensions(e, allDims)

    for (const ei of e.children) {
        layoutElement(ei)
    }
}

// Checks in order of ID (#), class (.), then HTML tag for any CSS defining dimensions.
// The condition of no mixing inline-block and block is verified.
function childrenDefinedDimensions(parent) {
    const dims = {
        widths: [],
        heights: [],
        displayInline: undefined
    }
    
    for (let i = 0; i < parent.children.length; i++) {
        const e = parent.children[i]

        let childInline
        const display = window.getComputedStyle(e).getPropertyValue('display')
        if (display === 'inline-block') {
            childInline = true
        } else if (display === 'inline') { // TODO: adapt to updated display specification
            childInline = true
        } else if (display === 'block') {
            childInline = false
        } else {
            throw new Error('display "' + display + '" not block or inline-block for ' + e.tagName)
        }

        if (dims.displayInline === undefined) {
            dims.displayInline = childInline
        } else if (dims.displayInline != childInline) {
            console.log(parent.children)
            throw new Error('mix of block and inline-block sibling elements')
        }

        let dim = undefined

        if (e.id !== '') {
            dim = cssDimensions('#'+e.id)
            if (dim !== undefined) {
                dims.widths[i] = dim.width
                dims.heights[i] = dim.height
                continue
            }
        }

        for (const cl of e.classList) {
            dim = cssDimensions('.'+cl)
            if (dim !== undefined) {
                dims.widths[i] = dim.width
                dims.heights[i] = dim.height
                break
            }
        }
        if (dim !== undefined) {
            continue
        }

        dim = cssDimensions(e.tagName.toLowerCase())
        if (dim !== undefined) {
            dims.widths[i] = dim.width
            dims.heights[i] = dim.height
            continue
        }

        dims.widths[i] = undefined
        dims.heights[i] = undefined
    }

    return dims
}

// Returns undefined if no dimensions are defined for this selector, or returns {width, height} percentages with maybe one undefined.
function cssDimensions(selector) {
    let rule = undefined
    for (const s of document.styleSheets) {
        for (const r of s.cssRules) {
            if (r.selectorText !== selector) {
                continue
            }
            rule = r
            break
        }
        if (rule !== undefined) {
            break
        }
    }
    if (rule === undefined) {
        return undefined
    }

    const w = rule.style.getPropertyValue('width')
    if (w !== '') {
        if (w.includes('%') === false) {
            throw new Error(selector + ' CSS width "' + w + '" not percent')
        }
    }

    const h = rule.style.getPropertyValue('height')
    if (h !== '') {
        if (h.includes('%') === false) {
            throw new Error(selector + ' CSS height "' + h + '" not percent')
        }
    }

    if ((w === '') && (h === '')) {
        return undefined
    }

    const dims = {
        width: undefined,
        height: undefined
    }

    if (w !== '') {
        dims.width = parseFloat(w)
    }

    if (h !== '') {
        dims.height = parseFloat(h)
    }

    return dims
}

function allChildrenDimensions(parent, definedDimensions) {
    let sameDim = undefined
    let varyingDim = undefined

    if (definedDimensions.displayInline === false) {
        // block elements are all the same height but can be any width
        sameDim = definedDimensions.heights
        varyingDim = definedDimensions.widths
    } else if (definedDimensions.displayInline === true) {
        // inline elements are all the samd width but can be any height
        sameDim = definedDimensions.widths
        varyingDim = definedDimensions.heights
        
        // inline elements must have 'vertical-align: top' for correct spacing
        for (const e of parent.children) {
            e.style.verticalAlign = 'top'
        }
    } else {
        throw new Error('undefined display inline')
    }

    let totalDim = 0
    let undefinedDims = 0
    for (let i = 0; i < sameDim.length; i++) {
        if (sameDim[i] !== undefined) {
            totalDim += sameDim[i]
        } else {
            undefinedDims++
        }

        if (varyingDim[i] === undefined) {
            varyingDim[i] = 100
        }
    }

    if (totalDim > 100) {
        console.log('parent:')
        console.log(parent)
        throw new Error('sibling total dimension greater than 100%')
    }

    if (undefinedDims > 0) {
        const undefinedDim = (100 - totalDim) / undefinedDims
        for (let i = 0; i < sameDim.length; i++) {
            if (sameDim[i] === undefined) {
                sameDim[i] = undefinedDim
            }
        }
    }

    if (definedDimensions.displayInline === false) {
        return {
            widths: varyingDim,
            heights: sameDim
        }
    } else {
        return {
            widths: sameDim,
            heights: varyingDim
        }
    }
}

function writeInlineCSSDimensions(parent, dimensions) {
    const pd = elementInteriorDimensions(parent)
    for (let i = 0; i < parent.children.length; i++) {
        parent.children[i].style.width = ((dimensions.widths[i] / 100) * pd.width) + 'px'
        parent.children[i].style.height = ((dimensions.heights[i] / 100) * pd.height) + 'px'
    }
}
