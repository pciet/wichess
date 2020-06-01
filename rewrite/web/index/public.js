import { publicMatching, matchPending } from './layouts.js'
import { armySelectionJSON } from './army.js'

import { layoutSelector } from '../layout.js'
import { layoutPage } from '../index.js'

// The public mode (renamed to "People") connects players by each entering the other's username
// into an input box. If both players enter each other then the game starts. The latest five
// players are displayed in a list of buttons for fast matching without having to enter a name.

let matching = ''
let cancelController

export function addPublicMatches() {
    if (matching !== '') {
        addMatching()
        return
    }
    layoutSelector('#publics', publicMatching)
    for (let i = 0; i < 5; i++) {
        if (RecentOpponents[i] === '') {
            continue
        }
        const pb = document.querySelector('#p'+i)
        pb.innerHTML = RecentOpponents[i]
        pb.classList.add('playernamed')
    }

    document.querySelector('#match').onclick = () => {
        const opponent = document.querySelector('#opponent').value
        if (opponent === '') {
            return
        }

        matching = opponent
        addMatching()
        cancelController = new AbortController()

        fetch('/people?o='+encodeURIComponent(opponent), {
            method: 'POST',
            signal: cancelController.signal,
            headers: {
                'Content-Type': 'application/json'
            },
            body: armySelectionJSON()
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
}

function cancelMatching() {
    matching = ''
    layoutPage()
    addPublicMatches()
}

function addMatching() {
    layoutSelector('body', matchPending)
    document.querySelector('#matchingopponent').innerHTML = matching
    document.querySelector('#matchingplayer').innerHTML = Name
    document.querySelector('#cancelbutton').onclick = () => {
        cancelController.abort()
    }
}
