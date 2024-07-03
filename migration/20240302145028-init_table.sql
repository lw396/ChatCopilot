-- +migrate Up 
CREATE TABLE if NOT EXISTS `group_contact` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `usr_name` VARCHAR(32) NOT NULL COMMENT '用户名称',
    `nickname` VARCHAR(255) NOT NULL COMMENT '昵称',
    `head_img` VARCHAR(600) NOT NULL COMMENT '头像',
    `group_member` TEXT NOT NULL COMMENT '群成员',
    `db_name` VARCHAR(255) NOT NULL COMMENT '数据库名称',
    `status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '状态',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_group_contact_usr_name` (`usr_name`),
    KEY `idx_group_contact_deleted_at` (`deleted_at`) USING BTREE
);

CREATE TABLE if NOT EXISTS `contact_person` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `usr_name` VARCHAR(32) NOT NULL COMMENT '用户名称',
    `nickname` VARCHAR(255) NOT NULL COMMENT '昵称',
    `remark` VARCHAR(255) NOT NULL COMMENT '备注',
    `head_img_url` VARCHAR(600) NOT NULL COMMENT '头像',
    `sex` TINYINT(1) UNSIGNED NOT NULL COMMENT '性别',
    `type` INT(10) UNSIGNED NOT NULL COMMENT '类型',
    `db_name` VARCHAR(255) NOT NULL COMMENT '数据库名称',
    `status` TINYINT(1) NOT NULL DEFAULT '1' COMMENT '状态',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_contact_person_usr_name` (`usr_name`),
    KEY `idx_contact_person_deleted_at` (`deleted_at`) USING BTREE
);

CREATE TABLE if NOT EXISTS `prompt_curation` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL COMMENT '标题',
    `prompt` LONGTEXT NOT NULL COMMENT '提示词',
    `start` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_contact_person_deleted_at` (`deleted_at`) USING BTREE
);

CREATE TABLE IF NOT EXISTS `chat_record` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(255) NOT NULL COMMENT '用户名称',
    `prompt_id` LONGTEXT NOT NULL COMMENT '提示词',
    `messages` JSON NOT NULL COMMENT '聊天记录',
    `start` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_contact_person_deleted_at` (`deleted_at`) USING BTREE
);

-- +migrate Down
DROP TABLE IF EXISTS `group_contact`;

DROP TABLE IF EXISTS `contact_person`;

DROP TABLE IF EXISTS `prompt_curation`;