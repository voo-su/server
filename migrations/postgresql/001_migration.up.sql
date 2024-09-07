CREATE TABLE projects
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255),
    created_by INT,
    created_at TIMESTAMPTZ
);
