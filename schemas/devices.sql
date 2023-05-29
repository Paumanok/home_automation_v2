CREATE TABLE IF NOT EXISTS devices (
    id BIGSERIAL PRIMARY KEY,
    nickname TEXT,
    mac TEXT,
    humiditycomp SMALLINT,
    temperaturecomp SMALLINT
);