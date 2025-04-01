CREATE INDEX idx_msg_id on message_read (msg_id);
CREATE INDEX idx_created_at on message_read (created_at);
CREATE INDEX idx_updated_at on message_read (updated_at);

ALTER TABLE users ALTER COLUMN created_at SET DEFAULT now();
ALTER TABLE users ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE bots alter COLUMN created_at SET DEFAULT now();
ALTER TABLE bots ADD creator_id INT DEFAULT NULL;
ALTER TABLE bots ADD COLUMN token VARCHAR(255) UNIQUE NOT NULL;

ALTER TABLE file_splits ALTER COLUMN upload_id type VARCHAR using upload_id::varchar;
ALTER TABLE file_splits ALTER COLUMN original_name type VARCHAR using original_name::varchar;

CREATE TABLE push_tokens
(
    id           SERIAL PRIMARY KEY,
    user_id      INTEGER      NOT NULL,
    platform     VARCHAR(255) NOT NULL,
    token        TEXT         NOT NULL,
    web_endpoint TEXT,
    web_p256dh   TEXT,
    web_auth     TEXT,
    is_active    BOOLEAN   DEFAULT TRUE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE file_splits ALTER COLUMN file_ext type VARCHAR(255) using file_ext::varchar(255);

ALTER TABLE users ADD notify_chats_mute_until INTEGER DEFAULT 0 NOT NULL;
ALTER TABLE users ADD notify_chats_show_previews BOOLEAN DEFAULT TRUE NOT NULL;
ALTER TABLE users ADD notify_chats_silent BOOLEAN DEFAULT FALSE NOT NULL;

ALTER TABLE users ADD notify_group_mute_until INTEGER DEFAULT 0 NOT NULL;
ALTER TABLE users ADD notify_group_show_previews BOOLEAN DEFAULT TRUE NOT NULL;
ALTER TABLE users ADD notify_group_silent BOOLEAN DEFAULT FALSE NOT NULL;

ALTER TABLE chats ADD notify_mute_until INTEGER DEFAULT 0 NOT NULL;
ALTER TABLE chats ADD notify_show_previews BOOLEAN DEFAULT TRUE NOT NULL;
ALTER TABLE chats ADD notify_silent BOOLEAN DEFAULT FALSE NOT NULL;

ALTER TABLE push_tokens ADD user_session_id INTEGER;


ALTER TABLE group_chat_members RENAME COLUMN min_record_id TO min_message_id;
ALTER TABLE message_votes RENAME COLUMN record_id TO message_id;
ALTER TABLE message_delete RENAME COLUMN record_id TO message_id;

ALTER TABLE users ALTER COLUMN email DROP NOT NULL;
ALTER TABLE users ALTER COLUMN email SET DEFAULT NULL;
ALTER TABLE users ALTER COLUMN about DROP NOT NULL;
ALTER TABLE users ALTER COLUMN about SET DEFAULT NULL;

CREATE UNIQUE INDEX unique_lower_email ON users (LOWER(email)) WHERE email IS NOT NULL;

CREATE TABLE files
(
    id            uuid      DEFAULT gen_random_uuid(),
    original_name TEXT         NOT NULL,
    object_name   TEXT         NOT NULL,
    size          INTEGER      NOT NULL,
    mime_type     VARCHAR(255) NOT NULL,
    created_by    INTEGER      NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

alter table messages add file_id uuid default NULL;

--UPDATE--

CREATE TABLE message_forwarded
(
    id                  SERIAL PRIMARY KEY,
    original_message_id INTEGER REFERENCES messages (id) ON DELETE CASCADE,
    new_message_id      INTEGER REFERENCES messages (id) ON DELETE CASCADE,
    user_id             INTEGER NOT NULL,
    created_at          TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE messages ADD COLUMN reply_to INTEGER REFERENCES messages(id) ON DELETE SET NULL;

CREATE TABLE message_invited_members
(
    id         SERIAL PRIMARY KEY,
    message_id INTEGER REFERENCES messages (id) ON DELETE CASCADE,
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (message_id, user_id)
);

CREATE TABLE message_login
(
    id         SERIAL PRIMARY KEY,
    message_id INTEGER REFERENCES messages (id) ON DELETE CASCADE,
    ip_address VARCHAR,
    user_agent TEXT,
    address    TEXT,
    user_id    INTEGER REFERENCES users (id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE message_media
(
    id         SERIAL PRIMARY KEY,
    message_id INTEGER NOT NULL,
    file_id    UUID    NOT NULL,
    drive      INTEGER,
    duration   INTEGER,
    url        TEXT,
    name       VARCHAR(255),
    size       INTEGER,
    cover      TEXT,
    mime_type  VARCHAR(255),
    width      INTEGER,
    height     INTEGER,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_message FOREIGN KEY (message_id) REFERENCES messages (id)
);

CREATE TABLE message_code
(
    id         serial PRIMARY KEY,
    message_id bigint      NOT NULL REFERENCES messages (id) ON DELETE CASCADE,
    lang       varchar(50) NOT NULL,
    code       text        NOT NULL,
    created_at TIMESTAMP   NOT NULL DEFAULT now()
);

CREATE TABLE message_location
(
    id          serial PRIMARY KEY,
    message_id  bigint           NOT NULL REFERENCES messages (id) ON DELETE CASCADE,
    longitude   double precision NOT NULL,
    latitude    double precision NOT NULL,
    description text,
    created_at  TIMESTAMP        NOT NULL DEFAULT now()
);
