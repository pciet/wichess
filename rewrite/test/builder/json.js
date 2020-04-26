export function CaseJSON(name, orientation, 
    previousFrom, previousTo, state, position, moves) {
    return JSON.stringify({
        n: name,
        o: orientation,
        m: {
            f: previousFrom,
            t: previousTo,
        },
        s: state,
        p: position,
        mo: moves
    })
}
