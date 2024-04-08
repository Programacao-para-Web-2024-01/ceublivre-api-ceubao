CREATE DATABASE IF NOT EXISTS CatalogoDeProdutos;

USE CatalogoDeProdutos;

CREATE TABLE Produtos
(
    id    INT PRIMARY KEY AUTO_INCREMENT,
    name  VARCHAR(255) NOT NULL,
    Descricao   VARCHAR(255)          NOT NULL,
    Preco float NOT NULL,
    Categoria VARCHAR(255)  NOT NULL
);