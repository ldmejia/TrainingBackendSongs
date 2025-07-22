CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    duration VARCHAR(50) NOT NULL,
    album VARCHAR(255),
    artwork VARCHAR(255),
    price DECIMAL(10, 2),
    origin VARCHAR(255)
);
