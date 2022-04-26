-- artists
INSERT INTO Artist (name, count_followers, count_listening)
VALUES ('Queen', 0, 0),
       ('Дайте танк (!)', 0, 0),
       ('Pink Floyd', 0, 0),
       ('Валерий Меладзе', 0, 0),
       ('Валентин Стрыкало', 0, 0),
       ('Twenty One Pilots', 0, 0),
       ('Химера', 0, 0),
       ('Дора', 0, 0),
       ('Shortparis', 0, 0),
       ('R.E.M', 0, 0);

-- -- album covers
INSERT
INTO AlbumCover (quote, is_dark)
VALUES ('Deluxe Edition 2011 Remaster', false),
       ('Нельзя в апокалипсис петь о том, как ты устал ходить на работу', true),
       ('This correlates with the inner world of a person and sets the mood for music describing the feelings he has experienced throughout his life. Amid the chaos, there is beauty and hope for humanity.',
        true),
       ('Вопреки — это несмотря ни на что… Несмотря на то, что обстоятельства вынуждают человека поступать одним образом, он поступает иначе, вопреки всем обстоятельствам.',
        true),
       ('Как и на первой пластинке, юмора будет предостаточно, разница лишь в том, что он будет более завуалированным. Ведь каждый любит вуаль.',
        false),
       ('Blurryface embodies all my fears and complexes and is an immaterial entity living inside, which makes me feel insecure in my creativity and in what I create',
        true),
       ('Я решил напечатать две тысячи плакатов и заклеить ими весь город, чтобы люди натыкались на них везде: в метро, во дворах и на заборах. Это не было рекламным трюком, хотя и могло сработать, как реклама. Если бы я в течение месяца натыкался бы на это слово, а потом, зайдя в магазин, увидел бы альбом с такой картинкой и таким названием, то я обязательно из любопытства его купил бы.',
        true),
       ('В нашей крови, в самых глубинах нашего существа сохраняется непреодолимым наследие древних времён; но я говорю не о первобытной жестокости и животных инстинктах, которые приписывают нам жидовский психоанализ и эволюционизм: это наследие древних, мифических времён, это наследие света.',
        true),
       ('Их музыка - революционна и стихийна, но в ней, в то же время, ощущается угрюмая, меланхолическая опасность; будучи по своей сути пост-дарквэйв группой, Shortparis обладает достаточной харизмой, чтобы собрать стадион подобно Depeche Mode',
        true),
       ('That me in the corner. That me in the spotlight', false);
--
--
-- -- albums
INSERT INTO Album (title, artist_id, count_likes, count_listening, date)
VALUES ('A Night At The Opera', 1, 0, 0, 1975),
       ('Человеко-часы', 2, 0, 0, 2020),
       ('The Dark Side of the Moon', 3, 0, 0, 1973),
       ('Вопреки', 4, 0, 0, 2008),
       ('Часть чего-то большего', 5, 0, 0, 2013),
       ('Blurryface', 6, 0, 0, 2015),
       ('ZUDWA-DWA', 7, 0, 0, 2003),
       ('Младшая сестра', 8, 0, 0, 2019),
       ('Так закалялась сталь', 9, 0, 0, 2019),
       ('Out Of Time', 9, 0, 0, 2019);

-- tracks
INSERT INTO Track (album_id, artist_id, title, duration, count_likes, count_listening)
VALUES (1, 1, 'Youre My Best Friend', 172, 0, 0),     -- 1
       (1, 1, 'Love Of My Life', 217, 0, 0),          -- 2
       (1, 1, 'Bohemian Rhapsody', 355, 0, 0),        -- 3
       (1, 1, 'Keep Yourself Alive', 244, 0, 0),      -- 4
       (2, 2, 'Профессионал', 219, 0, 0),             -- 5
       (2, 2, 'Люди', 162, 0, 0),                     -- 6
       (2, 2, 'Альтернатива', 230, 0, 0),             -- 7
       (2, 2, 'Ретро', 107, 0, 0),                    -- 8
       (3, 3, 'Speak To Me', 90, 0, 0),               -- 9
       (3, 3, 'The Great Gig in the Sky', 276, 0, 0), -- 10
       (3, 3, 'Money', 382, 0, 0),                    -- 11
       (3, 3, 'Us And Them', 423, 0, 0),              -- 12
       (4, 4, 'Параллельные', 227, 0, 0),             -- 13
       (4, 4, 'Ей никогда не быть моей', 255, 0, 0),  -- 14
       (4, 4, 'Иностранец', 246, 0, 0),               -- 15
       (4, 4, 'Вопреки', 265, 0, 0),                  -- 16
       (5, 5, 'Самый лучший друг', 165, 0, 0),        -- 17
       (5, 5, 'Все мои друзья', 160, 0, 0),           -- 18
       (5, 5, 'Космос нас ждёт', 202, 0, 0),          -- 19
       (5, 5, 'Кладбище самолётов', 352, 0, 0),       -- 20
       (5, 5, 'Улица Сталеваров', 258, 0, 0),         -- 21
       (6, 6, 'Heavydirtysoul', 234, 0, 0),           -- 22
       (6, 6, 'Stressed Out', 202, 0, 0),             -- 23
       (6, 6, 'Ride', 214, 0, 0),                     -- 24
       (7, 7, 'Пётр', 135, 0, 0),                     -- 25
       (7, 7, 'Ай-лю-ли', 187, 0, 0),                 -- 26
       (7, 7, 'Зайцы', 221, 0, 0),                    -- 27
       (7, 7, 'Вороны', 207, 0, 0),                   -- 28
       (8, 8, 'Дорадура', 134, 0, 0),                 -- 29
       (8, 8, 'Младшая сестра', 222, 0, 0),           -- 30
       (8, 8, 'Подружки', 142, 0, 0),                 -- 31
       (8, 8, 'Задолбал меня игнорить', 179, 0, 0),   -- 32
       (9, 9, 'Поломало', 275, 0, 0),                 -- 33
       (9, 9, 'Стыд', 268, 0, 0),                     -- 34
       (9, 9, 'Страшно', 289, 0, 0),                  -- 35
       (9, 9, 'Так закалялась сталь', 247, 0, 0),     -- 36
       (10, 10, 'Radio Song', 252, 0, 0),             -- 37
       (10, 10, 'Losing My Religion', 266, 0, 0),     -- 38
       (10, 10, 'Low', 295, 0, 0); -- 39


CREATE INDEX album_search_ru
    ON album
        USING gin (to_tsvector('russian', "title"));
CREATE INDEX album_search_en
    ON album
        USING gin (to_tsvector('english', "title"));
CREATE INDEX album_search_fr
    ON album
        USING gin (to_tsvector('french', "title"));

CREATE INDEX track_search_ru
    ON track
        USING gin (to_tsvector('russian', "title"));
CREATE INDEX track_search_en
    ON track
        USING gin (to_tsvector('english', "title"));
CREATE INDEX track_search_fr
    ON track
        USING gin (to_tsvector('french', "title"));


-- SELECT * from album;
--
-- SELECT * FROM Album WHERE to_tsvector(title)
--                               @@ plainto_tsquery('Часы');

EXPLAIN ANALYSE
SELECT *
FROM album
WHERE to_tsvector("title") @@ plainto_tsquery('rkpnysiugz')
ORDER BY ts_rank(to_tsvector("title"), plainto_tsquery('rkpnysiugz')) DESC;

-- EXPLAIN ANALYSE SELECT *
--                 FROM album
--                 WHERE title = 'rkpnysiugz'
--                 ORDER BY title DESC;
