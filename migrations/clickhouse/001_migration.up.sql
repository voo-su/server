CREATE TABLE auth_codes
(
    email         String,
    code          String,
    token         String   DEFAULT '',
    error_message String   DEFAULT '',
    created_at    DateTime DEFAULT now()
) ENGINE = MergeTree()
      ORDER BY (email, created_at);

CREATE TABLE access_logs
(
    remote_addr       String,
    request_uri       String,
    request_query     String,
    response_header   String,
    response_time     DateTime,
    http_user_agent   String,
    response_body     String,
    response_body_raw String,
    request_method    String,
    request_body      String,
    request_time      DateTime,
    request_header    String,
    request_body_raw  String,
    host_name         String,
    server_name       String,
    request_id        String,
    request_duration  String,
    http_status       Int32,
    created_at        DateTime DEFAULT now()
) ENGINE = MergeTree()
      ORDER BY created_at;
