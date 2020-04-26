import { armySelectionJSON } from './army.js'

export function addComputerPlayClick() {
    document.querySelector('#playbutton').onclick = () => {
        fetch('/computer', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: armySelectionJSON()
        }).then(() => {
            window.location = '/computer'
        })
    }
}
