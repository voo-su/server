CREATE TABLE file_splits
(
    id            SERIAL PRIMARY KEY,
    type          INTEGER      DEFAULT 1                     NOT NULL,
    drive         INTEGER      DEFAULT 1                     NOT NULL,
    upload_id     VARCHAR(100) DEFAULT ''::CHARACTER VARYING NOT NULL,
    user_id       INTEGER      DEFAULT 0                     NOT NULL,
    original_name VARCHAR(100) DEFAULT ''::CHARACTER VARYING NOT NULL,
    split_index   INTEGER      DEFAULT 0                     NOT NULL,
    split_num     INTEGER      DEFAULT 0                     NOT NULL,
    path          VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    file_ext      VARCHAR(10)  DEFAULT ''::CHARACTER VARYING NOT NULL,
    file_size     INTEGER                                    NOT NULL,
    is_delete     INTEGER      DEFAULT 0                     NOT NULL,
    attr          JSONB                                      NOT NULL,
    created_at    TIMESTAMP                                  NOT NULL,
    updated_at    TIMESTAMP                                  NOT NULL
);

CREATE TABLE bots
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER      DEFAULT 0                     NOT NULL,
    bot_type    INTEGER      DEFAULT 0,
    name        VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    description VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    avatar      VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    created_at  TIMESTAMP                                  NOT NULL
);
