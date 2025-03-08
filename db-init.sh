#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE TABLE IF NOT EXISTS wallets (
    	wallet_id uuid primary key,
    	amount numeric(32, 2) not null default 0.00
    );

    INSERT INTO
    	wallets
    VALUES
    	('acabd643-5065-4e7a-83a2-bb9245f4cca9', 40000),
    	('1115d743-1a60-441b-aad7-5b1662086af1', 0),
    	('e0e32c12-12b6-4612-aace-210bc997c4d4', -500),
    	('09e711a7-7812-497e-9661-7aea014a8a56', 250.95),
    	('99aa1f90-25c0-4fd1-a141-299af5f57fb7', 0.50);
EOSQL