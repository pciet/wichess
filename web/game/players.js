import { ActiveOrientation } from '../game.js'
import { Orientation } from '../piece.js'

export function writePlayersIndicator() {
    document.querySelector('#blacknametext').innerHTML = GameInformation.Black.Name
    document.querySelector('#whitenametext').innerHTML = GameInformation.White.Name
    if (hasComputerPlayer() === true) {
        return
    }
    writeActivePlayerIndicator()
}

export function writeActivePlayerIndicator() {
    if (hasComputerPlayer() === true) {
        return
    }
    if (ActiveOrientation === Orientation.WHITE) {
        document.querySelector('#blacknameactive').hidden = true
        document.querySelector('#whitenameactive').hidden = false
    } else {
        document.querySelector('#blacknameactive').hidden = false
        document.querySelector('#whitenameactive').hidden = true
    }
}

export const ComputerPlayerName = 'Computer Player'

export function hasComputerPlayer() {
    if ((GameInformation.Black.Name === ComputerPlayerName) ||
        (GameInformation.White.Name === ComputerPlayerName)) {
        return true
    }
    return false
}
