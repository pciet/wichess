export function button(buttonclass, id, text, inline) {
    let b = '<div class="buttonmargin'
    if (inline === true) {
        b += ' inline'
    }
    b += `">
            <div class="buttonsurface ` + buttonclass + `" id="` + id + `">
                <div class="buttonverticalcentering"></div>
                <div class="buttontext">` + text + `</div>
                <div class="buttonverticalcentering"></div>
            </div>
        </div>
    `
    return b
}

export function selectButton(id) {
    const s = document.querySelector('.buttonselected')
    if (s !== null) {
        s.classList.remove('buttonselected')
        s.querySelector('.buttontext').classList.remove('buttontextselected')
    }
    const e = document.querySelector(id)
    if (e === null) {
        throw new Error('no button element ' + id)
    }
    e.classList.add('buttonselected')
    e.querySelector('.buttontext').classList.add('buttontextselected')
}

