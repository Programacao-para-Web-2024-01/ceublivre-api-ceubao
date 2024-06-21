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
                                                   `image` VARCHAR(10000) NOT NULL,
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
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`,`Image`) VALUES ('Monaliza', 'Quadro', 15.60, 1,'https://s2-techtudo.glbimg.com/SSAPhiaAy_zLTOu3Tr3ZKu2H5vg=/0x0:1024x609/888x0/smart/filters:strip_icc()/i.s3.glbimg.com/v1/AUTH_08fbf48bc0524877943fe86e43087e7a/internal_photos/bs/2022/c/u/15eppqSmeTdHkoAKM0Uw/dall-e-2.jpg');
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`,`Image`) VALUES ('Fusca', 'Carro', 13000.00, 3,'https://upload.wikimedia.org/wikipedia/commons/2/22/Chevrolet_Opala_SS_1978_Inca.jpg');
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`,`Image`) VALUES ('Processador', 'PC', 775.00, 3,'https://s2-autoesporte.glbimg.com/x8eFp8dCk6dEurWSZ3VRK0xiEkg=/0x0:1298x865/888x0/smart/filters:strip_icc()/i.s3.glbimg.com/v1/AUTH_cf9d035bf26b4646b105bd958f32089d/internal_photos/bs/2023/a/F/bNed1NR5KVKNqlvDiYrQ/8731b386-2f60-4690-80cf-1f9359c38864-medium.jpeg');
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`,`Image`) VALUES ('Opala', 'Carro', 15000, 2,'https://encrypted-tbn3.gstatic.com/shopping?q=tbn:ANd9GcQHQzno7_zyLtqUwdI4o-6KSArMLyTpkedDOOqSYGPWfXhgL91zcbYM1OenwRuU0CLC2xiP-SXIo8BYmNtSzZbzcRbBTRppLNJn_i6Vx_JCPlTtsfBnxjml7PAizRNcvEOXCgFsGA&usqp=CAc');
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`,`Image`) VALUES ('Placa mãe', 'PC', 2000.00, 3,'https://encrypted-tbn1.gstatic.com/shopping?q=tbn:ANd9GcR7wbzvtJaEQmLs7PJnfom_D-Ua_jAZS1BCawB3kzM5x0HnCnJy0Ny24mFgAC-WBC8wdPGFk4fmqfyhhydR9dnXY68TX7ByiObZXODHdTjdry7SjF5Xr9jTe-l4Koe3syV7f6yZfi4&usqp=CAc');
INSERT INTO `market`.`products` (`name`, `description`, `price`, `categoria_idcategoria`,`Image`) VALUES ('Memoria ram', 'PC', 3300.00, 3,'https://encrypted-tbn0.gstatic.com/shopping?q=tbn:ANd9GcTzJ6L1R8XNgqG6H3XgMoK0XSJKOjHBs2OcsccM_cCA45fhXl3fe0g6PYdcD_Dp4Z73rwOFusNkVT9tVV4DgVNIbuXFG2Zgr1ZHStzgO8krkVsblNVztXfT3au2nBqy0ULuLIWlZc4&usqp=CAc');