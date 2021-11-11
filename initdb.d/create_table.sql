DROP TABLE IF EXISTS users;

create table IF not exists users
(
    `id`               INT(20) AUTO_INCREMENT,
    `name`             VARCHAR(20) NOT NULL,
    `created_at`       Datetime DEFAULT NULL,
    `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

DROP TABLE IF EXISTS tasks;

create table IF not exists tasks
(
    `id`               INT(20) AUTO_INCREMENT,
    `title`             VARCHAR(20) NOT NULL,
    `text`             VARCHAR(20) NOT NULL,
    `user_id`               INT(20) NOT NULL,
    `created_at`       Datetime DEFAULT NULL,
    `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY user_id(`user_id`) REFERENCES users(`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
