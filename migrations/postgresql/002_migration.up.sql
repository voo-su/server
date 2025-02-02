CREATE TABLE chats
(
    id          SERIAL PRIMARY KEY,
    chat_type   INTEGER DEFAULT 1 NOT NULL,
    user_id     INTEGER DEFAULT 0 NOT NULL,
    receiver_id INTEGER DEFAULT 0 NOT NULL,
    is_top      INTEGER DEFAULT 0 NOT NULL,
    is_disturb  INTEGER DEFAULT 0 NOT NULL,
    is_delete   INTEGER DEFAULT 0 NOT NULL,
    is_bot      INTEGER DEFAULT 0 NOT NULL,
    created_at  TIMESTAMP         NOT NULL,
    updated_at  TIMESTAMP         NOT NULL
);

CREATE INDEX chats_chat_type_user_id_receiver_id_idx ON chats (chat_type, user_id, receiver_id);

CREATE TABLE messages
(
    id          bigSERIAL PRIMARY KEY,
    msg_id      VARCHAR(50) DEFAULT ''::CHARACTER VARYING NOT NULL,
    sequence    INTEGER     DEFAULT 0                     NOT NULL,
    chat_type   INTEGER     DEFAULT 1                     NOT NULL,
    msg_type    INTEGER     DEFAULT 1                     NOT NULL,
    user_id     INTEGER     DEFAULT 0                     NOT NULL,
    receiver_id INTEGER     DEFAULT 0                     NOT NULL,
    is_revoke   INTEGER     DEFAULT 0                     NOT NULL,
    is_mark     INTEGER     DEFAULT 0                     NOT NULL,
    is_read     INTEGER     DEFAULT 0                     NOT NULL,
    quote_id    VARCHAR(50)                               NOT NULL,
    content     TEXT,
    extra       JSONB                                     NOT NULL
        CONSTRAINT message_records_extra_check CHECK (extra IS JSON
) ,
    created_at  TIMESTAMP                                 NOT NULL,
    updated_at  TIMESTAMP                                 NOT NULL
);
