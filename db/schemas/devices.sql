CREATE TABLE IF NOT EXISTS devices (
    id BIGSERIAL PRIMARY KEY,
    nickname TEXT,
    mac TEXT,
    humidityComp SMALLINT,
    temperatureComp SMALLINT
);
