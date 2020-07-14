import { collectionKindImage, collectionImage, NotInCollection } from '../collection.js'
import { NoKind } from '../pieceDefs.js'

let FloatingSelection
export let DescribedPiece
export let LeftPicked = NotInCollection
export let RightPicked = NotInCollection
export let RewardPicked = NotInCollection

const LeftSelector = '#left'
const RightSelector = '#right'
const RewardSelector = '#reward'

export function addPieceClicks() {
    document.querySelector(LeftSelector).onclick = rewardClick(LeftSelector, window.LeftPiece)
    document.querySelector(RightSelector).onclick = rewardClick(RightSelector, window.RightPiece)
    document.querySelector(RewardSelector).onclick = 
        rewardClick(RewardSelector, window.RewardPiece)

    for (let i = 0; i < window.Collection.length; i++) {
        const id = '#c'+i
        document.querySelector(id).onclick = collectionClick(id, window.Collection[i], i)
    }
}

function rewardClick(sourceElementID, kind) {
    return () => {
        if (kind === NoKind) {
            return
        }

        if (FloatingSelection !== undefined) {
            document.querySelector(FloatingSelection.selector).classList.remove('selected')
            if (FloatingSelection.selector === sourceElementID) {
                FloatingSelection = undefined
                return
            }
        }

        DescribedPiece = kind

        switch(sourceElementID) {
        case LeftSelector:
            if (LeftPicked !== NotInCollection) {
                deselectReward(LeftSelector, LeftPicked-1)
                LeftPicked = NotInCollection
                return
            }
            break
        case RightSelector:
            if (RightPicked !== NotInCollection) {
                deselectReward(RightSelector, RightPicked-1)
                RightPicked = NotInCollection
                return
            }
            break
        case RewardSelector:
            if (RewardPicked !== NotInCollection) {
                deselectReward(RewardSelector, RewardPicked-1)
                RewardPicked = NotInCollection
                return
            }
            break
        }

        FloatingSelection = {
            kind: kind,
            selector: sourceElementID
        }
        document.querySelector(sourceElementID).classList.add('selected')
    }
}

function deselectReward(sourceSelector, collectionIndex) {
    document.querySelector(sourceSelector).classList.remove('picked')
    const e = document.querySelector('#c'+(collectionIndex))
    e.classList.remove('replaced')
    e.src = collectionImage(collectionIndex)
}

function collectionClick(sourceElementID, kind, collectionIndex) {
    return () => {
        if (FloatingSelection === undefined) {
            DescribedPiece = window.Collection[collectionIndex]
            if (LeftPicked === (collectionIndex+1)) {
                deselectReward(LeftSelector, collectionIndex)
                LeftPicked = NotInCollection
            } else if (RightPicked === (collectionIndex+1)) {
                deselectReward(RightSelector, collectionIndex)
                RightPicked = NotInCollection
            } else if (RewardPicked === (collectionIndex+1)) {
                deselectReward(RewardSelector, collectionIndex)
                RewardPicked = NotInCollection
            }
            return
        }

        if (collectionIndex === (LeftPicked-1)) {
            deselectReward(LeftSelector, collectionIndex)
            LeftPicked = NotInCollection
        } else if (collectionIndex === (RightPicked-1)) {
            deselectReward(RightSelector, collectionIndex)
            RightPicked = NotInCollection
        } else if (collectionIndex === (RewardPicked-1)) {
            deselectReward(RewardSelector, collectionIndex)
            RewardPicked = NotInCollection
        }

        const ce = document.querySelector('#c'+collectionIndex)
        ce.src = collectionKindImage(FloatingSelection.kind)
        ce.classList.add('replaced')

        switch (FloatingSelection.selector) {
        case LeftSelector:
            LeftPicked = collectionIndex+1
            break
        case RightSelector:
            RightPicked = collectionIndex+1
            break
        case RewardSelector:
            RewardPicked = collectionIndex+1
            break
        }

        const e = document.querySelector(FloatingSelection.selector)
        e.classList.remove('selected')
        e.classList.add('picked')
        FloatingSelection = undefined
    }
}
