import { button } from '../button.js'

export const landscape = `
<div class="inline" id="modepick">
    <div id="logoutdiv">
        <a href="/quit">` + button('', 'logout', 'Quit', false) + `</a>
    </div>
    <div>
        <div></div>
        <div id="title">Wisconsin Chess</div>
        <div id="name">`+Name+`</div>
        <div></div>
    </div>
    <div class="modebuttonmargin">
        <div class="modebutton" id="public">
            <div></div>
            <div>People</div>
            <div></div>
        </div>
    </div>
    <div class="modepickspacer"></div>
    <div class="modebuttonmargin">
        <div class="modebutton" id="computer">
            <div></div>
            <div>Computer</div>
            <div></div>
        </div>
    </div>
</div>
<div class="inline" id="contentdiv">
    <div class="inline" id="contentspacer"></div>
    <div class="inline" id="content"></div>
</div>
`

export const playButton = `
<div id="play">
    <div class="inline"></div>
    <div class="inline" id="playbuttonmargin">
        <div id="playbutton">
            <div class="playspacer"></div>
            <div>Play</div>
            <div class="playspacer"></div>
        </div>
    </div>
    <div class="inline"></div>
</div>
`

export const content = `
<div id="army"></div>
<div id="picks">
    <div class="inline pickcell" id="leftpick"></div>
    <div class="inline" id="pickspacer"></div>
    <div class="inline pickcell" id="rightpick"></div>
</div>
`

export const publicMatching = `
<div class="inline">
    <div class="playerbuttonmargin">
        <div class="playerbutton" id="p0">
            <div></div>
            <div>Recent Opponent 1</div>
            <div></div>
        </div>
    </div>
    <div class="playerbuttonmargin">
        <div class="playerbutton" id="p1">
            <div></div>
            <div>Recent Opponent 2</div>
            <div></div>
        </div>
    </div>
    <div class="playerbuttonmargin">
        <div class="playerbutton" id="p2">
            <div></div>
            <div>Recent Opponent 3</div>
            <div></div>
        </div>
    </div>
    <div class="playerbuttonmargin">
        <div class="playerbutton" id="p3">
            <div></div>
            <div>Recent Opponent 4</div>
            <div></div>
        </div>
    </div>
    <div class="playerbuttonmargin">
        <div class="playerbutton" id="p4">
            <div></div>
            <div>Recent Opponent 5</div>
            <div></div>
        </div>
    </div>
</div>
<div class="inline" id="matchingspacer"></div>
<div class="inline">
    <div></div>
    <div id="opponentinputmargin">
        <input type="text" id="opponent">
    </div>
    <span id="opponentlabel">Type opponent's username.</span>
    <div id="matchbuttonmargin">
        <div id="match">
            <div></div>
            <div>New Match</div>
            <div></div>
        </div>
    </div>
</div>
`

export const matchPending = `
    <div class="matchcancelspacer"></div>
    <div>
        <div class="inline"></div>
        <div class="inline">
            <span class="matchingtext">Matching</span>
            <div class="matchingplayer" id="matchingopponent"></div>
            <span class="matchingtext">against</span>
            <div class="matchingplayer" id="matchingplayer"></div>
            <div>
                <div class="inline cancelbuttonspacer"></div>
                <div class="inline" id="cancelbuttonmargin">
                    <div id="cancelbutton">
                        <div></div>
                        <div>Cancel</div>
                        <div></div>
                    </div>
                </div>
                <div class="inline cancelbuttonspacer"></div>
            </div>
        </div>
        <div class="inline"></div>
    </div>
    <div class="matchcancelspacer"></div>
`
