--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

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
SET default_with_oids = false;
SET default_table_access_method = heap;
SET search_path = public, pg_catalog;

CREATE TABLE Users
(
    id              serial       NOT NULL,
    email           varchar(255) NOT NULL UNIQUE,
    username        varchar(128) NOT NULL UNIQUE,
    avatar          varchar(255),
    password_hash   varchar(64),
    count_following integer      NOT NULL,
    CONSTRAINT Users_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );



CREATE TABLE Tracks
(
    id              serial       NOT NULL,
    album_id        integer      NOT NULL,
    author_id       integer      NOT NULL,
    title           varchar(255) NOT NULL,
    duration        integer      NOT NULL,
    mp4             varchar(255) NOT NULL,
    count_likes     integer      NOT NULL,
    count_listening integer      NOT NULL,
    CONSTRAINT Tracks_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserTracksLikes
(
    user_id  integer NOT NULL,
    track_id integer NOT NULL
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserArtistsFollowing
(
    user_id   integer NOT NULL,
    artist_id integer NOT NULL
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserListenedTracks
(
    user_id        integer  NOT NULL,
    track_id       integer  NOT NULL,
    listening_date DATE NOT NULL DEFAULT 'now'
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserPlayer
(
    user_id     integer    NOT NULL UNIQUE,
    track_id    integer    NOT NULL,
    timestamp   integer    NOT NULL,
    track_from  varchar(1) NOT NULL,
    playlist_id integer,
    album_id    integer
) WITH (
      OIDS= FALSE
    );



CREATE TABLE Playlists
(
    id              serial       NOT NULL,
    title           varchar(255) NOT NULL,
    author_id       integer      NOT NULL,
    date_create     DATE     NOT NULL DEFAULT 'now',
    count_likes     integer      NOT NULL,
    count_added     integer      NOT NULL,
    count_listening integer      NOT NULL,
    CONSTRAINT Playlists_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );



CREATE TABLE PlaylistTracks
(
    playlist_id integer NOT NULL,
    track_id    integer NOT NULL
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserPlaylists
(
    user_id     integer NOT NULL,
    playlist_id integer NOT NULL
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserListenedPlaylists
(
    user_id        integer  NOT NULL,
    playlist_id    integer  NOT NULL,
    date_listening DATE NOT NULL
) WITH (
      OIDS= FALSE
    );



CREATE TABLE Albums
(
    id              serial       NOT NULL,
    title           varchar(255) NOT NULL,
    author_id       integer      NOT NULL,
    count_likes     integer      NOT NULL,
    count_listening integer      NOT NULL,
    date            DATE     NOT NULL,
    cover_id        integer      NOT NULL UNIQUE,
    CONSTRAINT Albums_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );



CREATE TABLE Artist
(
    id              serial       NOT NULL,
    name            varchar(255) NOT NULL,
    photo           varchar(255) NOT NULL,
    count_followers integer      NOT NULL,
    count_listening integer      NOT NULL,
    CONSTRAINT Artist_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );



CREATE TABLE AlbumCovers
(
    id      serial       NOT NULL,
    title   varchar(255) NOT NULL,
    photo   varchar(255) NOT NULL,
    quote   varchar(512) NOT NULL,
    is_dark BOOLEAN      NOT NULL,
    CONSTRAINT AlbumCovers_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserAlbumsLikes
(
    user_id  integer NOT NULL,
    album_id integer NOT NULL
) WITH (
      OIDS= FALSE
    );



ALTER TABLE Tracks
    ADD CONSTRAINT Tracks_fk0 FOREIGN KEY (album_id) REFERENCES Albums (id);
ALTER TABLE Tracks
    ADD CONSTRAINT Tracks_fk1 FOREIGN KEY (author_id) REFERENCES Artist (id);

ALTER TABLE UserTracksLikes
    ADD CONSTRAINT UserTracksLikes_fk0 FOREIGN KEY (user_id) REFERENCES Users (id);
ALTER TABLE UserTracksLikes
    ADD CONSTRAINT UserTracksLikes_fk1 FOREIGN KEY (track_id) REFERENCES Tracks (id);

ALTER TABLE UserArtistsFollowing
    ADD CONSTRAINT UserArtistsFollowing_fk0 FOREIGN KEY (user_id) REFERENCES Users (id);
ALTER TABLE UserArtistsFollowing
    ADD CONSTRAINT UserArtistsFollowing_fk1 FOREIGN KEY (artist_id) REFERENCES Artist (id);

ALTER TABLE UserListenedTracks
    ADD CONSTRAINT UserListenedTracks_fk0 FOREIGN KEY (user_id) REFERENCES Users (id);
ALTER TABLE UserListenedTracks
    ADD CONSTRAINT UserListenedTracks_fk1 FOREIGN KEY (track_id) REFERENCES Tracks (id);

ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk0 FOREIGN KEY (user_id) REFERENCES Users (id);
ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk1 FOREIGN KEY (track_id) REFERENCES Tracks (id);
ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk2 FOREIGN KEY (playlist_id) REFERENCES Tracks (id);
ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk3 FOREIGN KEY (album_id) REFERENCES Tracks (id);

ALTER TABLE Playlists
    ADD CONSTRAINT Playlists_fk0 FOREIGN KEY (author_id) REFERENCES Users (id);

ALTER TABLE PlaylistTracks
    ADD CONSTRAINT PlaylistTracks_fk0 FOREIGN KEY (playlist_id) REFERENCES Playlists (id);
ALTER TABLE PlaylistTracks
    ADD CONSTRAINT PlaylistTracks_fk1 FOREIGN KEY (track_id) REFERENCES Tracks (id);

ALTER TABLE UserPlaylists
    ADD CONSTRAINT UserPlaylists_fk0 FOREIGN KEY (user_id) REFERENCES Users (id);
ALTER TABLE UserPlaylists
    ADD CONSTRAINT UserPlaylists_fk1 FOREIGN KEY (playlist_id) REFERENCES Playlists (id);

ALTER TABLE UserListenedPlaylists
    ADD CONSTRAINT UserListenedPlaylists_fk0 FOREIGN KEY (user_id) REFERENCES Users (id);
ALTER TABLE UserListenedPlaylists
    ADD CONSTRAINT UserListenedPlaylists_fk1 FOREIGN KEY (playlist_id) REFERENCES Playlists (id);

ALTER TABLE Albums
    ADD CONSTRAINT Albums_fk0 FOREIGN KEY (author_id) REFERENCES Artist (id);
ALTER TABLE Albums
    ADD CONSTRAINT Albums_fk1 FOREIGN KEY (cover_id) REFERENCES AlbumCovers (id);



ALTER TABLE UserAlbumsLikes
    ADD CONSTRAINT UserAlbumsLikes_fk0 FOREIGN KEY (user_id) REFERENCES Users (id);
ALTER TABLE UserAlbumsLikes
    ADD CONSTRAINT UserAlbumsLikes_fk1 FOREIGN KEY (album_id) REFERENCES Albums (id);














