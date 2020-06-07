export function writePlayersIndicator() {
    document.querySelector('#blacknametext').innerHTML = GameInformation.Black.Name
    document.querySelector('#whitenametext').innerHTML = GameInformation.White.Name
}

export const ComputerPlayerName = 'Computer Player'

export function hasComputerPlayer() {
    if ((GameInformation.Black.Name === ComputerPlayerName) ||
        (GameInformation.White.Name === ComputerPlayerName)) {
        return true
    }
    return false
}
