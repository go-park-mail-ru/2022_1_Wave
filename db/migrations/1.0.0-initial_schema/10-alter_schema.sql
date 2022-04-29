SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET
    check_function_bodies = false;
SET
    xmloption = content;
SET
    client_min_messages = warning;
SET
    row_security = off;
SET
    default_tablespace = '';
SET
    default_with_oids = false;
SET
    default_table_access_method = heap;
SET
    search_path = public, pg_catalog;



DROP TABLE if exists UserAlbumsLike;
DROP TABLE if exists UserFavoriteTracks;
DROP TABLE if exists UserFavoriteArtists;
DROP TABLE if exists UserFavoriteAlbums;
DROP TABLE if exists UserArtistsLike;
DROP TABLE if exists UserArtistsFollowing;
DROP TABLE if exists UserTracksLike;
DROP TABLE if exists UserListenedTrack;
DROP TABLE if exists UserPlaylist;
DROP TABLE if exists UserListenedPlaylist;
DROP TABLE if exists UserPlayer;
DROP TABLE if exists PlaylistTrack;
DROP TABLE if exists Playlist;
DROP TABLE if exists Users;
DROP TABLE if exists Track;
DROP TABLE if exists AlbumCover;
DROP TABLE if exists Album;
DROP TABLE if exists Single;
DROP TABLE if exists Artist;
DROP TABLE if exists place;



CREATE TABLE Users
(
    id              serial       NOT NULL,
    email           varchar(255) NOT NULL UNIQUE,
    username        varchar(128) NOT NULL UNIQUE,
    avatar          varchar(255),
    password_hash   varchar(64),
    count_following integer DEFAULT 0,
    CONSTRAINT Users_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );



CREATE TABLE Track
(
    id              serial       NOT NULL,
    album_id        integer,
    artist_id       integer      NOT NULL,
--     cover_id        serial       NOT NULL,
    title           varchar(255) NOT NULL,
    duration        integer      NOT NULL,
--     mp4_id          serial       NOT NULL,
    count_likes     integer      NOT NULL,
    count_listening integer      NOT NULL,
    CONSTRAINT Track_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );


CREATE TABLE UserTracksLike
(
    user_id  integer NOT NULL,
    track_id integer NOT NULL
) WITH (
      OIDS = FALSE
    );

CREATE TABLE UserFavoriteTracks
(
    user_id  integer NOT NULL,
    track_id integer NOT NULL
) WITH (
      OIDS = FALSE
    );

CREATE TABLE UserFavoriteAlbums
(
    user_id  integer NOT NULL,
    album_id integer NOT NULL
) WITH (
      OIDS = FALSE
    );
CREATE TABLE UserFavoriteArtists
(
    user_id   integer NOT NULL,
    artist_id integer NOT NULL
) WITH (
      OIDS = FALSE
    );

CREATE TABLE UserArtistsFollowing
(
    user_id   integer NOT NULL,
    artist_id integer NOT NULL
) WITH (
      OIDS= FALSE
    );



