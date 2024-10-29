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
