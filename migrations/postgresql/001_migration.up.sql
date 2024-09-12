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
    created_by  INT,
    created_at  TIMESTAMPTZ
);

CREATE TABLE project_task_comments
(
    id           SERIAL PRIMARY KEY,
    task_id      INT,
    comment_text TEXT,
    created_by   INT,
    created_at   TIMESTAMPTZ
);
