-- Currently the table is set, and any migration will be added and made manually. Room for improvement there
CREATE TABLE IF NOT EXISTS
  users (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    games_played INTEGER,
    score INTEGER,
    friends text
  );
