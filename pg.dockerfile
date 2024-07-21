FROM postgres:15.3

COPY ./schemas/measurements.sql /docker-entrypoint-initdb.d/
COPY ./schemas/devices.sql /docker-entrypoint-initdb.d/
