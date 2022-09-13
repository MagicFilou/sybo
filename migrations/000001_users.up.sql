CREATE TABLE IF NOT EXISTS
users (
id TEXT PRIMARY KEY,
name TEXT NOT NULL,
games_played INTEGER,
score INTEGER,
friends TEXT
created BIGINT DEFAULT date_part('epoch'::text, now()),
);
