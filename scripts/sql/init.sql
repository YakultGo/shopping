DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `address` varchar(255) DEFAULT NULL,
    `telephone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
    `id_card` varchar(30) DEFAULT NULL,
    `birthday` date DEFAULT NULL,
    `delete_at` time DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `telephone` (`telephone`),
    UNIQUE KEY `id_card` (`id_card`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;