CREATE TABLE contact_folders
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER     DEFAULT 0                      NOT NULL,
    name       VARCHAR(50) DEFAULT '':: CHARACTER VARYING NOT NULL,
    num        INTEGER     DEFAULT 0                      NOT NULL,
    sort       INTEGER     DEFAULT 0                      NOT NULL,
    created_at TIMESTAMP                                  NOT NULL,
    updated_at TIMESTAMP                                  NOT NULL
);

CREATE TABLE group_chat_ads
(
    id            SERIAL PRIMARY KEY,
    group_id      INTEGER     DEFAULT 0                      NOT NULL,
    creator_id    INTEGER     DEFAULT 0                      NOT NULL,
    title         VARCHAR(50) DEFAULT '':: CHARACTER VARYING NOT NULL,
    content       TEXT                                       NOT NULL,
    confirm_users JSONB,
    is_delete     INTEGER     DEFAULT 0                      NOT NULL,
    is_top        INTEGER     DEFAULT 0                      NOT NULL,
    is_confirm    INTEGER     DEFAULT 0                      NOT NULL,
    created_at    TIMESTAMP                                  NOT NULL,
    updated_at    TIMESTAMP                                  NOT NULL,
    deleted_at    TIMESTAMP
);

CREATE TABLE group_chat_requests
(
    id         SERIAL PRIMARY KEY,
    group_id   INTEGER DEFAULT 0 NOT NULL,
    user_id    INTEGER DEFAULT 0 NOT NULL,
    status     INTEGER DEFAULT 1 NOT NULL,
    created_at TIMESTAMP         NOT NULL,
    updated_at TIMESTAMP         NOT NULL
);
