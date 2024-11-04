create table message_votes
(
    id            serial primary key,
    record_id     integer      default 0                     not null,
    user_id       integer      default 0                     not null,
    title         varchar(255) default ''::character varying not null,
    answer_mode   smallint     default 0                     not null,
    answer_option jsonb                                      not null,
    answer_num    smallint     default 0                     not null,
    answered_num  smallint     default 0                     not null,
    is_anonymous  smallint     default 0                     not null,
    status        smallint     default 0                     not null,
    created_at    timestamp                                  not null,
    updated_at    timestamp                                  not null
);

create table message_vote_answers
(
    id         serial primary key,
    vote_id    integer default 0          not null,
    user_id    integer default 0          not null,
    option     char    default ''::bpchar not null,
    created_at timestamp                  not null
);

INSERT INTO schema_migrations (version, dirty) VALUES (4, false);
