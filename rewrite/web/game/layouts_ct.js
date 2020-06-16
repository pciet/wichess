// ct stands for "centered text", a div structure compatible with layout.js that can be used
// to vertically and horizontally center text in a styled box. Use the padding of #[id]margin
// for margin, text is in #[id]text, and any classes in the argument are applied to #[id].
// The click handler should be applied to #[id] to create a button.
export function ct(id, classes = '', inline = false, noselect = true, text = '') {
    let t = '<div '
    if (inline === true) {
        t += 'class="inline" '
    }
    t += `id="`+id+`margin">
    <div `
    if (classes !== '') {
    t += 'class="'+classes+'" '
    }
    t += `id="`+id+`">
        <div></div>
    <div `
    if (noselect === true) {
        t += 'class="noselect" '
    }
    t += `id="`+id+`text">`
    if (text !== '') {
        t += text
    }
    return t + `</div>
        <div></div>
    </div>
</div>
`
}

