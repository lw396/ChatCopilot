-- +migrate Up 
CREATE TABLE if NOT EXISTS `group_contact` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `usr_name` VARCHAR(32) UNIQUE NOT NULL COMMENT '用户名称',
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
    `usr_name` VARCHAR(32) UNIQUE NOT NULL COMMENT '用户名称',
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

CREATE TABLE IF NOT EXISTS `chat_copilot` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `usr_name` VARCHAR(255) UNIQUE NOT NULL COMMENT '用户名称',
    `type` ENUM('person', 'group') NOT NULL COMMENT '聊天类型',
    `prompt_id` BIGINT UNSIGNED NOT NULL COMMENT '提示词id',
    `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_chat_copilot_deleted_at` (`deleted_at`) USING BTREE
);

CREATE TABLE IF NOT EXISTS `copilot_config` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `api_type` INT UNSIGNED NOT NULL COMMENT 'api类型',
    `url` VARCHAR(255) NOT NULL COMMENT '链接',
    `token` VARCHAR(255) NOT NULL COMMENT '令牌',
    `model` VARCHAR(100) NOT NULL COMMENT '模型名称',
    `temperature` DOUBLE NOT NULL DEFAULT 0.9 COMMENT '温度',
    `top_p` DOUBLE NOT NULL DEFAULT 0.7 COMMENT 'top_p',
    `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL,
    INDEX `idx_model` (`model`),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_copilot_config_deleted_at` (`deleted_at`) USING BTREE
);

-- +migrate Down
DROP TABLE IF EXISTS ` group_contact `;

DROP TABLE IF EXISTS ` contact_person `;

DROP TABLE IF EXISTS ` prompt_curation `;