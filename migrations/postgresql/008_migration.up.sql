create index idx_msg_id on message_read (msg_id);
create index idx_created_at on message_read (created_at);
create index idx_updated_at on message_read (updated_at);

ALTER TABLE users ALTER COLUMN created_at SET DEFAULT now();
ALTER TABLE users ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE bots alter COLUMN created_at SET DEFAULT now();
ALTER TABLE bots ADD creator_id INT DEFAULT NULL;
ALTER TABLE bots ADD COLUMN token VARCHAR(255) UNIQUE NOT NULL;

alter table splits alter column upload_id type varchar using upload_id::varchar;
alter table splits alter column original_name type varchar using original_name::varchar;

INSERT INTO schema_migrations (version, dirty) VALUES (8, false);
