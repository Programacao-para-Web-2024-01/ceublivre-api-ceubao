SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

CREATE SCHEMA IF NOT EXISTS `market` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `market` ;

CREATE TABLE IF NOT EXISTS `market`.`categoria` (
                                                    `idcategoria` INT NOT NULL,
                                                    `Nome_Categoria` VARCHAR(45) NOT NULL,
                                                    PRIMARY KEY (`idcategoria`))
    ENGINE = InnoDB;
CREATE TABLE IF NOT EXISTS `market`.`products` (
                                                   `id` INT NOT NULL AUTO_INCREMENT,
                                                   `name` VARCHAR(255) NOT NULL,
                                                   `description` VARCHAR(255) NOT NULL,
                                                   `price` FLOAT(10,2) NOT NULL,
                                                   `categoria_idcategoria` INT NOT NULL,
                                                   PRIMARY KEY (`id`),
                                                   INDEX `fk_products_categoria_idx` (`categoria_idcategoria` ASC) VISIBLE,
                                                   CONSTRAINT `fk_products_categoria`
                                                       FOREIGN KEY (`categoria_idcategoria`)
                                                           REFERENCES `market`.`categoria` (`idcategoria`)
                                                           ON DELETE NO ACTION
                                                           ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- Inserir valores na tabela categoria
INSERT INTO `market`.`categoria` (`idcategoria`, `Nome_Categoria`) VALUES (1, 'Eletrônicos');
INSERT INTO `market`.`categoria` (`idcategoria`, `Nome_Categoria`) VALUES (2, 'Roupas');
INSERT INTO `market`.`categoria` (`idcategoria`, `Nome_Categoria`) VALUES (3, 'Alimentos');

-- Inserir valores na tabela products
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`) VALUES ('Monaliza', 'Quadro', 15.60, 1);
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`) VALUES ('Opala', 'Carro', 15000, 2);
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`) VALUES ('Fusca', 'Carro', 13000.00, 3);
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`) VALUES ('Processador', 'PC', 775.00, 3);
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`) VALUES ('Placa mãe', 'PC', 2000.00, 3);
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`) VALUES ('Memoria ram', 'PC', 3300.00, 3);