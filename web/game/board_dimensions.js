import { addCSSRuleProperty } from '../layout.js'

export function writeBoardDimension(lowerSquareRatio, upperSquareRatio) {
    const r = window.innerWidth / window.innerHeight
    if ((r <= upperSquareRatio) && (r > lowerSquareRatio)) {
        // depends on height of .squareboardbar, -10%*2
        const h = window.innerHeight * 0.8
        // width is in the board column which is window width - #squarebar width of 25%
        const w = window.innerWidth * 0.75

        if (h > w) {
            addCSSRuleProperty('#board', 'width', '100%')
            addCSSRuleProperty('#board', 'height', ((w / window.innerHeight) * 100) + '%')
        } else {
            addCSSRuleProperty('#board', 'height', '80%')
            addCSSRuleProperty('#board', 'width', ((h / w) * 100) + '%')
        }

        return
    }

    // if not the square layout then this:

    let d
    if (window.innerWidth < window.innerHeight) {
        d = window.innerWidth
    } else {
        d = window.innerHeight
    }

    addCSSRuleProperty('#board', 'width', ((d / window.innerWidth) * 100) + '%')
    addCSSRuleProperty('#board', 'height', ((d / window.innerHeight) * 100) + '%')
}
