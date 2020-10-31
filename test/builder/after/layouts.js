export const newOrLoadCase = `
<div class="optionrow">
    <label for="cases">Case:</label>
    <select class="inline" id="cases"></select>
    <button class="inline" id="loadcase" type="button">Load</button>
    <label for="newcase">&emsp;&emsp;&emsp;&emsp;Title:</label>
    <input class="inline" type="text" id="newcasetext">
    <button class="inline" id="newcase" type="button">New</button>
</div>
`

export const caseMaker = `
<div>
    <div class="inline">
        <div id="board"></div>
    </div>
    <div class="inline">
        <h2>Position</h2>
        <div class="optionrow">
            <select class="inline" id="pieces"></select>
            <select class="inline" id="pieceorientation">
                <option value="0">White</option>
                <option value="1">Black</option>
            </select>
        </div>
        <div class="optionrow">
            <label for="moved">Moved:</label>
            <select class="inline" id="moved">
                <option value="0">False</option>
                <option value="1">True</option>
            </select>
        </div>
        <div class="optionrow">
            <button class="inline" id="pickstart" type="button">Select Start</button>
        </div>
        <h2>Move</h2>
        <div class="optionrow">
            <span class="inline">From:</span>
            <span class="inline" id="from">-</span>
            <span class="inline">To:</span>
            <span class="inline" id="to">-</span>
        </div>
        <div class="optionrow">
            <button id="move" type="button">Select Move</button>
        </div>
        <h2>Previous Move</h2>
        <div class="optionrow">
            <span class="inline">From:</span>
            <span class="inline" id="prevfrom">-</span>
            <span class="inline">To:</span>
            <span class="inline" id="prevto">-</span>
        </div>
        <div class="optionrow">
            <button id="prevmove" type="button">Select Previous Move</button>
        </div>
    </div>
</div>

<br>

<div>
    <div id="changesboard" class="inline"></div>
    <div class="inline">
        <h2>Changes</h2>
        <div class="optionrow">
            <select class="inline" id="changepieces"></select>
            <select class="inline" id="changepieceorientation">
                <option value="0">White</option>
                <option value="1">Black</option>
            </select>
        </div>
        <div class="optionrow">
            <button id="deletechange" type="button">Delete Square</button>
        </div>
    </div>
</div>

<br>

<button id="savecase" type="button">Save</button>
`
