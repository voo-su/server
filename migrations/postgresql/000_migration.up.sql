create table users
(
    id         serial primary key,
    email      varchar(255) default ''::character varying not null,
    username   varchar(255) default ''::character varying not null,
    name       varchar(255),
    surname    varchar(255),
    avatar     varchar(255) default ''::character varying not null,
    gender     smallint     default 0                     not null,
    about      varchar(100) default ''::character varying not null,
    birthday   varchar(10)  default ''::character varying not null,
    is_bot     smallint     default 0                     not null,
    created_at timestamp                                  not null,
    updated_at timestamp                                  not null
);

create table user_sessions
(
    id           serial primary key,
    user_id      integer      not null,
    access_token varchar(255) not null,
    is_logout    boolean   default false,
    updated_at   timestamp,
    logout_at    timestamp,
    user_ip      inet,
    user_agent   varchar(255),
    created_at   timestamp default CURRENT_TIMESTAMP
);

create table stickers
(
    id         serial primary key,
    name       varchar(50)  default ''::character varying not null,
    icon       varchar(255) default ''::character varying not null,
    status     smallint     default 0                     not null,
    created_at timestamp                                  not null,
    updated_at timestamp                                  not null
);

create table sticker_user
(
    id          serial primary key,
    user_id     integer                                    not null,
    sticker_ids varchar(255) default ''::character varying not null,
    created_at  timestamp                                  not null
);

create table sticker_items
(
    id          serial primary key,
    sticker_id  integer      default 0                     not null,
    user_id     integer      default 0                     not null,
    description varchar(255) default ''::character varying not null,
    url         varchar(255) default ''::character varying not null,
    file_suffix varchar(10)  default ''::character varying not null,
    file_size   bigint       default 0                     not null,
    created_at  timestamp                                  not null,
    updated_at  timestamp                                  not null
);

create table splits
(
    id            serial primary key,
    type          smallint     default 1                     not null,
    drive         smallint     default 1                     not null,
    upload_id     varchar(100) default ''::character varying not null,
    user_id       integer      default 0                     not null,
    original_name varchar(100) default ''::character varying not null,
    split_index   integer      default 0                     not null,
    split_num     integer      default 0                     not null,
    path          varchar(255) default ''::character varying not null,
    file_ext      varchar(10)  default ''::character varying not null,
    file_size     integer                                    not null,
    is_delete     smallint     default 0                     not null,
    attr          jsonb                                      not null,
    created_at    timestamp                                  not null,
    updated_at    timestamp                                  not null
);

create table messages
(
    id          bigserial primary key,
    msg_id      varchar(50) default ''::character varying not null,
    sequence    integer     default 0                     not null,
    dialog_type smallint    default 1                     not null,
    msg_type    integer     default 1                     not null,
    user_id     integer     default 0                     not null,
    receiver_id integer     default 0                     not null,
    is_revoke   smallint    default 0                     not null,
    is_mark     smallint    default 0                     not null,
    is_read     smallint    default 0                     not null,
    quote_id    varchar(50)                               not null,
    content     text,
    extra       jsonb                                     not null
        constraint dialog_records_extra_check
            check (extra IS JSON),
    created_at timestamp not null,
    updated_at timestamp not null
);

create table message_votes
(
    id            serial primary key,
    record_id     integer      default 0                     not null,
    user_id       integer      default 0                     not null,
    title         varchar(255) default ''::character varying not null,
    answer_mode   smallint     default 0                     not null,
    answer_option jsonb                                      not null,
    answer_num    smallint     default 0                     not null,
    answered_num  smallint     default 0                     not null,
    is_anonymous  smallint     default 0                     not null,
    status        smallint     default 0                     not null,
    created_at    timestamp                                  not null,
    updated_at    timestamp                                  not null
);

create table message_vote_answers
(
    id         serial primary key,
    vote_id    integer default 0          not null,
    user_id    integer default 0          not null,
    option     char    default ''::bpchar not null,
    created_at timestamp                  not null
);

create table message_delete
(
    id         serial primary key,
    record_id  integer default 0 not null,
    user_id    integer default 0 not null,
    created_at timestamp         not null
);

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

create table group_chat_requests
(
    id         serial primary key,
    group_id   integer default 0 not null,
    user_id    integer default 0 not null,
    status     integer default 1 not null,
    created_at timestamp         not null,
    updated_at timestamp         not null
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

create table dialogs
(
    id          serial primary key,
    dialog_type smallint default 1 not null,
    user_id     integer  default 0 not null,
    receiver_id integer  default 0 not null,
    is_top      smallint default 0 not null,
    is_disturb  smallint default 0 not null,
    is_delete   smallint default 0 not null,
    is_bot      smallint default 0 not null,
    created_at  timestamp          not null,
    updated_at  timestamp          not null
);

create index dialogs_dialog_type_user_id_receiver_id_idx on dialogs (dialog_type, user_id, receiver_id);

create table contacts
(
    id         serial primary key,
    user_id    integer     default 0                     not null,
    friend_id  integer     default 0                     not null,
    remark     varchar(20) default ''::character varying not null,
    status     smallint    default 0                     not null,
    group_id   integer     default 0                     not null,
    created_at timestamp   default CURRENT_TIMESTAMP     not null,
    updated_at timestamp   default CURRENT_TIMESTAMP     not null
);

create table contact_requests
(
    id         serial primary key,
    user_id    integer     default 0                     not null,
    friend_id  integer     default 0                     not null,
    remark     varchar(50) default ''::character varying not null,
    created_at timestamp                                 not null
);

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

create table bots
(
    id          serial primary key,
    user_id     integer      default 0                     not null,
    bot_type    integer      default 0,
    name        varchar(255) default ''::character varying not null,
    description varchar(255) default ''::character varying not null,
    avatar      varchar(255) default ''::character varying not null,
    created_at  timestamp                                  not null
);

create table schema_migrations
(
    version bigint  not null primary key,
    dirty   boolean not null
);

INSERT INTO schema_migrations (version, dirty) VALUES (1, false);
