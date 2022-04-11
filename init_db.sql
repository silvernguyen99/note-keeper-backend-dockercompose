-- Table: public.tb_users

-- DROP TABLE IF EXISTS public.tb_users;

CREATE SEQUENCE tb_users_user_id_seq;

CREATE TABLE IF NOT EXISTS public.tb_users
(
    user_id integer DEFAULT nextval('tb_users_user_id_seq'::regclass) PRIMARY KEY,
    user_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password_hash character varying(255) COLLATE pg_catalog."default",
    access_token character varying COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT tb_users__access_token_unique UNIQUE (access_token)
);

ALTER SEQUENCE tb_users_user_id_seq
OWNED BY tb_users.user_id;

ALTER TABLE IF EXISTS public.tb_users
    OWNER to postgres;

CREATE UNIQUE INDEX tb_users__idx_access_token
ON public.tb_users(access_token);

CREATE UNIQUE INDEX tb_users__idx_user_id
ON public.tb_users(user_id);

-- Table: public.tb_login_social

-- DROP TABLE IF EXISTS public.tb_login_social;
CREATE TABLE IF NOT EXISTS public.tb_login_social
(
    user_id integer NOT NULL,
    type_provider character varying(255) COLLATE pg_catalog."default" NOT NULL,
    social_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE IF EXISTS public.tb_login_social
    OWNER to postgres;

CREATE UNIQUE INDEX tb_login_social__user_provider_unique ON tb_login_social (user_id,type_provider);
CREATE UNIQUE INDEX tb_login_social__provider_social_unique ON tb_login_social (social_id,type_provider);


-- Table: public.tb_notes

-- DROP TABLE IF EXISTS public.tb_notes;
CREATE SEQUENCE tb_notes_note_id_seq;

CREATE TABLE IF NOT EXISTS public.tb_notes
(
    user_id integer NOT NULL,
    note_id integer DEFAULT nextval('tb_notes_note_id_seq'::regclass) PRIMARY KEY,
    note_name character varying(255) COLLATE pg_catalog."default" NOT NULL,
    content text COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER SEQUENCE tb_notes_note_id_seq
OWNED BY tb_notes.note_id;

CREATE INDEX tb_notes__idx_user_id ON tb_notes (user_id);
CREATE INDEX tb_notes__idx_user_id_note_id ON tb_notes (user_id,note_id);
CREATE UNIQUE INDEX tb_notes__idx_user_id_note_name ON tb_notes (user_id,note_name);



