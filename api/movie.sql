CREATE TABLE movies
(
    id               INTEGER PRIMARY KEY NOT NULL,
    imdb_id          VARCHAR(20)         NOT NULL,
    title            VARCHAR(255)        NOT NULL,
    director         VARCHAR(255)        NOT NULL,
    year             INT                 NOT NULL,
    rating           VARCHAR(8)          NOT NULL,
    genres           VARCHAR(255)        NOT NULL,
    runtime          INT                 NOT NULL,
    country          VARCHAR(255)        NOT NULL,
    language         VARCHAR(255)        NOT NULL,
    imdb_score       NUMERIC             NOT NULL,
    imdb_votes       INT                 NOT NULL,
    metacritic_score NUMERIC             NOT NULL
)