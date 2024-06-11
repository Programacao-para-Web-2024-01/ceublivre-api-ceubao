CREATE DATABASE IF NOT EXISTS market;

USE market;

CREATE TABLE IF NOT EXISTS products
(
    id          INT          PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    price       FLOAT(5,2)   NOT NULL
);
