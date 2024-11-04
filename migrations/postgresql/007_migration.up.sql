create table contact_folders
(
    id         serial primary key,
    user_id    integer     default 0                     not null,
    name       varchar(50) default ''::character varying not null,
    num        integer     default 0                     not null,
    sort       integer     default 0                     not null,
    created_at timestamp                                 not null,
    updated_at timestamp                                 not null
);

create table group_chat_ads
(
    id            serial primary key,
    group_id      integer     default 0                     not null,
    creator_id    integer     default 0                     not null,
    title         varchar(50) default ''::character varying not null,
    content       text                                      not null,
    confirm_users jsonb,
    is_delete     smallint    default 0                     not null,
    is_top        smallint    default 0                     not null,
    is_confirm    smallint    default 0                     not null,
    created_at    timestamp                                 not null,
    updated_at    timestamp                                 not null,
    deleted_at    timestamp
);

create table group_chat_requests
(
    id         serial primary key,
    group_id   integer default 0 not null,
    user_id    integer default 0 not null,
    status     integer default 1 not null,
    created_at timestamp         not null,
    updated_at timestamp         not null
);

INSERT INTO schema_migrations (version, dirty) VALUES (7, false);
