CREATE TABLE public.genres
(
    id         SERIAL PRIMARY KEY,
    genre      VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

INSERT INTO public.genres (genre, created_at, updated_at)
VALUES ('Comedy', now(), now()),
       ('Sci-Fi', now(), now()),
       ('Horror', now(), now()),
       ('Romance', now(), now()),
       ('Action', now(), now()),
       ('Thriller', now(), now()),
       ('Drama', now(), now()),
       ('Mystery', now(), now()),
       ('Crime', now(), now()),
       ('Animation', now(), now()),
       ('Adventure', now(), now()),
       ('Fantasy', now(), now()),
       ('Superhero', now(), now())
;

CREATE TABLE public.movies
(
    id           SERIAL PRIMARY KEY,
    title        VARCHAR(512),
    release_date DATE,
    runtime      INTEGER,
    mpaa_rating  VARCHAR(10),
    description  TEXT,
    image        VARCHAR(255),
    created_at   TIMESTAMP WITHOUT TIME ZONE,
    updated_at   TIMESTAMP WITHOUT TIME ZONE
);

INSERT INTO public.movies (title, release_date, runtime, mpaa_rating, description, image, created_at, updated_at)
VALUES ('Highlander', '1986-03-07', 116, 'R',
        'He fought his first battle on the Scottish Highlands in 1536. He will fight his greatest battle on the streets of New York City in 1986. His name is Connor MacLeod. He is immortal.',
        '/8Z8dptJEypuLoOQro1WugD855YE.jpg', now(), now()),
       ('Raiders of the Lost Ark', '1981-06-12', 115, 'PG-13',
        'Archaeology professor Indiana Jones ventures to seize a biblical artefact known as the Ark of the Covenant.While doing so, he puts up a fight against Renee and a troop of Nazis.',
        '/ceG9VzoRAVGwivFU403Wc3AHRys.jpg', now(), now()),
       ('The Godfather', '1972-03-24', 175, '18 A',
        'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.',
        '/3 bhkrj58Vtu7enYsRolD1fZdja1.jpg', now(), now())
;

CREATE TABLE public.movie_genres
(
    id       SERIAL PRIMARY KEY,
    movie_id INTEGER,
    genre_id INTEGER
);

ALTER TABLE public.movie_genres
    ADD CONSTRAINT movie_genres_movie_id_fkey FOREIGN KEY (movie_id) REFERENCES public.movies (id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE public.movie_genres
    ADD CONSTRAINT movie_genres_genre_id_fkey FOREIGN KEY (genre_id) REFERENCES public.genres (id) ON UPDATE CASCADE ON DELETE CASCADE;

INSERT INTO public.movie_genres (movie_id, genre_id)
VALUES (1, 5),
       (1, 12),
       (2, 5),
       (2, 11),
       (3, 9),
       (3, 7)
;

CREATE TABLE public.users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name  VARCHAR(255),
    email      VARCHAR(255),
    password   VARCHAR(255),
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE
);

INSERT INTO public.users (first_name, last_name, email, password, created_at, updated_at)
VALUES ('Admin', 'User', 'admin@example.com', '$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy',
        now(),
        now())
;