import { Pieces } from '../pieceDefs.js'

export function collectionImage(index) { 
    return '/web/img/pick_'+Pieces[window.Collection[index]]+'.png'
}

export function pickImage(kind) { return '/web/img/pick_'+Pieces[kind]+'.png' }

export const LeftPick = -1
export const RightPick = -2
export const NotInCollection = 0

export function collectionSelector(slot) {
    switch (slot) {
    case NotInCollection:
        return ''
    case LeftPick:
        return '#leftpick'
    case RightPick:
        return '#rightpick'
    }
    return '#c' + (slot-1)
}
