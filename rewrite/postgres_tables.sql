CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name VARCHAR(24),

    crypt CHAR(64),
    session BYTEA,

    recent_opponents INTEGER[5],
    people_game INTEGER,

    left_kind INTEGER,
    right_kind INTEGER,
    collection BIGINT[21],

    computer_streak INTEGER,
    best_computer_streak INTEGER
);

CREATE TABLE games (
    id SERIAL PRIMARY KEY,

    conceded BOOLEAN,

    white VARCHAR(24),
    white_ack BOOLEAN,
    white_left_kind INTEGER,
    white_right_kind INTEGER,
    white_reward INTEGER,
    white_captures INTEGER[15],

    black VARCHAR(24),
    black_ack BOOLEAN,
    black_left_kind INTEGER,
    black_right_kind INTEGER,
    black_reward INTEGER,
    black_captures INTEGER[15],

    -- 0/false for white,
    -- 1/true for black
    active BOOLEAN,
    previous_active BOOLEAN,

    move_from SMALLINT,
    move_to SMALLINT,

    draw_turns SMALLINT,
    turn SMALLINT,

    board BIGINT[64]
);