CREATE TABLE UserListenedTrack
(
    user_id        integer NOT NULL,
    track_id       integer NOT NULL,
    listening_date DATE    NOT NULL DEFAULT 'now'
) WITH (
      OIDS = FALSE
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



CREATE TABLE Playlist
(
    id        serial       NOT NULL,
    title     varchar(255) NOT NULL,
    CONSTRAINT Playlist_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );


CREATE TABLE PlaylistTrack
(
    playlist_id integer NOT NULL,
    track_id    integer NOT NULL
) WITH (
      OIDS = FALSE
    );

CREATE TABLE UserPlaylist
(
    user_id     integer NOT NULL,
    playlist_id integer NOT NULL
) WITH (
      OIDS = FALSE
    );



CREATE TABLE UserListenedPlaylist
(
    user_id        integer NOT NULL,
    playlist_id    integer NOT NULL,
    date_listening DATE    NOT NULL
) WITH (
      OIDS = FALSE
    );



CREATE TABLE Album
(
    id              serial       NOT NULL,
    title           varchar(255) NOT NULL,
    artist_id       integer      NOT NULL,
    count_likes     integer      NOT NULL,
    count_listening integer      NOT NULL,
    date            bigint       NOT NULL,
--     cover_id        serial       NOT NULL,
    CONSTRAINT Album_pk PRIMARY KEY (id)
) WITH (
      OIDS = FALSE
    );

-- CREATE TABLE Single
-- (
--     id         serial  NOT NULL,
--     artist_id  integer NOT NULL UNIQUE,
--     singles_id integer NOT NULL,
--     CONSTRAINT Single_pk PRIMARY KEY (id)
-- ) WITH (
--       OIDS= FALSE
--     );


CREATE TABLE Artist
(
    id              serial       NOT NULL,
    name            varchar(255) NOT NULL,
--     photo_id        serial      NOT NULL,
    count_likes     integer      NOT NULL,
    count_followers integer      NOT NULL,
    count_listening integer      NOT NULL,
    CONSTRAINT Artist_pk PRIMARY KEY (id)
) WITH (
      OIDS= FALSE
    );

CREATE TABLE UserArtistsLike
(
    user_id   integer NOT NULL,
    artist_id integer NOT NULL
) WITH (
      OIDS = FALSE
    );


CREATE TABLE AlbumCover
(
    id      serial       NOT NULL,
    quote   varchar(512) NOT NULL,
    is_dark BOOLEAN      NOT NULL,
    CONSTRAINT AlbumCover_pk PRIMARY KEY (id)
) WITH (
      OIDS = FALSE
    );


CREATE TABLE UserAlbumsLike
(
    user_id  integer NOT NULL,
    album_id integer NOT NULL
) WITH (
      OIDS = FALSE
    );



ALTER TABLE Track
    ADD CONSTRAINT Tracks_fk1 FOREIGN KEY (artist_id) REFERENCES Artist (id) ON DELETE CASCADE;

ALTER TABLE UserTracksLike
    ADD CONSTRAINT UserTracksLike_fk0 FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserTracksLike
    ADD CONSTRAINT UserTracksLike_fk1 FOREIGN KEY (track_id) REFERENCES Track (id) ON DELETE CASCADE;
ALTER TABLE UserTracksLike
    ADD CONSTRAINT uniq_user_track_like UNIQUE (user_id, track_id);

ALTER TABLE UserFavoriteTracks
    ADD CONSTRAINT real_user_for_favorite FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserFavoriteTracks
    ADD CONSTRAINT real_track_to_favorite FOREIGN KEY (track_id) REFERENCES Track (id) ON DELETE CASCADE;

ALTER TABLE UserFavoriteArtists
    ADD CONSTRAINT real_user_for_favorite FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserFavoriteArtists
    ADD CONSTRAINT real_track_to_favorite FOREIGN KEY (artist_id) REFERENCES Artist (id) ON DELETE CASCADE;

ALTER TABLE UserFavoriteAlbums
    ADD CONSTRAINT real_user_for_favorite FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserFavoriteAlbums
    ADD CONSTRAINT real_track_to_favorite FOREIGN KEY (album_id) REFERENCES Album (id) ON DELETE CASCADE;

ALTER TABLE UserArtistsFollowing
    ADD CONSTRAINT UserArtistsFollowing_fk0 FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserArtistsFollowing
    ADD CONSTRAINT UserArtistsFollowing_fk1 FOREIGN KEY (artist_id) REFERENCES Artist (id) ON DELETE CASCADE;

ALTER TABLE UserListenedTrack
    ADD CONSTRAINT UserListenedTrack_fk0 FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserListenedTrack
    ADD CONSTRAINT UserListenedTrack_fk1 FOREIGN KEY (track_id) REFERENCES Track (id) ON DELETE CASCADE;

ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk0 FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk1 FOREIGN KEY (track_id) REFERENCES Track (id);
ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk2 FOREIGN KEY (playlist_id) REFERENCES Track (id);
ALTER TABLE UserPlayer
    ADD CONSTRAINT UserPlayer_fk3 FOREIGN KEY (album_id) REFERENCES Track (id);

ALTER TABLE PlaylistTrack
    ADD CONSTRAINT PlaylistTrack_fk0 FOREIGN KEY (playlist_id) REFERENCES Playlist (id) ON DELETE CASCADE;
ALTER TABLE PlaylistTrack
    ADD CONSTRAINT PlaylistTrack_fk1 FOREIGN KEY (track_id) REFERENCES Track (id) ON DELETE CASCADE;
ALTER TABLE PlaylistTrack
    ADD CONSTRAINT uniq_playlist_track UNIQUE (playlist_id, track_id);


ALTER TABLE UserPlaylist
    ADD CONSTRAINT UserPlaylist_fk0 FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserPlaylist
    ADD CONSTRAINT UserPlaylist_fk1 FOREIGN KEY (playlist_id) REFERENCES Playlist (id) ON DELETE CASCADE;

ALTER TABLE UserListenedPlaylist
    ADD CONSTRAINT UserListenedPlaylist_fk0 FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserListenedPlaylist
    ADD CONSTRAINT UserListenedPlaylist_fk1 FOREIGN KEY (playlist_id) REFERENCES Playlist (id) ON DELETE CASCADE;

ALTER TABLE Album
    ADD CONSTRAINT Album_fk0 FOREIGN KEY (artist_id) REFERENCES Artist (id) ON DELETE CASCADE;
ALTER TABLE AlbumCover
    ADD CONSTRAINT AlbumCover_fk0 FOREIGN KEY (id) REFERENCES Album (id) ON DELETE CASCADE;

ALTER TABLE UserAlbumsLike
    ADD CONSTRAINT UserAlbumsLike_fk0 FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserAlbumsLike
    ADD CONSTRAINT UserAlbumsLike_fk1 FOREIGN KEY (album_id) REFERENCES Album (id) ON DELETE CASCADE;
ALTER TABLE UserAlbumsLike
    ADD CONSTRAINT uniq_user_album_like UNIQUE (user_id, album_id);

ALTER TABLE UserArtistsLike
    ADD CONSTRAINT real_user_for_like_track FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE;
ALTER TABLE UserArtistsLike
    ADD CONSTRAINT real_artist_to_be_liked FOREIGN KEY (artist_id) REFERENCES Album (id) ON DELETE CASCADE;
ALTER TABLE UserArtistsLike
    ADD CONSTRAINT uniq_user_artist_like UNIQUE (user_id, artist_id);
