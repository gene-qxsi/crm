CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(24) NOT NULL UNIQUE,
    age INT NOT NULL CHECK(age >= 0 AND age <= 117)
)