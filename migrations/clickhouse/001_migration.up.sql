CREATE TABLE auth_codes
(
    email         String,
    code          String,
    token         String   DEFAULT '',
    error_message String   DEFAULT '',
    created_at    DateTime DEFAULT now()
) ENGINE = MergeTree()
      ORDER BY (email, created_at);
