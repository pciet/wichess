// randomInt returns a random integer in the 
// inclusive range [0-maxPossible].
export function randomInt(maxPossible) {
    return Math.floor(Math.random() * (maxPossible+1))
}
