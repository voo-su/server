CREATE TABLE message_delete
(
    id         serial primary key,
    record_id  integer default 0 NOT NULL,
    user_id    integer default 0 NOT NULL,
    created_at timestamp         NOT NULL
);

CREATE TABLE message_read
(
    id          serial primary key,
    msg_id      varchar(64) default '':: character varying NOT NULL,
    user_id     integer     default 0                      NOT NULL,
    receiver_id integer     default 0                      NOT NULL,
    created_at  timestamp with time zone                   NOT NULL,
    updated_at  timestamp with time zone                   NOT NULL,
    constraint unique_user_receiver_msg
        unique (user_id, receiver_id, msg_id)
);

-- INSERT INTO schema_migrations (version, dirty) VALUES (5, false);
