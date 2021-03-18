-- Type: typeenum

-- DROP TYPE public.typeenum;

CREATE TYPE public.typeenum AS ENUM
    ('Yes', 'No');

ALTER TYPE public.typeenum
    OWNER TO postgres;



-- Table: public.apps

-- DROP TABLE public.apps;

CREATE TABLE public.apps
(
    app_code uuid NOT NULL DEFAULT uuid_generate_v4(),
    app_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    app_logo character varying(70) COLLATE pg_catalog."default" NOT NULL DEFAULT 'default.png'::character varying,
    create_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp with time zone,
    delete_at timestamp with time zone,
    app_publish typeenum,
    CONSTRAINT apps_pkey PRIMARY KEY (app_code)
)

TABLESPACE pg_default;


ALTER TABLE public.apps
    OWNER to postgres;
-- Index: apps_app_code_idx

-- DROP INDEX public.apps_app_code_idx;

CREATE INDEX apps_app_code_idx
    ON public.apps USING btree
    (app_code ASC NULLS LAST)
    TABLESPACE pg_default;



-- Table: public.category

-- DROP TABLE public.category;

CREATE TABLE public.category
(
    category_code uuid NOT NULL DEFAULT uuid_generate_v4(),
    category_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    category_desc text COLLATE pg_catalog."default",
    category_icon text COLLATE pg_catalog."default" NOT NULL,
    app_code uuid NOT NULL,
    create_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp with time zone,
    delete_at timestamp with time zone,
    app_publish typeenum,
    CONSTRAINT category_pkey PRIMARY KEY (category_code),
    CONSTRAINT category_category_code_fkey FOREIGN KEY (category_code)
        REFERENCES public.apps (app_code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.category
    OWNER to postgres;
-- Index: category_category_code_app_code_idx

-- DROP INDEX public.category_category_code_app_code_idx;

CREATE INDEX category_category_code_app_code_idx
    ON public.category USING btree
    (category_code ASC NULLS LAST, app_code ASC NULLS LAST)
    TABLESPACE pg_default;



-- Table: public.content

-- DROP TABLE public.content;

CREATE TABLE public.content
(
    content_code uuid NOT NULL DEFAULT uuid_generate_v4(),
    content_title character varying(255) COLLATE pg_catalog."default" NOT NULL,
    content_desc text COLLATE pg_catalog."default" NOT NULL,
    category_code uuid NOT NULL,
    manual_code uuid NOT NULL,
    create_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp with time zone,
    delete_at timestamp with time zone,
    app_publish typeenum,
    CONSTRAINT content_pkey PRIMARY KEY (content_code),
    CONSTRAINT content_category_code_fkey FOREIGN KEY (category_code)
        REFERENCES public.category (category_code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT content_manual_code_fkey FOREIGN KEY (manual_code)
        REFERENCES public.manual (manual_code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.content
    OWNER to postgres;
-- Index: content_category_code_manual_code_content_code_idx

-- DROP INDEX public.content_category_code_manual_code_content_code_idx;

CREATE INDEX content_category_code_manual_code_content_code_idx
    ON public.content USING btree
    (category_code ASC NULLS LAST, manual_code ASC NULLS LAST, content_code ASC NULLS LAST)
    TABLESPACE pg_default;


-- Table: public.manual

-- DROP TABLE public.manual;

CREATE TABLE public.manual
(
    manual_code uuid NOT NULL DEFAULT uuid_generate_v4(),
    manual_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    manual_desc text COLLATE pg_catalog."default" NOT NULL,
    manual_icon text COLLATE pg_catalog."default" NOT NULL,
    app_code uuid NOT NULL,
    create_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp with time zone,
    delete_at timestamp with time zone,
    app_publish typeenum,
    CONSTRAINT manual_pkey PRIMARY KEY (manual_code),
    CONSTRAINT manual_app_code_fkey FOREIGN KEY (app_code)
        REFERENCES public.apps (app_code) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE public.manual
    OWNER to postgres;
-- Index: manual_manual_code_app_code_idx

-- DROP INDEX public.manual_manual_code_app_code_idx;

CREATE INDEX manual_manual_code_app_code_idx
    ON public.manual USING btree
    (manual_code ASC NULLS LAST, app_code ASC NULLS LAST)
    TABLESPACE pg_default;