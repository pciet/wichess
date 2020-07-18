const whitespaceKey = 'whitespace'

// This whitespace var is used in layouts.js to dynamically make the board sublayout.
export let whitespace = true

export function initializeWhitespace() {
    // the default is to not have whitespace, but the default layout has whitespace
    if (window.localStorage.getItem(whitespaceKey) === 'true') {
        return
    }
    swapWhitespace()
}

export function swapWhitespace() {
    if (whitespace === false) {
        whitespace = true
        window.localStorage.setItem(whitespaceKey, true)
    } else {
        whitespace = false
        window.localStorage.setItem(whitespaceKey, false)
    }
}

