import { publicMatching } from './layouts.js'
import { armySelectionJSON } from './army.js'

import { layoutSelector } from '../layout.js'

// The public mode (renamed to "People") connects players by each entering the other's username
// into an input box. If both players enter each other then the game starts. The latest five
// players are displayed in a list of buttons for fast matching without having to enter a name.

export function addPublicMatches() {
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
        fetch('/people?o='+encodeURIComponent(opponent), {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: armySelectionJSON()
        }).then(r => r.json()).then(id => {
            window.location = '/people/'+id.id
        })
    }
}
