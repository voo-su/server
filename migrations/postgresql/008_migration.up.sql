create index idx_msg_id on message_read (msg_id);
create index idx_created_at on message_read (created_at);
create index idx_updated_at on message_read (updated_at);

ALTER TABLE users ALTER COLUMN created_at SET DEFAULT now();
ALTER TABLE users ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE bots alter COLUMN created_at SET DEFAULT now();
ALTER TABLE bots ADD creator_id INT DEFAULT NULL;
ALTER TABLE bots ADD COLUMN token VARCHAR(255) UNIQUE NOT NULL;

ALTER TABLE splits ALTER COLUMN upload_id type VARCHAR using upload_id::varchar;
ALTER TABLE splits ALTER COLUMN original_name type VARCHAR using original_name::varchar;

CREATE TABLE push_tokens
(
    id           SERIAL PRIMARY KEY,
    user_id      INT          NOT NULL,
    platform     VARCHAR(255) NOT NULL,
    token        TEXT         NOT NULL,
    web_endpoint TEXT,
    web_p256dh   TEXT,
    web_auth     TEXT,
    is_active    BOOLEAN   DEFAULT TRUE,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

alter table splits alter column file_ext type varchar(255) using file_ext::varchar(255);
alter table splits rename to file_splits;

-- INSERT INTO schema_migrations (version, dirty) VALUES (8, false);
