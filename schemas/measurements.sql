CREATE TABLE IF NOT EXISTS Measurements (
    id BIGSERIAL PRIMARY KEY,
    mac TEXT,
    temp REAL,
    humidity REAL,
    pressure REAL,
    pm25 REAL,
    createdat timestamptz DEFAULT now()
)