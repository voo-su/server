CREATE TABLE message_delete
(
    id         SERIAL PRIMARY KEY,
    record_id  INTEGER DEFAULT 0 NOT NULL,
    user_id    INTEGER DEFAULT 0 NOT NULL,
    created_at TIMESTAMP         NOT NULL
);

CREATE TABLE message_read
(
    id          SERIAL PRIMARY KEY,
    msg_id      VARCHAR(64) DEFAULT '':: CHARACTER VARYING NOT NULL,
    user_id     INTEGER     DEFAULT 0                      NOT NULL,
    receiver_id INTEGER     DEFAULT 0                      NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE                   NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE                   NOT NULL,
    CONSTRAINT unique_user_receiver_msg
        UNIQUE (user_id, receiver_id, msg_id)
);
