CREATE KEYSPACE IF NOT EXISTS apartments
WITH replication = {
  'class': 'SimpleStrategy',
  'replication_factor': 1
};

USE apartments;

CREATE TABLE IF NOT EXISTS user_matches (
    user_id int,
    week int,
    apartment_ids list<int>,
    PRIMARY KEY (user_id, week)
);
