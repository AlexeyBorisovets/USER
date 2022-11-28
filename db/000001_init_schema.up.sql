CREATE TABLE IF NOT EXISTS users(
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    balance INT NOT NULL,
    usertype TEXT NOT NULL,
    password TEXT NOT NULL,
    refreshtoken TEXT
    );