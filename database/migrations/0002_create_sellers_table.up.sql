CREATE TABLE IF NOT EXISTS sellers (
                                       id SERIAL PRIMARY KEY,
                                       name TEXT NOT NULL,
                                       phone TEXT NOT NULL
);

INSERT INTO sellers (name, phone) VALUES ('Jack','17182222222');
INSERT INTO sellers (name, phone) VALUES ('John','17183333333');
INSERT INTO sellers (name, phone) VALUES ('Donald','17184444444');