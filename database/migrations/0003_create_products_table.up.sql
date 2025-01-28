CREATE TABLE IF NOT EXISTS products (
                                        id SERIAL PRIMARY KEY,
                                        name TEXT NOT NULL,
                                        description TEXT,
                                        price NUMERIC NOT NULL,
                                        seller_id INT NOT NULL,
                                        FOREIGN KEY (seller_id) REFERENCES sellers(id)
);