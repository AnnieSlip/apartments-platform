CREATE TABLE IF NOT EXISTS user_filters (
    user_id INT REFERENCES users(id),
    min_price NUMERIC,
    max_price NUMERIC,
    room_numbers INT[],
    bedroom_numbers INT[],
    bathroom_numbers INT[],
    city TEXT,
    district TEXT,
    PRIMARY KEY (user_id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
