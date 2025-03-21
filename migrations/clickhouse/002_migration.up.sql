CREATE TABLE IF NOT EXISTS loggers
(
    log_message String,
    created_at  DateTime DEFAULT now()
) ENGINE = MergeTree()
      ORDER BY created_at;

CREATE TABLE IF NOT EXISTS access_grpc_logs
(
    full_method String,
    created_at  DateTime DEFAULT now()
) ENGINE = MergeTree()
      ORDER BY created_at;

CREATE TABLE IF NOT EXISTS access_grpc_stream_logs
(
    full_method String,
    created_at  DateTime DEFAULT now()
) ENGINE = MergeTree()
      ORDER BY created_at;
