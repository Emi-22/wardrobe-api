CREATE TABLE closet_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    classification VARCHAR(50) NOT NULL,
    color VARCHAR(30) NOT NULL,
    brand VARCHAR(50),
    favorite BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
)