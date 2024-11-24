CREATE TABLE contact_folders
(
    id         serial primary key,
    user_id    integer     default 0                     NOT NULL,
    name       varchar(50) default ''::character varying NOT NULL,
    num        integer     default 0                     NOT NULL,
    sort       integer     default 0                     NOT NULL,
    created_at timestamp                                 NOT NULL,
    updated_at timestamp                                 NOT NULL
);

CREATE TABLE group_chat_ads
(
    id            serial primary key,
    group_id      integer     default 0                     NOT NULL,
    creator_id    integer     default 0                     NOT NULL,
    title         varchar(50) default ''::character varying NOT NULL,
    content       text                                      NOT NULL,
    confirm_users jsonb,
    is_delete     smallint    default 0                     NOT NULL,
    is_top        smallint    default 0                     NOT NULL,
    is_confirm    smallint    default 0                     NOT NULL,
    created_at    timestamp                                 NOT NULL,
    updated_at    timestamp                                 NOT NULL,
    deleted_at    timestamp
);

CREATE TABLE group_chat_requests
(
    id         serial primary key,
    group_id   integer default 0 NOT NULL,
    user_id    integer default 0 NOT NULL,
    status     integer default 1 NOT NULL,
    created_at timestamp         NOT NULL,
    updated_at timestamp         NOT NULL
);

-- INSERT INTO schema_migrations (version, dirty) VALUES (7, false);
