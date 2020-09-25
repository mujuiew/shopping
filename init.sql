-- Table: public.account

-- DROP TABLE public.account;

CREATE TABLE public.account
(
    "ID" character varying(20) COLLATE pg_catalog."default" NOT NULL,
    "Facebook_ID" character varying(30) COLLATE pg_catalog."default",
    account_fname character varying(20) COLLATE pg_catalog."default" NOT NULL,
    account_lname character varying(20) COLLATE pg_catalog."default" NOT NULL,
    account_email character varying(20) COLLATE pg_catalog."default",
    account_phone character varying(10) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "Account_pkey" PRIMARY KEY ("ID")
)

TABLESPACE pg_default;

ALTER TABLE public.account
    OWNER to postgres;