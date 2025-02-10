CREATE TABLE projects
(
    id         uuid DEFAULT gen_random_uuid(),
    name       VARCHAR(255),
    created_by INTEGER,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_members
(
    id         SERIAL PRIMARY KEY,
    project_id uuid,
    user_id    INTEGER,
    created_by INTEGER,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_task_types
(
    id         SERIAL PRIMARY KEY,
    project_id uuid,
    title      VARCHAR(255),
    created_by INTEGER,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_tasks
(
    id          uuid DEFAULT gen_random_uuid(),
    project_id  uuid,
    type_id     INTEGER,
    title       VARCHAR(255),
    description TEXT,
    assigner_id INTEGER,
    executor_id INTEGER,
    created_by  INTEGER,
    created_at  TIMESTAMPTZ
);

CREATE TABLE project_task_coexecutors
(
    id         SERIAL PRIMARY KEY,
    task_id    uuid,
    member_id  INTEGER,
    created_by INTEGER,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_task_watchers
(
    id         SERIAL PRIMARY KEY,
    task_id    uuid,
    member_id  INTEGER,
    created_by INTEGER,
    created_at TIMESTAMPTZ
);

CREATE TABLE project_task_comments
(
    id           SERIAL PRIMARY KEY,
    task_id      uuid,
    comment_text TEXT,
    created_by   INTEGER,
    created_at   TIMESTAMPTZ
);
