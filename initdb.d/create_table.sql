DROP TABLE IF EXISTS users;
create table IF not exists users (
  `id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8 COLLATE = utf8_bin;
DROP TABLE IF EXISTS tasks;
create table IF not exists tasks (
  `id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `title` VARCHAR(20) NOT NULL,
  `text` VARCHAR(20) NOT NULL,
  `user_id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`id`)
) DEFAULT CHARSET = utf8 COLLATE = utf8_bin;