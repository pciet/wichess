// A game is in one of these six states, only continuing from the abnormal states of 
// promotion, reverse promotion, and check.
export const State = {
    NORMAL: 0,
    PROMOTION: 1,
    CHECK: 2,
    CHECKMATE: 3,
    DRAW: 4,
    CONCEDED: 5,
    TIME_OVER: 6,
    REVERSE_PROMOTION: 7
}

function StateDef(name) {
    this.name = name
}

export const States = [
    new StateDef('Normal'),
    new StateDef('Promotion'),
    new StateDef('Check'),
    new StateDef('Checkmate'),
    new StateDef('Draw'),
    new StateDef('Conceded'),
    new StateDef('Time Over'),
    new StateDef('Reverse Promotion')
]
