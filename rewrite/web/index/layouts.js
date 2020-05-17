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
            <div>Public</div>
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
