CREATE DATABASE kaspi_db OWNER postgres;

CREATE SCHEMA IF NOT EXISTS kaspi;

CREATE TABLE IF NOT EXISTS kaspi.users (
                                           id      SERIAL PRIMARY KEY,
                                           first_name    VARCHAR(255),
                                           last_name    VARCHAR(255),
                                           patronymic    VARCHAR(255),
                                           fullname text generated always as (last_name || ' ' || first_name || ' ' || patronymic) stored,
                                           iin     VARCHAR(12) NOT null,
                                           phone   VARCHAR(15) DEFAULT NULL,
                                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

                                           CONSTRAINT iin_uq UNIQUE (iin)
);

CREATE INDEX peoples_idx_iin ON kaspi.users (iin);
CREATE INDEX peoples_idx_phone ON kaspi.users (phone);
CREATE INDEX peoples_idx_name ON kaspi.users (fullname);