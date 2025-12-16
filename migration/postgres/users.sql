CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY ,
    email TEXT UNIQUE NOT NULL
);

-- insert test user
INSERT INTO users (email) VALUES ('test@example.com');
