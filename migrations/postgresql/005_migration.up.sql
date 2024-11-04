create table message_delete
(
    id         serial primary key,
    record_id  integer default 0 not null,
    user_id    integer default 0 not null,
    created_at timestamp         not null
);

create table message_read
(
    id          serial primary key,
    msg_id      varchar(64) default '':: character varying not null,
    user_id     integer     default 0                      not null,
    receiver_id integer     default 0                      not null,
    created_at  timestamp with time zone                   not null,
    updated_at  timestamp with time zone                   not null,
    constraint unique_user_receiver_msg
        unique (user_id, receiver_id, msg_id)
);

INSERT INTO schema_migrations (version, dirty) VALUES (5, false);
