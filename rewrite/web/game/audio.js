import { layoutSelector } from '../layout.js'

const audio = new AudioContext()
let moveAudio
const soundCount = 325
const muteKey = 'mute'
const unmutedChar = '&#x1f50a;'
const mutedChar = '&#x1f507;'

const filter = audio.createBiquadFilter()
filter.type = 'highshelf'
filter.frequency.value = 2500
filter.gain.value = -6

export function muted() {
    if (window.localStorage.getItem(muteKey) === 'true') {
        return true
    }
    return false
}

export function toggleMute() {
    const m = window.localStorage.getItem(muteKey)
    if ((m === undefined) || (m === 'false')) {
        window.localStorage.setItem(muteKey, true)
        setMuteIcon(true)
        return
    }
    window.localStorage.setItem(muteKey, false)
    setMuteIcon(false)
}

export function setMuteIcon(isMuted) {
    if (isMuted === false) {
        layoutSelector('#mute', unmutedChar)
    } else {
        layoutSelector('#mute', mutedChar)
    }
}

export function fetchNextMoveSound() {
    moveAudio = audio.createBufferSource()
    moveAudio.connect(filter)
    filter.connect(audio.destination)
    fetch('/web/sound/click'+Math.floor(Math.random()*soundCount)+'.wav').then(
        r => r.arrayBuffer()).then(r => {
        audio.decodeAudioData(r, b => {
            moveAudio.buffer = b
        }, () => {
            console.log('failed to decode audio clip') 
        })
    })
}

export function moveSound() {
    if (muted()) {
        return
    }
    moveAudio.start(0)
    fetchNextMoveSound()
}