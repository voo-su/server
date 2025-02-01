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
    created_at timestamp    default now()                 NOT NULL,
    updated_at timestamp    default now()                 NOT NULL
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

CREATE TABLE stickers
(
    id         serial primary key,
    name       varchar(50)  default ''::character varying NOT NULL,
    icon       varchar(255) default ''::character varying NOT NULL,
    status     smallint     default 0                     NOT NULL,
    created_at timestamp                                  NOT NULL,
    updated_at timestamp                                  NOT NULL
);

CREATE TABLE sticker_users
(
    id          serial primary key,
    user_id     integer                                    NOT NULL,
    sticker_ids varchar(255) default ''::character varying NOT NULL,
    created_at  timestamp                                  NOT NULL
);

CREATE TABLE sticker_items
(
    id          serial primary key,
    sticker_id  integer      default 0                     NOT NULL,
    user_id     integer      default 0                     NOT NULL,
    description varchar(255) default ''::character varying NOT NULL,
    url         varchar(255) default ''::character varying NOT NULL,
    file_suffix varchar(10)  default ''::character varying NOT NULL,
    file_size   bigint       default 0                     NOT NULL,
    created_at  timestamp                                  NOT NULL,
    updated_at  timestamp                                  NOT NULL
);

CREATE TABLE splits
(
    id            serial primary key,
    type          smallint     default 1                     NOT NULL,
    drive         smallint     default 1                     NOT NULL,
    upload_id     varchar      default ''::character varying NOT NULL,
    user_id       integer      default 0                     NOT NULL,
    original_name varchar      default ''::character varying NOT NULL,
    split_index   integer      default 0                     NOT NULL,
    split_num     integer      default 0                     NOT NULL,
    path          varchar(255) default ''::character varying NOT NULL,
    file_ext      varchar(10)  default ''::character varying NOT NULL,
    file_size     integer                                    NOT NULL,
    is_delete     smallint     default 0                     NOT NULL,
    attr          jsonb                                      NOT NULL,
    created_at    timestamp                                  NOT NULL,
    updated_at    timestamp                                  NOT NULL
);

CREATE TABLE push_tokens
(
    id           serial primary key,
    user_id      integer      NOT NULL,
    platform     varchar(255) NOT NULL,
    token        text         NOT NULL,
    web_endpoint text,
    web_p256dh   text,
    web_auth     text,
    is_active    boolean   default true,
    created_at   timestamp default CURRENT_TIMESTAMP
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
        constraint dialog_records_extra_check check (extra IS JSON),
    created_at  timestamp                                 NOT NULL,
    updated_at  timestamp                                 NOT NULL
);

CREATE TABLE message_votes
(
    id            serial primary key,
    record_id     integer      default 0                     NOT NULL,
    user_id       integer      default 0                     NOT NULL,
    title         varchar(255) default ''::character varying NOT NULL,
    answer_mode   smallint     default 0                     NOT NULL,
    answer_option jsonb                                      NOT NULL,
    answer_num    smallint     default 0                     NOT NULL,
    answered_num  smallint     default 0                     NOT NULL,
    is_anonymous  smallint     default 0                     NOT NULL,
    status        smallint     default 0                     NOT NULL,
    created_at    timestamp                                  NOT NULL,
    updated_at    timestamp                                  NOT NULL,
    new_column    integer
);

CREATE TABLE message_vote_answers
(
    id         serial primary key,
    vote_id    integer default 0          NOT NULL,
    user_id    integer default 0          NOT NULL,
    option     char    default ''::bpchar NOT NULL,
    created_at timestamp                  NOT NULL,
    new_column integer
);

CREATE TABLE message_read
(
    id          serial primary key,
    msg_id      varchar(64) default ''::character varying NOT NULL,
    user_id     integer     default 0                     NOT NULL,
    receiver_id integer     default 0                     NOT NULL,
    created_at  timestamp with time zone                  NOT NULL,
    updated_at  timestamp with time zone                  NOT NULL,
    constraint unique_user_receiver_msg
        unique (user_id, receiver_id, msg_id)
);

create index idx_msg_id on message_read (msg_id);
create index idx_created_at on message_read (created_at);
create index idx_updated_at on message_read (updated_at);

CREATE TABLE message_delete
(
    id         serial primary key,
    record_id  integer default 0 NOT NULL,
    user_id    integer default 0 NOT NULL,
    created_at timestamp         NOT NULL
);

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

CREATE TABLE group_chat_requests
(
    id         serial primary key,
    group_id   integer default 0 NOT NULL,
    user_id    integer default 0 NOT NULL,
    status     integer default 1 NOT NULL,
    created_at timestamp         NOT NULL,
    updated_at timestamp         NOT NULL
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
    deleted_at    timestamp,
    new_column    integer
);

CREATE TABLE contacts
(
    id         serial primary key,
    user_id    integer     default 0                     NOT NULL,
    friend_id  integer     default 0                     NOT NULL,
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
    created_at timestamp                                 NOT NULL
);

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

CREATE TABLE bots
(
    id          serial primary key,
    user_id     integer      default 0                     NOT NULL,
    bot_type    integer      default 0,
    name        varchar(255) default ''::character varying NOT NULL,
    description varchar(255) default ''::character varying NOT NULL,
    avatar      varchar(255) default ''::character varying NOT NULL,
    created_at  timestamp    default now()                 NOT NULL,
    token       varchar(255)                               NOT NULL unique,
    creator_id  integer
);

CREATE TABLE schema_migrations
(
    version bigint  NOT NULL primary key,
    dirty   boolean NOT NULL
);

INSERT INTO schema_migrations (version, dirty) VALUES (8, false);
