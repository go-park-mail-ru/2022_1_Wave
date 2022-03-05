--
-- PostgreSQL database dump
--

-- Dumped from database version 13.3 (Debian 13.3-1.pgdg100+1)
-- Dumped by pg_dump version 13.3 (Debian 13.3-1.pgdg100+1)

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
-- Name: albumcovers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.albumcovers (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    photo character varying(255) NOT NULL,
    quote character varying(512) NOT NULL,
    is_dark boolean NOT NULL
);


ALTER TABLE public.albumcovers OWNER TO postgres;

--
-- Name: albumcovers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.albumcovers_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.albumcovers_id_seq OWNER TO postgres;

--
-- Name: albumcovers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.albumcovers_id_seq OWNED BY public.albumcovers.id;


--
-- Name: albums; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.albums (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    author_id integer NOT NULL,
    count_likes integer NOT NULL,
    count_listening integer NOT NULL,
    date date NOT NULL,
    cover_id integer NOT NULL
);


ALTER TABLE public.albums OWNER TO postgres;

--
-- Name: albums_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.albums_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.albums_id_seq OWNER TO postgres;

--
-- Name: albums_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.albums_id_seq OWNED BY public.albums.id;


--
-- Name: artist; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.artist (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    photo character varying(255) NOT NULL,
    count_followers integer NOT NULL,
    count_listening integer NOT NULL
);


ALTER TABLE public.artist OWNER TO postgres;

--
-- Name: artist_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.artist_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.artist_id_seq OWNER TO postgres;

--
-- Name: artist_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.artist_id_seq OWNED BY public.artist.id;


--
-- Name: playlists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.playlists (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    author_id integer NOT NULL,
    date_create date DEFAULT '2022-03-03'::date NOT NULL,
    count_likes integer NOT NULL,
    count_added integer NOT NULL,
    count_listening integer NOT NULL
);


ALTER TABLE public.playlists OWNER TO postgres;

--
-- Name: playlists_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.playlists_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.playlists_id_seq OWNER TO postgres;

--
-- Name: playlists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.playlists_id_seq OWNED BY public.playlists.id;


--
-- Name: playlisttracks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.playlisttracks (
    playlist_id integer NOT NULL,
    track_id integer NOT NULL
);


ALTER TABLE public.playlisttracks OWNER TO postgres;

--
-- Name: tracks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tracks (
    id integer NOT NULL,
    album_id integer NOT NULL,
    author_id integer NOT NULL,
    title character varying(255) NOT NULL,
    duration integer NOT NULL,
    mp4 character varying(255) NOT NULL,
    count_likes integer NOT NULL,
    count_listening integer NOT NULL
);


ALTER TABLE public.tracks OWNER TO postgres;

--
-- Name: tracks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tracks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tracks_id_seq OWNER TO postgres;

--
-- Name: tracks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tracks_id_seq OWNED BY public.tracks.id;


--
-- Name: useralbumslikes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.useralbumslikes (
    user_id integer NOT NULL,
    album_id integer NOT NULL
);


ALTER TABLE public.useralbumslikes OWNER TO postgres;

--
-- Name: userartistsfollowing; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userartistsfollowing (
    user_id integer NOT NULL,
    artist_id integer NOT NULL
);


ALTER TABLE public.userartistsfollowing OWNER TO postgres;

--
-- Name: userlistenedplaylists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userlistenedplaylists (
    user_id integer NOT NULL,
    playlist_id integer NOT NULL,
    date_listening date NOT NULL
);


ALTER TABLE public.userlistenedplaylists OWNER TO postgres;

--
-- Name: userlistenedtracks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userlistenedtracks (
    user_id integer NOT NULL,
    track_id integer NOT NULL,
    listening_date date DEFAULT '2022-03-03'::date NOT NULL
);


ALTER TABLE public.userlistenedtracks OWNER TO postgres;

--
-- Name: userplayer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userplayer (
    user_id integer NOT NULL,
    track_id integer NOT NULL,
    "timestamp" integer NOT NULL,
    track_from character varying(1) NOT NULL,
    playlist_id integer,
    album_id integer
);


