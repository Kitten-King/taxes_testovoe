CREATE SCHEMA IF NOT EXISTS testovoe;

CREATE TABLE testovoe.tax_rates (
                                    city_id INTEGER PRIMARY KEY,
                                    rate    FLOAT NOT NULL
);