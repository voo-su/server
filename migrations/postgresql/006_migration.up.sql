create table stickers
(
    id         serial primary key,
    name       varchar(50)  default ''::character varying not null,
    icon       varchar(255) default ''::character varying not null,
    status     smallint     default 0                     not null,
    created_at timestamp                                  not null,
    updated_at timestamp                                  not null
);

create table sticker_users
(
    id          serial primary key,
    user_id     integer                                    not null,
    sticker_ids varchar(255) default ''::character varying not null,
    created_at  timestamp                                  not null
);

create table sticker_items
(
    id          serial primary key,
    sticker_id  integer      default 0                     not null,
    user_id     integer      default 0                     not null,
    description varchar(255) default ''::character varying not null,
    url         varchar(255) default ''::character varying not null,
    file_suffix varchar(10)  default ''::character varying not null,
    file_size   bigint       default 0                     not null,
    created_at  timestamp                                  not null,
    updated_at  timestamp                                  not null
);

INSERT INTO schema_migrations (version, dirty) VALUES (6, false);
