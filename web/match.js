// People play a game by entering each other's username into the input text box. The latest five
// matches are displayed as a list of buttons for fast matching without having to reenter the name.

for (let i = 0; i < window.RecentOpponents.length; i++) {
    const e = document.querySelector('#r'+i)
    const o = window.RecentOpponents[i]
    if (o === '') {
        e.innerHTML = '(No Recent Opponent)'
        continue
    }
    e.innerHTML = o
    e.classList.add('activebutton')
    e.onclick = () => {
        matchOnClick(o)
    }
}

document.querySelector('#comp').onclick = () => {
    fetch('/computer', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: window.localStorage.getItem('army')
    }).then(() => {
        window.location = '/computer'
    })
}

document.querySelector('#oppsubmit').onclick = () => {
    const opponent = document.querySelector('#opp').value
    if (opponent === '') {
        return
    }
    matchOnClick(opponent)
}

let matching = ''
let cancelController

function matchOnClick(opp) {
    matching = opp
    showMatching()
    cancelController = new AbortController()

    fetch('/people?o='+encodeURIComponent(opp), {
        method: 'POST',
        signal: cancelController.signal,
        headers: {
            'Content-Type': 'application/json'
        },
        body: window.localStorage.getItem('army')
    }).then(r => r.json()).then(r => {
        if (r.id === 0) {
            // then there was a timeout
            cancelMatching()
        } else {
            // otherwise a new game is ready to be played
            window.location = '/people/'+r.id
        }
    }).catch(err => {
        if (err.name === 'AbortError') {
            // not an error, player pressed the matching cancel button
            cancelMatching()
            return
        }
        console.error(err)
    })
}

function cancelMatching() {
    matching = ''
    document.querySelector('#matchingname').innerHTML = ''
    document.querySelector('#matchingtext').classList.add('hidden')
    document.querySelector('#cancel').classList.add('hidden')
}

document.querySelector('#cancel').onclick = () => {
    cancelController.abort()
    cancelMatching()
}

function showMatching() {
    document.querySelector('#matchingname').innerHTML = matching
    document.querySelector('#matchingtext').classList.remove('hidden')
    document.querySelector('#cancel').classList.remove('hidden')
}
