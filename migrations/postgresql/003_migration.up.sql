CREATE TABLE group_chats
(
    id           serial primary key,
    creator_id   integer      default 0                     NOT NULL,
    type         smallint     default 1                     NOT NULL,
    group_name   varchar(30)  default ''::character varying NOT NULL,
    description  varchar(100) default ''::character varying NOT NULL,
    avatar       varchar(255) default ''::character varying NOT NULL,
    max_num      smallint     default 200                   NOT NULL,
    is_overt     smallint     default 0                     NOT NULL,
    is_mute      smallint     default 0                     NOT NULL,
    is_dismiss   smallint     default 0                     NOT NULL,
    created_at   timestamp                                  NOT NULL,
    updated_at   timestamp                                  NOT NULL,
    dismissed_at timestamp
);

CREATE TABLE group_chat_members
(
    id            serial primary key,
    group_id      integer     default 0                     NOT NULL,
    user_id       integer     default 0                     NOT NULL,
    leader        smallint    default 0                     NOT NULL,
    user_card     varchar(20) default ''::character varying NOT NULL,
    is_quit       smallint    default 0                     NOT NULL,
    is_mute       smallint    default 0                     NOT NULL,
    min_record_id integer     default 0                     NOT NULL,
    join_time     timestamp,
    created_at    timestamp                                 NOT NULL,
    updated_at    timestamp                                 NOT NULL
);

-- INSERT INTO schema_migrations (version, dirty) VALUES (3, false);
