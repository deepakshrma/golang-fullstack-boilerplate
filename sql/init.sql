
--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
                              id integer NOT NULL,
                              first_name character varying(255),
                              last_name character varying(255),
                              email character varying(255),
                              password character varying(60),
                              is_admin integer,
                              created_at timestamp without time zone,
                              updated_at timestamp without time zone
);


ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
    );



insert into users (first_name, last_name, email, password, is_admin, created_at, updated_at)
values ('Admin', 'User', 'admin@example.com', '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK', 1, now(), now())
