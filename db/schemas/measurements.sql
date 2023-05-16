CREATE TABLE IF NOT EXISTS measurements (
    id BIGSERIAL PRIMARY KEY,
    mac TEXT,
    temp REAL,
    humidity REAL,
    pressure REAL,
    pm25 REAL,
    createdAt timestamptz DEFAULT now()
)