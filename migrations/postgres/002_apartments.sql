CREATE TABLE IF NOT EXISTS apartments (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    price_per_month NUMERIC NOT NULL,
    room_numbers INT NOT NULL,
    bedroom_numbers INT NOT NULL,
    bathroom_numbers INT NOT NULL,
    city TEXT NOT NULL,
    district TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- insert some test data
INSERT INTO apartments (title, price_per_month, room_numbers, bedroom_numbers, bathroom_numbers, city, district)
VALUES 
('Apartment 1', 1000, 3, 2, 1, 'Tbilisi', 'Saburtalo'),
('Apartment 2', 2000, 4, 3, 2, 'Kutaisi', 'Something');