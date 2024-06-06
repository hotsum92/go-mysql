-- create test database
CREATE DATABASE test;

use test;

-- create user table
CREATE TABLE user (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

-- insert some data
INSERT INTO user (name, email) VALUES ('Alice', 'mail');
