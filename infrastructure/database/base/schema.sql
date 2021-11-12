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

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cashes; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.cashes
(
    code        uuid                           DEFAULT public.uuid_generate_v4() NOT NULL,
    user_dni    character(8)                                                     NOT NULL,
    amount      numeric(12, 2)                 DEFAULT 0.00                      NOT NULL,
    state       boolean                        DEFAULT true                      NOT NULL,
    description character varying(300),
    created_at  timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6),
    updated_at  timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6)
);


--
-- Name: categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.categories
(
    code        uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name        character varying(120)                 NOT NULL,
    description character varying(300)
);


--
-- Name: contacts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.contacts
(
    code        uuid                           DEFAULT public.uuid_generate_v4() NOT NULL,
    name        character varying(255)                                           NOT NULL,
    type        character(1)                                                     NOT NULL,
    email       character varying(50),
    phone       character varying(30),
    address     character varying(120),
    description character varying(300),
    created_at  timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6),
    updated_at  timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6)
);


--
-- Name: expenses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.expenses
(
    code        uuid                           DEFAULT public.uuid_generate_v4() NOT NULL,
    description character varying(350)                                           NOT NULL,
    user_dni    character(8)                                                     NOT NULL,
    cash_code   uuid                                                             NOT NULL,
    total       numeric(12, 2)                 DEFAULT 0.00                      NOT NULL,
    created_at  timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6),
    updated_at  timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6)
);


--
-- Name: materials; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.materials
(
    code          uuid                           DEFAULT public.uuid_generate_v4() NOT NULL,
    name          character varying(150)                                           NOT NULL,
    description   character varying(300),
    category_code uuid                                                             NOT NULL,
    created_at    timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6)
);


--
-- Name: migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.migrations
(
    version character varying(255) NOT NULL
);


--
-- Name: purchase_details; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.purchase_details
(
    purchase_code uuid                        NOT NULL,
    material_code uuid                        NOT NULL,
    quantity      integer        DEFAULT 0    NOT NULL,
    price         numeric(12, 2) DEFAULT 0.00 NOT NULL
);


--
-- Name: purchases; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.purchases
(
    expense_code uuid NOT NULL,
    contact_code uuid NOT NULL,
    voucher_num  character varying(20)
);


--
-- Name: roles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.roles
(
    code        uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name        character varying(20)                  NOT NULL,
    description character varying(255)
);


--
-- Name: stocks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.stocks
(
    material_code uuid                                     NOT NULL,
    quantity      integer                        DEFAULT 0 NOT NULL,
    state         character(1),
    updated_at    timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6)
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users
(
    dni        character(8)          NOT NULL,
    name       character varying(20) NOT NULL,
    lastname   character varying(50) NOT NULL,
    rol_code   uuid                  NOT NULL,
    gender     character(1),
    image      character varying(100),
    address    character varying(100),
    phone      character varying(40),
    email      character varying(50),
    password   character varying(100),
    created_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP(6)
);


--
-- Name: cashes cashes_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cashes
    ADD CONSTRAINT cashes_pkey PRIMARY KEY (code);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (code);


--
-- Name: contacts contacts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.contacts
    ADD CONSTRAINT contacts_pkey PRIMARY KEY (code);


--
-- Name: expenses expenses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.expenses
    ADD CONSTRAINT expenses_pkey PRIMARY KEY (code);


--
-- Name: materials materials_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.materials
    ADD CONSTRAINT materials_pkey PRIMARY KEY (code);


--
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (version);


--
-- Name: purchase_details purchase_details_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.purchase_details
    ADD CONSTRAINT purchase_details_pkey PRIMARY KEY (purchase_code, material_code);


--
-- Name: purchases purchases_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.purchases
    ADD CONSTRAINT purchases_pkey PRIMARY KEY (expense_code);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (code);


--
-- Name: stocks stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.stocks
    ADD CONSTRAINT stocks_pkey PRIMARY KEY (material_code);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (dni);


--
-- Name: cashes cashes_user_dni_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cashes
    ADD CONSTRAINT cashes_user_dni_fkey FOREIGN KEY (user_dni) REFERENCES public.users (dni);


--
-- Name: expenses expenses_cash_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.expenses
    ADD CONSTRAINT expenses_cash_code_fkey FOREIGN KEY (cash_code) REFERENCES public.cashes (code);


--
-- Name: expenses expenses_user_dni_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.expenses
    ADD CONSTRAINT expenses_user_dni_fkey FOREIGN KEY (user_dni) REFERENCES public.users (dni);


--
-- Name: materials materials_category_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.materials
    ADD CONSTRAINT materials_category_code_fkey FOREIGN KEY (category_code) REFERENCES public.categories (code);


--
-- Name: purchase_details purchase_details_material_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.purchase_details
    ADD CONSTRAINT purchase_details_material_code_fkey FOREIGN KEY (material_code) REFERENCES public.materials (code);


--
-- Name: purchase_details purchase_details_purchase_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.purchase_details
    ADD CONSTRAINT purchase_details_purchase_code_fkey FOREIGN KEY (purchase_code) REFERENCES public.purchases (expense_code);


--
-- Name: purchases purchases_contact_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.purchases
    ADD CONSTRAINT purchases_contact_code_fkey FOREIGN KEY (contact_code) REFERENCES public.contacts (code);


--
-- Name: purchases purchases_expense_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.purchases
    ADD CONSTRAINT purchases_expense_code_fkey FOREIGN KEY (expense_code) REFERENCES public.expenses (code);


--
-- Name: stocks stocks_material_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.stocks
    ADD CONSTRAINT stocks_material_code_fkey FOREIGN KEY (material_code) REFERENCES public.materials (code);


--
-- Name: users users_rol_code_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_rol_code_fkey FOREIGN KEY (rol_code) REFERENCES public.roles (code);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.migrations (version)
VALUES ('20211031000216'),
       ('20211031000223'),
       ('20211103172711'),
       ('20211103182357'),
       ('20211103183704'),
       ('20211103235446'),
       ('20211104000825'),
       ('20211104002126'),
       ('20211104002223'),
       ('20211104210153');
