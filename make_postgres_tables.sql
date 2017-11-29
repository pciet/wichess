CREATE TABLE Players (
    name VARCHAR(64) PRIMARY KEY,
    crypt CHAR(64) UNIQUE NOT NULL,
    wins INTEGER,
    losses INTEGER,
    draws INTEGER,
    rating INTEGER,
    c5 INTEGER,
    c15 INTEGER
);

CREATE TABLE Games (
    game_id SERIAL PRIMARY KEY,
    piece INTEGER,
    competitive BOOLEAN,
    recorded BOOLEAN,
    white VARCHAR(64) NOT NULL,
    white_ack BOOLEAN,
    white_latestmove TIMESTAMPTZ,
    white_elapsed BIGINT,
    white_elapsedupdated TIMESTAMPTZ,
    black VARCHAR(64) NOT NULL,
    black_ack BOOLEAN,
    black_latestmove TIMESTAMPTZ,
    black_elapsed BIGINT,
    black_elapsedupdated TIMESTAMPTZ,
    active VARCHAR(64) NOT NULL,
    previous_active VARCHAR(64) NOT NULL,
    move_from SMALLINT,
    move_to SMALLINT,
    draw_turns SMALLINT,
    turn SMALLINT,
    s0 BIGINT, s1 BIGINT, s2 BIGINT, s3 BIGINT,
    s4 BIGINT, s5 BIGINT, s6 BIGINT, s7 BIGINT,
    s8 BIGINT, s9 BIGINT, s10 BIGINT, s11 BIGINT,
    s12 BIGINT, s13 BIGINT, s14 BIGINT, s15 BIGINT,
    s16 BIGINT, s17 BIGINT, s18 BIGINT, s19 BIGINT,
    s20 BIGINT, s21 BIGINT, s22 BIGINT, s23 BIGINT,
    s24 BIGINT, s25 BIGINT, s26 BIGINT, s27 BIGINT,
    s28 BIGINT, s29 BIGINT, s30 BIGINT, s31 BIGINT,
    s32 BIGINT, s33 BIGINT, s34 BIGINT, s35 BIGINT,
    s36 BIGINT, s37 BIGINT, s38 BIGINT, s39 BIGINT,
    s40 BIGINT, s41 BIGINT, s42 BIGINT, s43 BIGINT,
    s44 BIGINT, s45 BIGINT, s46 BIGINT, s47 BIGINT,
    s48 BIGINT, s49 BIGINT, s50 BIGINT, s51 BIGINT,
    s52 BIGINT, s53 BIGINT, s54 BIGINT, s55 BIGINT,
    s56 BIGINT, s57 BIGINT, s58 BIGINT, s59 BIGINT,
    s60 BIGINT, s61 BIGINT, s62 BIGINT, s63 BIGINT
);

CREATE TABLE Pieces (
    piece_id SERIAL PRIMARY KEY,
    kind INTEGER,
    owner VARCHAR(64),
    ingame BOOLEAN
);
