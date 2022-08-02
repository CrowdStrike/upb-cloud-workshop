SET client_encoding = 'UTF8';

DROP ROLE IF EXISTS upb;
CREATE USER upb WITH PASSWORD 'upb';

BEGIN;

CREATE TABLE IF NOT EXISTS actors(
id char(36) primary key,
name varchar(250) not null,
quote text
);

CREATE TABLE IF NOT EXISTS movies(
id char(36) primary key,
name varchar(250) unique not null,
description text,
actors char(250)[]
);

COPY actors (id, name, quote) from stdin (DELIMITER '|');
001|Keanu Reeves|Sometimes life imitates art
002|Lawrence Fishburne|People think that I'm haughty and stuck up, but really I'm just very shy
003|Leonardo DiCaprio|Only you and you alone can change your situation
\.


COPY movies (id, name, description, actors) from stdin (DELIMITER '|');
001|Matrix|Best movie|{001,002}
002|Inception|Another best movie|{003}
\.

GRANT ALL privileges ON actors TO upb;
GRANT ALL privileges ON movies TO upb;

-- TODO: task 6: create an index on the actors table to improve the search speed for actor names

COMMIT;
