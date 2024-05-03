create database if not exists shop;
use shop;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL,
    `address` varchar(255) DEFAULT NULL,
    `telephone` varchar(30) NOT NULL ,
    `birthday` date DEFAULT NULL,
    `delete_at` time DEFAULT NULL,
    `password` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `telephone` (`telephone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;