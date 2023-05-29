FROM postgres:latest

COPY ./schemas/measurements.sql /docker-entrypoint-initdb.d/
COPY ./schemas/devices.sql /docker-entrypoint-initdb.d/