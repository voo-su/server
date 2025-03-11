CREATE TABLE message_votes
(
    id            SERIAL PRIMARY KEY,
    message_id    INTEGER      DEFAULT 0                     NOT NULL,
    user_id       INTEGER      DEFAULT 0                     NOT NULL,
    title         VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    answer_mode   INTEGER      DEFAULT 0                     NOT NULL,
    answer_option JSONB                                      NOT NULL,
    answer_num    INTEGER      DEFAULT 0                     NOT NULL,
    answered_num  INTEGER      DEFAULT 0                     NOT NULL,
    is_anonymous  INTEGER      DEFAULT 0                     NOT NULL,
    status        INTEGER      DEFAULT 0                     NOT NULL,
    created_at    TIMESTAMP                                  NOT NULL,
    updated_at    TIMESTAMP                                  NOT NULL
);

CREATE TABLE message_vote_answers
(
    id         SERIAL PRIMARY KEY,
    vote_id    INTEGER DEFAULT 0          NOT NULL,
    user_id    INTEGER DEFAULT 0          NOT NULL,
    option     CHAR    DEFAULT ''::bpchar NOT NULL,
    created_at TIMESTAMP                  NOT NULL
);
