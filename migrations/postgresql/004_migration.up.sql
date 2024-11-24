CREATE TABLE message_votes
(
    id            serial primary key,
    record_id     integer      default 0                     NOT NULL,
    user_id       integer      default 0                     NOT NULL,
    title         varchar(255) default ''::character varying NOT NULL,
    answer_mode   smallint     default 0                     NOT NULL,
    answer_option jsonb                                      NOT NULL,
    answer_num    smallint     default 0                     NOT NULL,
    answered_num  smallint     default 0                     NOT NULL,
    is_anonymous  smallint     default 0                     NOT NULL,
    status        smallint     default 0                     NOT NULL,
    created_at    timestamp                                  NOT NULL,
    updated_at    timestamp                                  NOT NULL
);

CREATE TABLE message_vote_answers
(
    id         serial primary key,
    vote_id    integer default 0          NOT NULL,
    user_id    integer default 0          NOT NULL,
    option     char    default ''::bpchar NOT NULL,
    created_at timestamp                  NOT NULL
);

-- INSERT INTO schema_migrations (version, dirty) VALUES (4, false);
