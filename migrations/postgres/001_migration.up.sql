CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    email      VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    username   VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    name       VARCHAR(255),
    surname    VARCHAR(255),
    avatar     VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    gender     INTEGER      DEFAULT 0                     NOT NULL,
    about      VARCHAR(100) DEFAULT ''::CHARACTER VARYING NOT NULL,
    birthday   VARCHAR(10)  DEFAULT ''::CHARACTER VARYING NOT NULL,
    is_bot     INTEGER      DEFAULT 0                     NOT NULL,
    created_at TIMESTAMP                                  NOT NULL,
    updated_at TIMESTAMP                                  NOT NULL
);

CREATE TABLE user_sessions
(
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER      NOT NULL,
    access_token VARCHAR(255) NOT NULL,
    is_logout    BOOLEAN   DEFAULT false,
    updated_at   TIMESTAMP,
    logout_at    TIMESTAMP,
    user_ip      INET,
    user_agent   VARCHAR(255),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE contacts
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER   DEFAULT 0                 NOT NULL,
    friend_id  INTEGER   DEFAULT 0                 NOT NULL,
    status     INTEGER   DEFAULT 0                 NOT NULL,
    group_id   INTEGER   DEFAULT 0                 NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE contact_requests
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER DEFAULT 0 NOT NULL,
    friend_id  INTEGER DEFAULT 0 NOT NULL,
    created_at TIMESTAMP         NOT NULL
);
