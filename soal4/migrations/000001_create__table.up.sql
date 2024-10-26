CREATE TABLE requests (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);
