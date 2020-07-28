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
<div class="inline">
    <div id="board"></div>
</div>
<div class="inline">
    <div class="optionrow">
        <label for="orientation">Active Player:</label>
        <select class="inline" id="orientation">
            <option value="0">White</option>
            <option value="1">Black</option>
        </select>
    </div>
    <div class="optionrow">
        <label for="state">State:</label>
        <select class="inline" id="state"></select>
    </div>
    <div class="optionrow">
        <span class="inline">From:</span>
        <span class="inline" id="from">-</span>
        <span class="inline">To:</span>
        <span class="inline" id="to">-</span>
    </div>
    <div class="optionrow">
        <button id="previous" type="button">Select Previous</button>
    </div>
    <p>-----</p>
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
    <p>-----</p>
    <div class="optionrow">
        <div class="inline" id="moves"></div>
        <div class="inline" id="next"></div>
    </div>
    <p>-----</p>
    <button id="savecase" type="button">Save</button>
</div>
`
