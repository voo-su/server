create table chats
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

create index chats_dialog_type_user_id_receiver_id_idx on chats (dialog_type, user_id, receiver_id);

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
        constraint dialog_records_extra_check  check (extra IS JSON),
    created_at  timestamp                                 not null,
    updated_at  timestamp                                 not null
);

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

create table schema_migrations
(
    version bigint  not null primary key,
    dirty   boolean not null
);

INSERT INTO schema_migrations (version, dirty) VALUES (1, false);


CREATE TABLE projects
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255),
    created_by INT,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_members
(
    id         SERIAL PRIMARY KEY,
    project_id INT,
    user_id    INT,
    created_by INT,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_task_types
(
    id         SERIAL PRIMARY KEY,
    project_id INT,
    title      VARCHAR(255),
    created_by INT,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_tasks
(
    id          SERIAL PRIMARY KEY,
    project_id  INT,
    type_id     INT,
    title       VARCHAR(255),
    description TEXT,
    assigner_id INT,
    executor_id INT,
    created_by  INT,
    created_at  TIMESTAMPTZ
);

CREATE TABLE project_task_coexecutors
(
    id         SERIAL PRIMARY KEY,
    task_id    INT,
    member_id  INT,
    created_by INT,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_task_watchers
(
    id         SERIAL PRIMARY KEY,
    task_id    INT,
    member_id  INT,
    created_by INT,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_task_comments
(
    id           SERIAL PRIMARY KEY,
    task_id      INT,
    comment_text TEXT,
    created_by   INT,
    created_at   TIMESTAMPTZ
);
