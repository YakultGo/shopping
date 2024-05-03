create database if not exists shop;
use shop;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for inventory
-- ----------------------------
DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory`(
    `id`      bigint(20) NOT NULL AUTO_INCREMENT,
    `good_id` bigint(20) DEFAULT NULL COMMENT '商品id',
    `stock`   int(11) DEFAULT NULL COMMENT '库存',
    FOREIGN KEY (`good_id`) REFERENCES `goods` (`id`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
