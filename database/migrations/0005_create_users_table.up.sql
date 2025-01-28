CREATE TABLE IF NOT EXISTS users (
                                       login TEXT PRIMARY KEY,
                                       password TEXT NOT NULL,
                                       role TEXT NOT NULL
);