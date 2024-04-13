--
-- PostgreSQL database dump
--

-- Dumped from database version 16.0 (Ubuntu 16.0-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.0 (Ubuntu 16.0-1.pgdg22.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: bots; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bots
(
    id          integer                                              NOT NULL,
    user_id     integer                DEFAULT 0                     NOT NULL,
    bot_type    integer                DEFAULT 0,
    name        character varying(255) DEFAULT ''::character varying NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    avatar      character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at  timestamp without time zone                          NOT NULL
);

ALTER TABLE public.bots
    OWNER TO postgres;

--
-- Name: bots_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bots_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.bots_id_seq OWNER TO postgres;

--
-- Name: bots_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bots_id_seq OWNED BY public.bots.id;


--
-- Name: contact_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contact_groups
(
    id         integer                                             NOT NULL,
    user_id    integer               DEFAULT 0                     NOT NULL,
    name       character varying(50) DEFAULT ''::character varying NOT NULL,
    num        integer               DEFAULT 0                     NOT NULL,
    sort       integer               DEFAULT 0                     NOT NULL,
    created_at timestamp without time zone                         NOT NULL,
    updated_at timestamp without time zone                         NOT NULL
);


ALTER TABLE public.contact_groups
    OWNER TO postgres;

--
-- Name: contact_groups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contact_groups_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.contact_groups_id_seq OWNER TO postgres;

--
-- Name: contact_groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contact_groups_id_seq OWNED BY public.contact_groups.id;


--
-- Name: contact_requests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contact_requests
(
    id         integer                                             NOT NULL,
    user_id    integer               DEFAULT 0                     NOT NULL,
    friend_id  integer               DEFAULT 0                     NOT NULL,
    remark     character varying(50) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone                         NOT NULL
);


ALTER TABLE public.contact_requests
    OWNER TO postgres;

--
-- Name: contact_requests_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contact_requests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.contact_requests_id_seq OWNER TO postgres;

--
-- Name: contact_requests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contact_requests_id_seq OWNED BY public.contact_requests.id;


--
-- Name: contacts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contacts
(
    id         integer                                                   NOT NULL,
    user_id    integer                     DEFAULT 0                     NOT NULL,
    friend_id  integer                     DEFAULT 0                     NOT NULL,
    remark     character varying(20)       DEFAULT ''::character varying NOT NULL,
    status     smallint                    DEFAULT 0                     NOT NULL,
    group_id   integer                     DEFAULT 0                     NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP     NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP     NOT NULL
);


ALTER TABLE public.contacts
    OWNER TO postgres;

--
-- Name: contacts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contacts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.contacts_id_seq OWNER TO postgres;

--
-- Name: contacts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contacts_id_seq OWNED BY public.contacts.id;


--
-- Name: dialogs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dialogs
(
    id          integer                     NOT NULL,
    dialog_type smallint DEFAULT 1          NOT NULL,
    user_id     integer  DEFAULT 0          NOT NULL,
    receiver_id integer  DEFAULT 0          NOT NULL,
    is_top      smallint DEFAULT 0          NOT NULL,
    is_disturb  smallint DEFAULT 0          NOT NULL,
    is_delete   smallint DEFAULT 0          NOT NULL,
    is_bot      smallint DEFAULT 0          NOT NULL,
    created_at  timestamp without time zone NOT NULL,
    updated_at  timestamp without time zone NOT NULL
);


ALTER TABLE public.dialogs
    OWNER TO postgres;

--
-- Name: dialogs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dialogs_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.dialogs_id_seq OWNER TO postgres;

--
-- Name: dialogs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dialogs_id_seq OWNED BY public.dialogs.id;


--
-- Name: group_chat_members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.group_chat_members
(
    id            integer                                             NOT NULL,
    group_id      integer               DEFAULT 0                     NOT NULL,
    user_id       integer               DEFAULT 0                     NOT NULL,
    leader        smallint              DEFAULT 0                     NOT NULL,
    user_card     character varying(20) DEFAULT ''::character varying NOT NULL,
    is_quit       smallint              DEFAULT 0                     NOT NULL,
    is_mute       smallint              DEFAULT 0                     NOT NULL,
    min_record_id integer               DEFAULT 0                     NOT NULL,
    join_time     timestamp without time zone,
    created_at    timestamp without time zone                         NOT NULL,
    updated_at    timestamp without time zone                         NOT NULL
);


ALTER TABLE public.group_chat_members
    OWNER TO postgres;

--
-- Name: group_chat_members_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.group_chat_members_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.group_chat_members_id_seq OWNER TO postgres;

--
-- Name: group_chat_members_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.group_chat_members_id_seq OWNED BY public.group_chat_members.id;


--
-- Name: group_chat_notice; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.group_chat_notice
(
    id            integer                                             NOT NULL,
    group_id      integer               DEFAULT 0                     NOT NULL,
    creator_id    integer               DEFAULT 0                     NOT NULL,
    title         character varying(50) DEFAULT ''::character varying NOT NULL,
    content       text                                                NOT NULL,
    confirm_users jsonb,
    is_delete     smallint              DEFAULT 0                     NOT NULL,
    is_top        smallint              DEFAULT 0                     NOT NULL,
    is_confirm    smallint              DEFAULT 0                     NOT NULL,
    created_at    timestamp without time zone                         NOT NULL,
    updated_at    timestamp without time zone                         NOT NULL,
    deleted_at    timestamp without time zone,
    new_column    integer
);


ALTER TABLE public.group_chat_notice
    OWNER TO postgres;

--
-- Name: group_chat_notice_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.group_chat_notice_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.group_chat_notice_id_seq OWNER TO postgres;

--
-- Name: group_chat_notice_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.group_chat_notice_id_seq OWNED BY public.group_chat_notice.id;


--
-- Name: group_chat_requests; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.group_chat_requests
(
    id         integer                                              NOT NULL,
    group_id   integer                DEFAULT 0                     NOT NULL,
    user_id    integer                DEFAULT 0                     NOT NULL,
    status     integer                DEFAULT 1                     NOT NULL,
    remark     character varying(255) DEFAULT ''::character varying NOT NULL,
    reason     character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone                          NOT NULL,
    updated_at timestamp without time zone                          NOT NULL
);


ALTER TABLE public.group_chat_requests
    OWNER TO postgres;

--
-- Name: group_chat_requests_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.group_chat_requests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.group_chat_requests_id_seq OWNER TO postgres;

--
-- Name: group_chat_requests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.group_chat_requests_id_seq OWNED BY public.group_chat_requests.id;


--
-- Name: group_chats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.group_chats
(
    id           integer                                              NOT NULL,
    creator_id   integer                DEFAULT 0                     NOT NULL,
    type         smallint               DEFAULT 1                     NOT NULL,
    group_name   character varying(30)  DEFAULT ''::character varying NOT NULL,
    description  character varying(100) DEFAULT ''::character varying NOT NULL,
    avatar       character varying(255) DEFAULT ''::character varying NOT NULL,
    max_num      smallint               DEFAULT 200                   NOT NULL,
    is_overt     smallint               DEFAULT 0                     NOT NULL,
    is_mute      smallint               DEFAULT 0                     NOT NULL,
    is_dismiss   smallint               DEFAULT 0                     NOT NULL,
    created_at   timestamp without time zone                          NOT NULL,
    updated_at   timestamp without time zone                          NOT NULL,
    dismissed_at timestamp without time zone
);


ALTER TABLE public.group_chats
    OWNER TO postgres;

--
-- Name: group_chats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.group_chats_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.group_chats_id_seq OWNER TO postgres;

--
-- Name: group_chats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.group_chats_id_seq OWNED BY public.group_chats.id;


--
-- Name: message_delete; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.message_delete
(
    id         integer                     NOT NULL,
    record_id  integer DEFAULT 0           NOT NULL,
    user_id    integer DEFAULT 0           NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.message_delete
    OWNER TO postgres;

--
-- Name: message_delete_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.message_delete_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.message_delete_id_seq OWNER TO postgres;

--
-- Name: message_delete_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.message_delete_id_seq OWNED BY public.message_delete.id;


--
-- Name: message_vote_answers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.message_vote_answers
(
    id         integer                         NOT NULL,
    vote_id    integer      DEFAULT 0          NOT NULL,
    user_id    integer      DEFAULT 0          NOT NULL,
    option     character(1) DEFAULT ''::bpchar NOT NULL,
    created_at timestamp without time zone     NOT NULL,
    new_column integer
);


ALTER TABLE public.message_vote_answers
    OWNER TO postgres;

--
-- Name: message_vote_answers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.message_vote_answers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.message_vote_answers_id_seq OWNER TO postgres;

--
-- Name: message_vote_answers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.message_vote_answers_id_seq OWNED BY public.message_vote_answers.id;


--
-- Name: message_votes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.message_votes
(
    id            integer                                              NOT NULL,
    record_id     integer                DEFAULT 0                     NOT NULL,
    user_id       integer                DEFAULT 0                     NOT NULL,
    title         character varying(255) DEFAULT ''::character varying NOT NULL,
    answer_mode   smallint               DEFAULT 0                     NOT NULL,
    answer_option jsonb                                                NOT NULL,
    answer_num    smallint               DEFAULT 0                     NOT NULL,
    answered_num  smallint               DEFAULT 0                     NOT NULL,
    is_anonymous  smallint               DEFAULT 0                     NOT NULL,
    status        smallint               DEFAULT 0                     NOT NULL,
    created_at    timestamp without time zone                          NOT NULL,
    updated_at    timestamp without time zone                          NOT NULL,
    new_column    integer
);


ALTER TABLE public.message_votes
    OWNER TO postgres;

--
-- Name: message_votes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.message_votes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.message_votes_id_seq OWNER TO postgres;

--
-- Name: message_votes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.message_votes_id_seq OWNED BY public.message_votes.id;


--
-- Name: messages; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.messages
(
    id          bigint                                              NOT NULL,
    msg_id      character varying(50) DEFAULT ''::character varying NOT NULL,
    sequence    integer               DEFAULT 0                     NOT NULL,
    dialog_type smallint              DEFAULT 1                     NOT NULL,
    msg_type    integer               DEFAULT 1                     NOT NULL,
    user_id     integer               DEFAULT 0                     NOT NULL,
    receiver_id integer               DEFAULT 0                     NOT NULL,
    is_revoke   smallint              DEFAULT 0                     NOT NULL,
    is_mark     smallint              DEFAULT 0                     NOT NULL,
    is_read     smallint              DEFAULT 0                     NOT NULL,
    quote_id    character varying(50)                               NOT NULL,
    content     text,
    extra       jsonb                                               NOT NULL,
    created_at  timestamp without time zone                         NOT NULL,
    updated_at  timestamp without time zone                         NOT NULL,
    CONSTRAINT dialog_records_extra_check CHECK ((extra IS JSON))
);


ALTER TABLE public.messages
    OWNER TO postgres;

--
-- Name: messages_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.messages_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.messages_id_seq OWNER TO postgres;

--
-- Name: messages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.messages_id_seq OWNED BY public.messages.id;


--
-- Name: splits; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.splits
(
    id            integer                                              NOT NULL,
    type          smallint               DEFAULT 1                     NOT NULL,
    drive         smallint               DEFAULT 1                     NOT NULL,
    upload_id     character varying(100) DEFAULT ''::character varying NOT NULL,
    user_id       integer                DEFAULT 0                     NOT NULL,
    original_name character varying(100) DEFAULT ''::character varying NOT NULL,
    split_index   integer                DEFAULT 0                     NOT NULL,
    split_num     integer                DEFAULT 0                     NOT NULL,
    path          character varying(255) DEFAULT ''::character varying NOT NULL,
    file_ext      character varying(10)  DEFAULT ''::character varying NOT NULL,
    file_size     integer                                              NOT NULL,
    is_delete     smallint               DEFAULT 0                     NOT NULL,
    attr          jsonb                                                NOT NULL,
    created_at    timestamp without time zone                          NOT NULL,
    updated_at    timestamp without time zone                          NOT NULL
);


ALTER TABLE public.splits
    OWNER TO postgres;

--
-- Name: splits_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.splits_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.splits_id_seq OWNER TO postgres;

--
-- Name: splits_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.splits_id_seq OWNED BY public.splits.id;


--
-- Name: sticker_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sticker_items
(
    id          integer                                              NOT NULL,
    sticker_id  integer                DEFAULT 0                     NOT NULL,
    user_id     integer                DEFAULT 0                     NOT NULL,
    description character varying(20)  DEFAULT ''::character varying NOT NULL,
    url         character varying(255) DEFAULT ''::character varying NOT NULL,
    file_suffix character varying(10)  DEFAULT ''::character varying NOT NULL,
    file_size   bigint                 DEFAULT 0                     NOT NULL,
    created_at  timestamp without time zone                          NOT NULL,
    updated_at  timestamp without time zone                          NOT NULL
);


ALTER TABLE public.sticker_items
    OWNER TO postgres;

--
-- Name: sticker_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sticker_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sticker_items_id_seq OWNER TO postgres;

--
-- Name: sticker_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sticker_items_id_seq OWNED BY public.sticker_items.id;


--
-- Name: sticker_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sticker_user
(
    id          integer                                              NOT NULL,
    user_id     integer                                              NOT NULL,
    sticker_ids character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at  timestamp without time zone                          NOT NULL
);


ALTER TABLE public.sticker_user
    OWNER TO postgres;

--
-- Name: sticker_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sticker_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sticker_user_id_seq OWNER TO postgres;

--
-- Name: sticker_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sticker_user_id_seq OWNED BY public.sticker_user.id;


--
-- Name: stickers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stickers
(
    id         integer                                              NOT NULL,
    name       character varying(50)  DEFAULT ''::character varying NOT NULL,
    icon       character varying(255) DEFAULT ''::character varying NOT NULL,
    status     smallint               DEFAULT 0                     NOT NULL,
    created_at timestamp without time zone                          NOT NULL,
    updated_at timestamp without time zone                          NOT NULL
);


ALTER TABLE public.stickers
    OWNER TO postgres;

--
-- Name: stickers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.stickers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.stickers_id_seq OWNER TO postgres;

--
-- Name: stickers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.stickers_id_seq OWNED BY public.stickers.id;


--
-- Name: user_sessions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_sessions
(
    id           integer                NOT NULL,
    user_id      integer                NOT NULL,
    access_token character varying(255) NOT NULL,
    is_logout    boolean                     DEFAULT false,
    updated_at   timestamp without time zone,
    logout_at    timestamp without time zone,
    user_ip      inet,
    user_agent   character varying(255),
    created_at   timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.user_sessions
    OWNER TO postgres;

--
-- Name: user_sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_sessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.user_sessions_id_seq OWNER TO postgres;

--
-- Name: user_sessions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_sessions_id_seq OWNED BY public.user_sessions.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users
(
    id         integer                                              NOT NULL,
    email      character varying(255) DEFAULT ''::character varying NOT NULL,
    username   character varying(255) DEFAULT ''::character varying NOT NULL,
    name       character varying(255),
    surname    character varying(255),
    avatar     character varying(255) DEFAULT ''::character varying NOT NULL,
    gender     smallint               DEFAULT 0                     NOT NULL,
    about      character varying(100) DEFAULT ''::character varying NOT NULL,
    birthday   character varying(10)  DEFAULT ''::character varying NOT NULL,
    is_bot     smallint               DEFAULT 0                     NOT NULL,
    created_at timestamp without time zone                          NOT NULL,
    updated_at timestamp without time zone                          NOT NULL
);


ALTER TABLE public.users
    OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: bots id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bots
ALTER COLUMN id SET DEFAULT nextval('public.bots_id_seq'::regclass);


--
-- Name: contact_groups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contact_groups
    ALTER COLUMN id SET DEFAULT nextval('public.contact_groups_id_seq'::regclass);


--
-- Name: contact_requests id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contact_requests
    ALTER COLUMN id SET DEFAULT nextval('public.contact_requests_id_seq'::regclass);


--
-- Name: contacts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contacts
    ALTER COLUMN id SET DEFAULT nextval('public.contacts_id_seq'::regclass);


--
-- Name: dialogs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dialogs
    ALTER COLUMN id SET DEFAULT nextval('public.dialogs_id_seq'::regclass);


--
-- Name: group_chat_members id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat_members
    ALTER COLUMN id SET DEFAULT nextval('public.group_chat_members_id_seq'::regclass);


--
-- Name: group_chat_notice id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat_notice
    ALTER COLUMN id SET DEFAULT nextval('public.group_chat_notice_id_seq'::regclass);


--
-- Name: group_chat_requests id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat_requests
    ALTER COLUMN id SET DEFAULT nextval('public.group_chat_requests_id_seq'::regclass);


--
-- Name: group_chats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chats
    ALTER COLUMN id SET DEFAULT nextval('public.group_chats_id_seq'::regclass);


--
-- Name: message_delete id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message_delete
    ALTER COLUMN id SET DEFAULT nextval('public.message_delete_id_seq'::regclass);


--
-- Name: message_vote_answers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message_vote_answers
    ALTER COLUMN id SET DEFAULT nextval('public.message_vote_answers_id_seq'::regclass);


--
-- Name: message_votes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message_votes
    ALTER COLUMN id SET DEFAULT nextval('public.message_votes_id_seq'::regclass);


--
-- Name: messages id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ALTER COLUMN id SET DEFAULT nextval('public.messages_id_seq'::regclass);


--
-- Name: splits id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.splits
    ALTER COLUMN id SET DEFAULT nextval('public.splits_id_seq'::regclass);


--
-- Name: sticker_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sticker_items
ALTER COLUMN id SET DEFAULT nextval('public.sticker_items_id_seq'::regclass);


--
-- Name: sticker_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sticker_user
ALTER COLUMN id SET DEFAULT nextval('public.sticker_user_id_seq'::regclass);


--
-- Name: stickers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stickers
ALTER COLUMN id SET DEFAULT nextval('public.stickers_id_seq'::regclass);


--
-- Name: user_sessions id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_sessions
    ALTER COLUMN id SET DEFAULT nextval('public.user_sessions_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: bots bots_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bots
    ADD CONSTRAINT bots_pkey PRIMARY KEY (id);


--
-- Name: contact_groups contact_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contact_groups
    ADD CONSTRAINT contact_groups_pkey PRIMARY KEY (id);


--
-- Name: contact_requests contact_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contact_requests
    ADD CONSTRAINT contact_requests_pkey PRIMARY KEY (id);


--
-- Name: contacts contacts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contacts
    ADD CONSTRAINT contacts_pkey PRIMARY KEY (id);


--
-- Name: dialogs dialogs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dialogs
    ADD CONSTRAINT dialogs_pkey PRIMARY KEY (id);


--
-- Name: group_chat_members group_chat_members_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat_members
    ADD CONSTRAINT group_chat_members_pkey PRIMARY KEY (id);


--
-- Name: group_chat_notice group_chat_notice_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat_notice
    ADD CONSTRAINT group_chat_notice_pkey PRIMARY KEY (id);


--
-- Name: group_chat_requests group_chat_requests_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chat_requests
    ADD CONSTRAINT group_chat_requests_pkey PRIMARY KEY (id);


--
-- Name: group_chats group_chats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.group_chats
    ADD CONSTRAINT group_chats_pkey PRIMARY KEY (id);


--
-- Name: message_delete message_delete_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message_delete
    ADD CONSTRAINT message_delete_pkey PRIMARY KEY (id);


--
-- Name: message_vote_answers message_vote_answers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message_vote_answers
    ADD CONSTRAINT message_vote_answers_pkey PRIMARY KEY (id);


--
-- Name: message_votes message_votes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.message_votes
    ADD CONSTRAINT message_votes_pkey PRIMARY KEY (id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (id);


--
-- Name: splits splits_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.splits
    ADD CONSTRAINT splits_pkey PRIMARY KEY (id);


--
-- Name: sticker_items sticker_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sticker_items
    ADD CONSTRAINT sticker_items_pkey PRIMARY KEY (id);


--
-- Name: sticker_user sticker_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sticker_user
    ADD CONSTRAINT sticker_user_pkey PRIMARY KEY (id);


--
-- Name: stickers stickers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stickers
    ADD CONSTRAINT stickers_pkey PRIMARY KEY (id);


--
-- Name: user_sessions user_sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_sessions
    ADD CONSTRAINT user_sessions_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: dialogs_dialog_type_user_id_receiver_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX dialogs_dialog_type_user_id_receiver_id_idx ON public.dialogs USING btree (dialog_type, user_id, receiver_id);


--
-- PostgreSQL database dump complete
--
