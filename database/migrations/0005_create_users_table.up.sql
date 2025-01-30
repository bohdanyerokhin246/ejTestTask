CREATE TABLE IF NOT EXISTS users (
                                       login TEXT PRIMARY KEY,
                                       password TEXT NOT NULL,
                                       role TEXT NOT NULL
);

INSERT INTO users (login, password, role) VALUES ('admin','admin123','admin');
INSERT INTO users (login, password, role) VALUES ('user','user123','user');
