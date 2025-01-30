CREATE TABLE IF NOT EXISTS buyers (
                                      id SERIAL PRIMARY KEY,
                                      name TEXT NOT NULL,
                                      phone TEXT NOT NULL
);

INSERT INTO buyers (name, phone) VALUES ('Bohdan','380990000000');
INSERT INTO buyers (name, phone) VALUES ('Ihor','380970000000');
INSERT INTO buyers (name, phone) VALUES ('Andrey','380660000000');