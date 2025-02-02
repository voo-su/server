CREATE TABLE group_chats
(
    id           SERIAL PRIMARY KEY,
    creator_id   INTEGER      DEFAULT 0                     NOT NULL,
    type         INTEGER      DEFAULT 1                     NOT NULL,
    group_name   VARCHAR(30)  DEFAULT ''::CHARACTER VARYING NOT NULL,
    description  VARCHAR(100) DEFAULT ''::CHARACTER VARYING NOT NULL,
    avatar       VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    max_num      INTEGER      DEFAULT 200                   NOT NULL,
    is_overt     INTEGER      DEFAULT 0                     NOT NULL,
    is_mute      INTEGER      DEFAULT 0                     NOT NULL,
    is_dismiss   INTEGER      DEFAULT 0                     NOT NULL,
    created_at   TIMESTAMP                                  NOT NULL,
    updated_at   TIMESTAMP                                  NOT NULL,
    dismissed_at TIMESTAMP
);

CREATE TABLE group_chat_members
(
    id            SERIAL PRIMARY KEY,
    group_id      INTEGER     DEFAULT 0                     NOT NULL,
    user_id       INTEGER     DEFAULT 0                     NOT NULL,
    leader        INTEGER     DEFAULT 0                     NOT NULL,
    user_card     VARCHAR(20) DEFAULT ''::CHARACTER VARYING NOT NULL,
    is_quit       INTEGER     DEFAULT 0                     NOT NULL,
    is_mute       INTEGER     DEFAULT 0                     NOT NULL,
    min_record_id INTEGER     DEFAULT 0                     NOT NULL,
    join_time     TIMESTAMP,
    created_at    TIMESTAMP                                 NOT NULL,
    updated_at    TIMESTAMP                                 NOT NULL
);
