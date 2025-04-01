INSERT INTO message_login (message_id, ip_address, user_agent, address, user_id, created_at)
SELECT id                                 AS message_id,
       extra ->> 'ip'                     AS ip_address,
       extra ->> 'agent'                  AS user_agent,
       extra ->> 'address'                AS address,
       user_id,
       (extra ->> 'datetime'):: timestamp AS created_at
FROM messages
WHERE msg_type = 10;

alter table messages drop column file_id;
alter table messages alter column quote_id drop not null;

ALTER TABLE message_read
    ADD COLUMN message_id BIGINT NOT NULL DEFAULT 0;

ALTER TABLE message_read
    ADD CONSTRAINT message_read_message_id_fk
        FOREIGN KEY (message_id) REFERENCES messages (id)
            ON DELETE CASCADE;

-- DROP INDEX IF EXISTS idx_msg_id;
-- ALTER TABLE message_read DROP CONSTRAINT IF EXISTS unique_user_receiver_msg;
-- ALTER TABLE message_read DROP COLUMN IF EXISTS msg_id;
