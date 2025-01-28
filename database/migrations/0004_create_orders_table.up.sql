CREATE TABLE IF NOT EXISTS orders (
                                      id SERIAL PRIMARY KEY,
                                      buyer_id INT NOT NULL,
                                      price INT NOT NULL ,
                                      products JSONB NOT NULL,
                                      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                      FOREIGN KEY (buyer_id) REFERENCES buyers(id)
);