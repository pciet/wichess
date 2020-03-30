import { addCSSRuleProperty } from '../layout.js'

export function writeBoardDimension(lowerSquareRatio, upperSquareRatio) {
    let d
    if (window.innerWidth < window.innerHeight) {
        d = window.innerWidth
    } else {
        d = window.innerHeight
    }

    const r = window.innerWidth / window.innerHeight
    if ((r <= upperSquareRatio) && (r > lowerSquareRatio)) {
        d = d * 0.75
        addCSSRuleProperty('#boardrow', 'height', ((d / window.innerHeight) * 100) + '%')
        addCSSRuleProperty('#board', 'height', '100%')
        addCSSRuleProperty('#board', 'width', ((d / window.innerWidth) * 100) + '%')
        return
    }

    // if not the square layout then this:
    addCSSRuleProperty('#board', 'width', ((d / window.innerWidth) * 100) + '%')
    addCSSRuleProperty('#board', 'height', ((d / window.innerHeight) * 100) + '%')
}
