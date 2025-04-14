CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email VARCHAR(255)UNIQUE,
                       hashed_password VARCHAR(255)
);

CREATE TABLE tasks (
                       id SERIAL PRIMARY KEY,
                       user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       name VARCHAR(255),
                       description TEXT,
                       category VARCHAR(255),
                       priority VARCHAR(50)CHECK (priority IN ('low', 'medium', 'high')),
                       deadline TIMESTAMP
);
