-- +migrate Up 
CREATE TABLE if NOT EXISTS `group_contact` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `usr_name` VARCHAR(32) NOT NULL COMMENT '用户名称',
    `nickname` VARCHAR(255) NOT NULL COMMENT '昵称',
    `head_img` VARCHAR(600) NOT NULL COMMENT '头像',
    `group_member` TEXT NOT NULL COMMENT '群成员',
    `db_name` VARCHAR(255) NOT NULL COMMENT '数据库名称',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_group_contact_usr_name` (`usr_name`),
    KEY `idx_group_contact_deleted_at` (`deleted_at`) USING BTREE
);

-- +migrate Down
DROP TABLE IF EXISTS `group_contact`;