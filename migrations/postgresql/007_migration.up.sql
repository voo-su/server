CREATE TABLE stickers
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50)  DEFAULT ''::CHARACTER VARYING NOT NULL,
    icon       VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    status     INTEGER      DEFAULT 0                     NOT NULL,
    created_at TIMESTAMP                                  NOT NULL,
    updated_at TIMESTAMP                                  NOT NULL
);

CREATE TABLE sticker_users
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER                                    NOT NULL,
    sticker_ids VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    created_at  TIMESTAMP                                  NOT NULL
);

CREATE TABLE sticker_items
(
    id          SERIAL PRIMARY KEY,
    sticker_id  INTEGER      DEFAULT 0                     NOT NULL,
    user_id     INTEGER      DEFAULT 0                     NOT NULL,
    description VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    url         VARCHAR(255) DEFAULT ''::CHARACTER VARYING NOT NULL,
    file_suffix VARCHAR(10)  DEFAULT ''::CHARACTER VARYING NOT NULL,
    file_size   BIGINT       DEFAULT 0                     NOT NULL,
    created_at  TIMESTAMP                                  NOT NULL,
    updated_at  TIMESTAMP                                  NOT NULL
);
