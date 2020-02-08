// layout.js aids making a website that doesn't scroll and adapts to the window size, useful for dashboards or games.
// The meaning of the HTML language is changed in some ways, with emphasis on an interface of buttons instead of text presentation.
//
// Why use HTML and web browsers?
// A web browser lets you make a network computing application with the convenience for people of not having to install another approved app.
//
// The browser window is divided by regular HTML blocks marked by the <div> tag in HTML.
// You set each block's CSS 'display' tag depending on how you want it to look, 'block' for a column or 'inline-block' for a row.
// 'block' and 'inline-block' cannot be mixed for sibling blocks.
// layout.css must be part of your page. It includes the 'inline' class to apply 'inline-block' for you, and it disables the text presentation box model.
// You set each block's CSS 'width' and 'height' properties to percentages to define the size of each block.
// If no dimensions are set then layout.js evenly sizes each block, or if only some are set then the rest evenly take the remaining space.
// The primary use of layout.js is to do this layout when the page is loaded and whenever the browser window is resized.
//
// layout.js lets your website look different depending on the the screen dimensions.
// Multiple layouts are set for the page that depend on the aspect ratio, width divided by height, which means portrait < 1 < landscape.
// In your page's JavaScript you call 'addLayout' first to define the different page looks.
// If part of that HTML is best made programatically then do that in the 'generateFunction' argument.
// 'generateFunction' is called after the HTML is initially rendered but before the layout process happens.
// 'maxRatio' is the upper value of the aspect ratio range the corresponding layout is used for.
// The layout with the smallest 'maxRatio' that is larger than the actual aspect ratio is choosen.

export { addLayout, layout }

class Layout {
    constructor(maxRatio, html, generateFunction) {
        this.maxRatio = maxRatio
        // TODO: should the HTML have spaces removed?
        this.html = html
        this.gen = generateFunction
    }
}

let layouts = []

function addLayout(maxAspectRatio, html, generateFunc) {
    // https://stackoverflow.com/questions/10805125/how-to-remove-all-line-breaks-from-a-string
    // https://stackoverflow.com/questions/12014441/remove-every-white-space-between-tags-using-javascript
    layouts.push(new Layout(maxAspectRatio, html.replace(/(<(pre|script|style|textarea)[^]+?<\/\2)|(^|>)\s+|\s+(?=<|$)/g, '$1$3'), generateFunc))
}

function layout() {
    const w = window.innerWidth
    const h = window.innerHeight

    document.body.style.width = w + 'px'
    document.body.style.height = h + 'px'

    const l = pick(w, h)

    document.body.innerHTML = l.html

    if (l.gen !== undefined) {
        l.gen()
    }

    layoutElement(document.body)
}

function pick(width, height) {
    if (layouts.length === 0) {
        throw new Error('no layouts set')
    }

    const ratio = width / height

    let use = undefined
    for (const l of layouts) {
        if (use === undefined) {
            use = l
            continue
        }
        if (use.maxRatio === l.maxRatio) {
            throw new Error('duplicate layout maxAspectRatio')
        }
        if (l.maxRatio < ratio) {
            break
        }
        use = l
    }

    if (use === undefined) {
        throw new Error('no layout')
    }

    return use
}

// Recursively travels the tree of document elements and applies the layout.js rules at each level.
function layoutElement(e) {
    // first determine if the children of e already have dimensions set, looking in order of ID, class, then tag
    let widths = []
    let heights = []

    const bs = window.getComputedStyle(e).getPropertyValue('box-sizing')
    if (bs !== 'border-box') {
        throw new Error(e.tagName + ' box-sizing property is "' + bs + '", should be set as "border-box"')
    }

    // selector can be an ID (#id), class (.class), or tag
    // readCSSDimensions is not the inline or computed style, it's what's actually typed in the CSS text
    // TODO: both width and height must be defined in one CSS property block, is that ok or is there value to defining width and height in separate places?
    const readCSSDimensions = (selector, index) => {
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
            return false
        }

        const w = rule.style.getPropertyValue('width')
        if (w !== '') {
            if (w.includes('%') === false) {
                throw new Error(selector + ' CSS width property not a percentage')
            }
        }

        const h = rule.style.getPropertyValue('height')
        if (h !== '') {
            if (w.includes('%') === false) {
                throw new Error(selector + ' CSS height property not a percentage')
            }
        }

        if ((w === '') && (h === '')) {
            return false
        }

        if (w === '') {
            widths[index] = undefined
        } else {
            widths[index] = parseFloat(w)
        }

        if (h === '') {
            heights[index] = undefined
        } else {
            heights[index] = parseFloat(h)
        }

        return true
    }

    let elements = []
    let inline = undefined
    for (let i = 0; i < e.children.length; i++) {
        const c = e.children[i]

        elements[i] = c

        let ei
        const inlineValue = window.getComputedStyle(c).getPropertyValue('display')
        if (inlineValue === 'inline-block') {
            ei = true
        } else {
            ei = false
        }

        if (inline === undefined) {
            inline = ei
        } else {
            if (((inline === false) && ei) || (inline && (ei === false))) {
                throw new Error('mix of block and inline-block sibling elements')
            }
        }

        if (c.id !== '') {
            if (readCSSDimensions(c.id, i)) {
                continue
            }
        }

        let t
        for (const cl of c.classList) {
            t = readCSSDimensions('.'+cl, i)
            if (t) {
                break
            }
        }
        if (t) {
            continue
        }

        if (readCSSDimensions(c.tagName.toLowerCase(), i)) {
            continue
        }

        widths[i] = undefined
        heights[i] = undefined
    }

    if (elements.length === 0) {
        return
    }

    // block elements are constrained together for height but can be any width
    if (inline === false) {
        let totalHeight = 0
        let undefinedHeights = 0
        for (let i = 0; i < elements.length; i++) {
            if (heights[i] !== undefined) {
                totalHeight += heights[i]
            } else {
                undefinedHeights++
            }
        }

        if (totalHeight > 100) {
            throw new Error('sibling block height greater than 100%')
        }

        if (undefinedHeights > 0) {
            const undefinedHeight = (100 - totalHeight) / undefinedHeights
            for (let i = 0; i < elements.length; i++) {
                if (heights[i] === undefined) {
                    heights[i] = undefinedHeight
                    widths[i] = 100
                }
            }
        }
    } else {
        let totalWidth = 0
        let undefinedWidths = 0
        for (let i = 0; i < elements.length; i++) {
            // inline-block elements require vertical-align: top to be set for correct spacing
            elements[i].style.verticalAlign = 'top'

            if (widths[i] !== undefined) {
                totalWidth += widths[i]
            } else {
                undefinedWidths++
            }
        }

        if (totalWidth > 100) {
            throw new Error('sibling inline-block width greater than 100%')
        }

        if (undefinedWidths > 0) {
            const undefinedWidth = (100 - totalWidth) / undefinedWidths
            for (let i = 0; i < elements.length; i++) {
                if (widths[i] === undefined) {
                    widths[i] = undefinedWidth
                    heights[i] = 100
                }
            }
        }
    }

    for (let i = 0; i < elements.length; i++) {
        elements[i].style.width = ((widths[i] / 100) * elements[i].parentElement.clientWidth) + 'px'
        elements[i].style.height = ((heights[i] / 100) * elements[i].parentElement.clientHeight) + 'px'
        layoutElement(elements[i])
    }
}
