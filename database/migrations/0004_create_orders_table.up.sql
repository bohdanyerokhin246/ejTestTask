CREATE TABLE IF NOT EXISTS orders (
                                      id SERIAL PRIMARY KEY,
                                      buyer_id INT NOT NULL,
                                      price NUMERIC NOT NULL ,
                                      products JSONB NOT NULL,
                                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      FOREIGN KEY (buyer_id) REFERENCES buyers(id)
);

INSERT INTO orders (buyer_id, price, products)
VALUES (1, 370.00, '[{"ID": 1, "price": 10.00, "quantity": 2}, {"ID": 2, "price": 350.00, "quantity": 1}]');

INSERT INTO orders (buyer_id, price, products)
VALUES (1, 100.00, '[{"ID": 1, "price": 10.00, "quantity": 10}]');

INSERT INTO orders (buyer_id, price, products)
VALUES (2, 850.00, '[{"ID": 2, "price": 350.00, "quantity": 1}, {"ID": 3, "price": 500.00, "quantity": 1}]');

INSERT INTO orders (buyer_id, price, products)
VALUES (3, 10850.00, '[{"ID": 2, "price": 350.00, "quantity": 1}, {"ID": 3, "price": 500.00, "quantity": 1},{"ID": 4, "price": 10000.00, "quantity": 1}]');