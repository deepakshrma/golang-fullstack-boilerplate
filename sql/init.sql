CREATE TABLE public.users
(
    id         integer NOT NULL unique,
    first_name character varying(255),
    last_name  character varying(255),
    email      character varying(255) unique,
    password   character varying(60),
    is_admin   integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.users
    ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
        SEQUENCE NAME public.users_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );


CREATE TABLE public.todos
(
    id          integer NOT NULL,
    user_id     integer,
    title       character varying(255),
    description character varying(255),
    done        integer                     default 0,
    created_at  timestamp without time zone default now(),
    updated_at  timestamp without time zone default now()
);


ALTER TABLE public.todos
    ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
        SEQUENCE NAME public.todos_id_seq
        START WITH 1
        INCREMENT BY 1
        NO MINVALUE
        NO MAXVALUE
        CACHE 1
        );



ALTER TABLE ONLY public.todos
    ADD CONSTRAINT user_todos_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.todos
    ADD CONSTRAINT user_images_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE CASCADE;


INSERT INTO public.users (first_name, last_name, email, password, is_admin, created_at, updated_at)
values ('Admin', 'User', 'admin@example.com', '$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK', 1, now(),
        now());

INSERT INTO public.todos (user_id, title, description)
values (1, 'Test Todo', 'This is test todo');