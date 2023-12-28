-- +migrate up

create table
    if not exists `pw_user_audit` (
        `id` int unsigned not null auto_increment primary key,
        `nickname` varchar(600) not null comment '昵称',
        `username` varchar(32) not null comment '真实姓名',
        `sex` tinyint default '1' comment '性别:1.男,2:女',
        `mobile` varchar(11) not null comment '手机号',
        `work` varchar(255) default null comment '职业',
        `education` varchar(255) default null comment '学历',
        `address` text default null comment '地址',
        `birthday` timestamp not null comment '生日',
        `loginip` varchar(50) not null comment '注册IP',
        `status` tinyint default '1' comment '状态:0:代审,1.通过,2.不通过',
        `photo` varchar(600) default null comment '照片',
        `card_front` varchar(600) default null comment '身份证正面',
        `card_back` varchar(600) default null comment '身份证反面',
        `createtime` int DEFAULT NULL COMMENT '创建时间',
        `updatetime` int DEFAULT NULL COMMENT '更新时间',
        index idx_created_at(createtime),
        index idx_updated_at(updatetime),
        key `username` (`username`) using btree,
        key `mobile` (`mobile`) using btree
    ) engine = innodb default charset = utf8mb4 comment = '会员审核表';

create table
    if not exists `pw_article` (
        `id` int unsigned not null auto_increment primary key,
        `title` varchar(255) not null comment '标题',
        `content` text comment '内容',
        `createtime` timestamp not null default current_timestamp,
        `updatetime` timestamp not null on update current_timestamp default current_timestamp,
        index idx_created_at(createtime),
        index idx_updated_at(updatetime),
        key `title` (`title`) using btree
    ) engine = innodb default charset = utf8mb4 comment = '文章表';

create table
    if not exists `pw_province_data`(
        `area_id` int unsigned not null auto_increment primary key,
        `area_type` tinyint default '1' comment '类型',
        `name` varchar(255) default null comment '名称',
        `parent_area_id` int unsigned default null comment '父级id',
        key `name` (`name`) using btree
    ) engine = innodb default charset = utf8mb4 comment = '省表';

create table
    if not exists `pw_city_data`(
        `area_id` int unsigned not null auto_increment primary key,
        `area_type` tinyint default '1' comment '类型',
        `name` varchar(255) default null comment '名称',
        `parent_area_id` int unsigned default null comment '父级id',
        key `name` (`name`) using btree
    ) engine = innodb default charset = utf8mb4 comment = '市表';

create table
    if not exists `pw_county_data`(
        `area_id` int unsigned not null auto_increment primary key,
        `area_type` tinyint default '1' comment '类型',
        `name` varchar(255) default null comment '名称',
        `parent_area_id` int unsigned default null comment '父级id',
        key `name` (`name`) using btree
    ) engine = innodb default charset = utf8mb4 comment = '区表';

-- +migrate down

drop table if exists `pw_user_audit`;

drop table if exists `pw_article`;

drop table if exists `pw_province_data`;

drop table if exists `pw_city_data`;

drop table if exists `pw_county_data`;