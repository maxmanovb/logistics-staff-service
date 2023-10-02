-- support schema
CREATE SCHEMA IF NOT EXISTS support;

-- vehicle
-- Create the "Driver" table
CREATE TABLE IF NOT EXISTS support.driver (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(255) UNIQUE,
    password_hash VARCHAR(255),
    password_salt VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Create the "DriverProfile" table
CREATE TABLE IF NOT EXISTS support.driver_profile (
    driver_id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    phone_number VARCHAR(255) UNIQUE,
    image_url VARCHAR(255),
    active BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- Create the "DriverSmsCode" table
CREATE TABLE IF NOT EXISTS support.driver_sms_code (
    id SERIAL PRIMARY KEY,
    driver_id SERIAL REFERENCES driver(id),
    code VARCHAR(255),
    expires_in INT,
    created_at TIMESTAMP
);

-- Create an index on "driver" table for "phone_number"
CREATE INDEX IF NOT EXISTS idx_driver_by_phone ON driver(phone_number);

-- Create an index on "driver_sms_code" table for "driver_id"
CREATE INDEX IF NOT EXISTS idx_sms_code_by_driver ON driver_sms_code(driver_id);


