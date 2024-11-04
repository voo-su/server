create table group_chats
(
    id           serial primary key,
    creator_id   integer      default 0                     not null,
    type         smallint     default 1                     not null,
    group_name   varchar(30)  default ''::character varying not null,
    description  varchar(100) default ''::character varying not null,
    avatar       varchar(255) default ''::character varying not null,
    max_num      smallint     default 200                   not null,
    is_overt     smallint     default 0                     not null,
    is_mute      smallint     default 0                     not null,
    is_dismiss   smallint     default 0                     not null,
    created_at   timestamp                                  not null,
    updated_at   timestamp                                  not null,
    dismissed_at timestamp
);

create table group_chat_members
(
    id            serial primary key,
    group_id      integer     default 0                     not null,
    user_id       integer     default 0                     not null,
    leader        smallint    default 0                     not null,
    user_card     varchar(20) default ''::character varying not null,
    is_quit       smallint    default 0                     not null,
    is_mute       smallint    default 0                     not null,
    min_record_id integer     default 0                     not null,
    join_time     timestamp,
    created_at    timestamp                                 not null,
    updated_at    timestamp                                 not null
);

INSERT INTO schema_migrations (version, dirty) VALUES (3, false);
