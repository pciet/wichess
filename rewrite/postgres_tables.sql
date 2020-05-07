CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    name VARCHAR(24),

    crypt CHAR(64),
    session BYTEA,

    left_kind INTEGER,
    right_kind INTEGER,
    collection BIGINT[24],

    computer_streak INTEGER,
    best_computer_streak INTEGER
);

CREATE TABLE games (
    id SERIAL PRIMARY KEY,

    conceded BOOLEAN,

    white VARCHAR(24),
    white_ack BOOLEAN,

    black VARCHAR(24),
    black_ack BOOLEAN,

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
