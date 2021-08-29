\c postgres
CREATE ROLE postgres with 
LOGIN
SUPERUSER
INHERIT
CREATEDB
CREATEROLE
REPLICATION
password 'postgres';
GRANT ALL PRIVILEGES ON DATABASE "postgres" TO postgres;

CREATE TABLE logbook (
	 id SERIAL PRIMARY KEY,
     cityname VARCHAR(64),
     timerequested VARCHAR(64),
     temperature REAL DEFAULT 0.0);