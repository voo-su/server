CREATE TABLE splits
(
    id            serial primary key,
    type          smallint     default 1                     NOT NULL,
    drive         smallint     default 1                     NOT NULL,
    upload_id     varchar(100) default ''::character varying NOT NULL,
    user_id       integer      default 0                     NOT NULL,
    original_name varchar(100) default ''::character varying NOT NULL,
    split_index   integer      default 0                     NOT NULL,
    split_num     integer      default 0                     NOT NULL,
    path          varchar(255) default ''::character varying NOT NULL,
    file_ext      varchar(10)  default ''::character varying NOT NULL,
    file_size     integer                                    NOT NULL,
    is_delete     smallint     default 0                     NOT NULL,
    attr          jsonb                                      NOT NULL,
    created_at    timestamp                                  NOT NULL,
    updated_at    timestamp                                  NOT NULL
);

CREATE TABLE bots
(
    id          serial primary key,
    user_id     integer      default 0                     NOT NULL,
    bot_type    integer      default 0,
    name        varchar(255) default ''::character varying NOT NULL,
    description varchar(255) default ''::character varying NOT NULL,
    avatar      varchar(255) default ''::character varying NOT NULL,
    created_at  timestamp                                  NOT NULL
);

-- INSERT INTO schema_migrations (version, dirty) VALUES (2, false);
