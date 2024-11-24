CREATE TABLE chats
(
    id          serial primary key,
    dialog_type smallint default 1 NOT NULL,
    user_id     integer  default 0 NOT NULL,
    receiver_id integer  default 0 NOT NULL,
    is_top      smallint default 0 NOT NULL,
    is_disturb  smallint default 0 NOT NULL,
    is_delete   smallint default 0 NOT NULL,
    is_bot      smallint default 0 NOT NULL,
    created_at  timestamp          NOT NULL,
    updated_at  timestamp          NOT NULL
);

create index chats_dialog_type_user_id_receiver_id_idx on chats (dialog_type, user_id, receiver_id);

CREATE TABLE contacts
(
    id         serial primary key,
    user_id    integer     default 0                     NOT NULL,
    friend_id  integer     default 0                     NOT NULL,
    remark     varchar(20) default ''::character varying NOT NULL,
    status     smallint    default 0                     NOT NULL,
    group_id   integer     default 0                     NOT NULL,
    created_at timestamp   default CURRENT_TIMESTAMP     NOT NULL,
    updated_at timestamp   default CURRENT_TIMESTAMP     NOT NULL
);

CREATE TABLE contact_requests
(
    id         serial primary key,
    user_id    integer     default 0                     NOT NULL,
    friend_id  integer     default 0                     NOT NULL,
    remark     varchar(50) default ''::character varying NOT NULL,
    created_at timestamp                                 NOT NULL
);

CREATE TABLE messages
(
    id          bigserial primary key,
    msg_id      varchar(50) default ''::character varying NOT NULL,
    sequence    integer     default 0                     NOT NULL,
    dialog_type smallint    default 1                     NOT NULL,
    msg_type    integer     default 1                     NOT NULL,
    user_id     integer     default 0                     NOT NULL,
    receiver_id integer     default 0                     NOT NULL,
    is_revoke   smallint    default 0                     NOT NULL,
    is_mark     smallint    default 0                     NOT NULL,
    is_read     smallint    default 0                     NOT NULL,
    quote_id    varchar(50)                               NOT NULL,
    content     text,
    extra       jsonb                                     NOT NULL
        constraint dialog_records_extra_check  check (extra IS JSON),
    created_at  timestamp                                 NOT NULL,
    updated_at  timestamp                                 NOT NULL
);

CREATE TABLE users
(
    id         serial primary key,
    email      varchar(255) default ''::character varying NOT NULL,
    username   varchar(255) default ''::character varying NOT NULL,
    name       varchar(255),
    surname    varchar(255),
    avatar     varchar(255) default ''::character varying NOT NULL,
    gender     smallint     default 0                     NOT NULL,
    about      varchar(100) default ''::character varying NOT NULL,
    birthday   varchar(10)  default ''::character varying NOT NULL,
    is_bot     smallint     default 0                     NOT NULL,
    created_at timestamp                                  NOT NULL,
    updated_at timestamp                                  NOT NULL
);

CREATE TABLE user_sessions
(
    id           serial primary key,
    user_id      integer      NOT NULL,
    access_token varchar(255) NOT NULL,
    is_logout    boolean   default false,
    updated_at   timestamp,
    logout_at    timestamp,
    user_ip      inet,
    user_agent   varchar(255),
    created_at   timestamp default CURRENT_TIMESTAMP
);

CREATE TABLE schema_migrations
(
    version bigint  NOT NULL primary key,
    dirty   boolean NOT NULL
);

-- INSERT INTO schema_migrations (version, dirty) VALUES (1, false);
