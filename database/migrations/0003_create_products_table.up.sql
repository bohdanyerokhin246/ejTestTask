CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        name TEXT NOT NULL,
                                        description TEXT,
                                        price NUMERIC NOT NULL,
                                        quantity INT NOT NULL,
                                        seller_id INT NOT NULL,
                                        FOREIGN KEY (seller_id) REFERENCES sellers(id)
);

INSERT INTO products (name, description, price, quantity,seller_id) VALUES ('Apple','This is green apple','10.00','20','1');
INSERT INTO products (name, description, price, quantity,seller_id) VALUES ('Iphone','This is Apple smartphone','350.00','20','2');
INSERT INTO products (name, description, price, quantity,seller_id) VALUES ('Laptop','This is new laptop','500.00','10','2');
INSERT INTO products (name, description, price, quantity,seller_id) VALUES ('BMW','This is fast car','10000.00','5','3');