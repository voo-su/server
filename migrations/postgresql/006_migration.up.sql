CREATE TABLE stickers
(
    id         serial primary key,
    name       varchar(50)  default ''::character varying NOT NULL,
    icon       varchar(255) default ''::character varying NOT NULL,
    status     smallint     default 0                     NOT NULL,
    created_at timestamp                                  NOT NULL,
    updated_at timestamp                                  NOT NULL
);

CREATE TABLE sticker_users
(
    id          serial primary key,
    user_id     integer                                    NOT NULL,
    sticker_ids varchar(255) default ''::character varying NOT NULL,
    created_at  timestamp                                  NOT NULL
);

CREATE TABLE sticker_items
(
    id          serial primary key,
    sticker_id  integer      default 0                     NOT NULL,
    user_id     integer      default 0                     NOT NULL,
    description varchar(255) default ''::character varying NOT NULL,
    url         varchar(255) default ''::character varying NOT NULL,
    file_suffix varchar(10)  default ''::character varying NOT NULL,
    file_size   bigint       default 0                     NOT NULL,
    created_at  timestamp                                  NOT NULL,
    updated_at  timestamp                                  NOT NULL
);

-- INSERT INTO schema_migrations (version, dirty) VALUES (6, false);