ALTER TABLE public.userplayer OWNER TO postgres;

--
-- Name: userplaylists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userplaylists (
    user_id integer NOT NULL,
    playlist_id integer NOT NULL
);


ALTER TABLE public.userplaylists OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    username character varying(128) NOT NULL,
    avatar character varying(255),
    password_hash character varying(64),
    count_following integer NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

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


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: usertrackslikes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usertrackslikes (
    user_id integer NOT NULL,
    track_id integer NOT NULL
);


ALTER TABLE public.usertrackslikes OWNER TO postgres;

--
-- Name: albumcovers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albumcovers ALTER COLUMN id SET DEFAULT nextval('public.albumcovers_id_seq'::regclass);


--
-- Name: albums id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albums ALTER COLUMN id SET DEFAULT nextval('public.albums_id_seq'::regclass);


--
-- Name: artist id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.artist ALTER COLUMN id SET DEFAULT nextval('public.artist_id_seq'::regclass);


--
-- Name: playlists id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlists ALTER COLUMN id SET DEFAULT nextval('public.playlists_id_seq'::regclass);


--
-- Name: tracks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks ALTER COLUMN id SET DEFAULT nextval('public.tracks_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: albumcovers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.albumcovers (id, title, photo, quote, is_dark) FROM stdin;
\.


--
-- Data for Name: albums; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.albums (id, title, author_id, count_likes, count_listening, date, cover_id) FROM stdin;
\.


--
-- Data for Name: artist; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.artist (id, name, photo, count_followers, count_listening) FROM stdin;
\.


--
-- Data for Name: playlists; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.playlists (id, title, author_id, date_create, count_likes, count_added, count_listening) FROM stdin;
\.


--
-- Data for Name: playlisttracks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.playlisttracks (playlist_id, track_id) FROM stdin;
\.


--
-- Data for Name: tracks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tracks (id, album_id, author_id, title, duration, mp4, count_likes, count_listening) FROM stdin;
\.


--
-- Data for Name: useralbumslikes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.useralbumslikes (user_id, album_id) FROM stdin;
\.


--
-- Data for Name: userartistsfollowing; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userartistsfollowing (user_id, artist_id) FROM stdin;
\.


--
-- Data for Name: userlistenedplaylists; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userlistenedplaylists (user_id, playlist_id, date_listening) FROM stdin;
\.


--
-- Data for Name: userlistenedtracks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userlistenedtracks (user_id, track_id, listening_date) FROM stdin;
\.


--
-- Data for Name: userplayer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userplayer (user_id, track_id, "timestamp", track_from, playlist_id, album_id) FROM stdin;
\.


--
-- Data for Name: userplaylists; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userplaylists (user_id, playlist_id) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, username, avatar, password_hash, count_following) FROM stdin;
\.


--
-- Data for Name: usertrackslikes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.usertrackslikes (user_id, track_id) FROM stdin;
\.


--
-- Name: albumcovers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.albumcovers_id_seq', 1, false);


--
-- Name: albums_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.albums_id_seq', 1, false);


--
-- Name: artist_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.artist_id_seq', 1, false);


--
-- Name: playlists_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.playlists_id_seq', 1, false);


--
-- Name: tracks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tracks_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: albumcovers albumcovers_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albumcovers
    ADD CONSTRAINT albumcovers_pk PRIMARY KEY (id);


--
-- Name: albums albums_cover_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_cover_id_key UNIQUE (cover_id);


--
-- Name: albums albums_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_pk PRIMARY KEY (id);


--
-- Name: artist artist_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.artist
    ADD CONSTRAINT artist_pk PRIMARY KEY (id);


--
-- Name: playlists playlists_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlists
    ADD CONSTRAINT playlists_pk PRIMARY KEY (id);


--
-- Name: tracks tracks_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks
    ADD CONSTRAINT tracks_pk PRIMARY KEY (id);


