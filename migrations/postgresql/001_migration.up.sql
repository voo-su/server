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

create index idx_msg_id on message_read (msg_id);
create index idx_created_at on message_read (created_at);
create index idx_updated_at on message_read (updated_at);

alter table users alter column created_at set default now();
alter table users alter column updated_at set default now();

alter table bots alter column created_at set default now();
alter table bots add creator_id int default null;

ALTER TABLE bots ADD COLUMN token VARCHAR(255) UNIQUE NOT NULL;
