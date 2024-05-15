CREATE DATABASE kaspi_db OWNER postgres;

CREATE TABLE IF NOT EXISTS kaspi_db.peoples (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    iin     VARCHAR(12) NOT NULL,
    phone   VARCHAR(15) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX peoples_idx_iin ON kaspi_db.peoples (iin);
CREATE INDEX peoples_idx_phone ON kaspi_db.peoples (phone)