--
-- Name: userplayer userplayer_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userplayer
    ADD CONSTRAINT userplayer_user_id_key UNIQUE (user_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: albums albums_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_fk0 FOREIGN KEY (author_id) REFERENCES public.artist(id);


--
-- Name: albums albums_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.albums
    ADD CONSTRAINT albums_fk1 FOREIGN KEY (cover_id) REFERENCES public.albumcovers(id);


--
-- Name: playlists playlists_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlists
    ADD CONSTRAINT playlists_fk0 FOREIGN KEY (author_id) REFERENCES public.users(id);


--
-- Name: playlisttracks playlisttracks_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlisttracks
    ADD CONSTRAINT playlisttracks_fk0 FOREIGN KEY (playlist_id) REFERENCES public.playlists(id);


--
-- Name: playlisttracks playlisttracks_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.playlisttracks
    ADD CONSTRAINT playlisttracks_fk1 FOREIGN KEY (track_id) REFERENCES public.tracks(id);


--
-- Name: tracks tracks_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks
    ADD CONSTRAINT tracks_fk0 FOREIGN KEY (album_id) REFERENCES public.albums(id);


--
-- Name: tracks tracks_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tracks
    ADD CONSTRAINT tracks_fk1 FOREIGN KEY (author_id) REFERENCES public.artist(id);


--
-- Name: useralbumslikes useralbumslikes_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.useralbumslikes
    ADD CONSTRAINT useralbumslikes_fk0 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: useralbumslikes useralbumslikes_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.useralbumslikes
    ADD CONSTRAINT useralbumslikes_fk1 FOREIGN KEY (album_id) REFERENCES public.albums(id);


--
-- Name: userartistsfollowing userartistsfollowing_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userartistsfollowing
    ADD CONSTRAINT userartistsfollowing_fk0 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: userartistsfollowing userartistsfollowing_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userartistsfollowing
    ADD CONSTRAINT userartistsfollowing_fk1 FOREIGN KEY (artist_id) REFERENCES public.artist(id);


--
-- Name: userlistenedplaylists userlistenedplaylists_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userlistenedplaylists
    ADD CONSTRAINT userlistenedplaylists_fk0 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: userlistenedplaylists userlistenedplaylists_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userlistenedplaylists
    ADD CONSTRAINT userlistenedplaylists_fk1 FOREIGN KEY (playlist_id) REFERENCES public.playlists(id);


--
-- Name: userlistenedtracks userlistenedtracks_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userlistenedtracks
    ADD CONSTRAINT userlistenedtracks_fk0 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: userlistenedtracks userlistenedtracks_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userlistenedtracks
    ADD CONSTRAINT userlistenedtracks_fk1 FOREIGN KEY (track_id) REFERENCES public.tracks(id);


--
-- Name: userplayer userplayer_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userplayer
    ADD CONSTRAINT userplayer_fk0 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: userplayer userplayer_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userplayer
    ADD CONSTRAINT userplayer_fk1 FOREIGN KEY (track_id) REFERENCES public.tracks(id);


--
-- Name: userplayer userplayer_fk2; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userplayer
    ADD CONSTRAINT userplayer_fk2 FOREIGN KEY (playlist_id) REFERENCES public.tracks(id);


--
-- Name: userplayer userplayer_fk3; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userplayer
    ADD CONSTRAINT userplayer_fk3 FOREIGN KEY (album_id) REFERENCES public.tracks(id);


--
-- Name: userplaylists userplaylists_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userplaylists
    ADD CONSTRAINT userplaylists_fk0 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: userplaylists userplaylists_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userplaylists
    ADD CONSTRAINT userplaylists_fk1 FOREIGN KEY (playlist_id) REFERENCES public.playlists(id);


--
-- Name: usertrackslikes usertrackslikes_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usertrackslikes
    ADD CONSTRAINT usertrackslikes_fk0 FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: usertrackslikes usertrackslikes_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usertrackslikes
    ADD CONSTRAINT usertrackslikes_fk1 FOREIGN KEY (track_id) REFERENCES public.tracks(id);


--
-- PostgreSQL database dump complete
--

