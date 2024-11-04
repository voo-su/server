create table splits
(
    id            serial primary key,
    type          smallint     default 1                     not null,
    drive         smallint     default 1                     not null,
    upload_id     varchar(100) default ''::character varying not null,
    user_id       integer      default 0                     not null,
    original_name varchar(100) default ''::character varying not null,
    split_index   integer      default 0                     not null,
    split_num     integer      default 0                     not null,
    path          varchar(255) default ''::character varying not null,
    file_ext      varchar(10)  default ''::character varying not null,
    file_size     integer                                    not null,
    is_delete     smallint     default 0                     not null,
    attr          jsonb                                      not null,
    created_at    timestamp                                  not null,
    updated_at    timestamp                                  not null
);

create table bots
(
    id          serial primary key,
    user_id     integer      default 0                     not null,
    bot_type    integer      default 0,
    name        varchar(255) default ''::character varying not null,
    description varchar(255) default ''::character varying not null,
    avatar      varchar(255) default ''::character varying not null,
    created_at  timestamp                                  not null
);

INSERT INTO schema_migrations (version, dirty) VALUES (2, false);
