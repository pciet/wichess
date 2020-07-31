import { addLayout, layout, scaleFont, setAllSquareDimensions } from './layout.js'
import { Pieces } from './pieceDefs.js'
import { NotInCollection } from './collection.js'

import { landscape } from './reward/layouts.js'
import { addPieceClicks, DescribedPiece, 
    LeftPicked, RightPicked, RewardPicked } from './reward/click.js'

addLayout(100, landscape)

export function layoutPage() {
    layout()
    scaleFont()

    setAllSquareDimensions('#left', '.rewardcell')
    setAllSquareDimensions('#cc0', '.collectioncell')

    document.querySelector('#ack').onclick = () => {
        const ack = () => {
            fetch('/acknowledge/' + window.GameID).then(() => { window.location = '/' })
        }
        if ((LeftPicked === NotInCollection) && 
            (RightPicked === NotInCollection) && (RewardPicked === NotInCollection)) {
            ack()
            return
        }
        fetch('/reward/' + window.GameID, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                l: LeftPicked,
                r: RightPicked,
                re: RewardPicked
            })
        }).then(ack)
    }

    document.querySelector('#details').onclick = () => {
        if (DescribedPiece === undefined) {
            return
        }
        window.open('/details?p=' + Pieces[DescribedPiece])
    }

    addPieceClicks()

    document.body.classList.add('visible')
}

let resizeWait

window.onresize = () => {
    clearTimeout(resizeWait)
    resizeWait = setTimeout(layoutPage, 150)
}

window.onload = layoutPage
