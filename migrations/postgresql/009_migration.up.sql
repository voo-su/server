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

--UPDATE--
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
