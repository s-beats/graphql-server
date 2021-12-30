DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users (
  `id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `name` VARCHAR(20) NOT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8 COLLATE = utf8_bin;
DROP TABLE IF EXISTS task_priorities;
CREATE TABLE IF NOT EXISTS task_priorities (
  `id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `value` VARCHAR(20) NOT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8 COLLATE = utf8_bin;
DROP TABLE IF EXISTS tasks;
CREATE TABLE IF NOT EXISTS tasks (
  `id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `title` VARCHAR(20) NOT NULL,
  `text` VARCHAR(20) NOT NULL,
  `user_id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `priority_id` char(36) COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` Datetime DEFAULT NULL,
  `updated_at` Datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`id`),
  FOREIGN KEY (`priority_id`) REFERENCES task_priorities(`id`)
) DEFAULT CHARSET = utf8 COLLATE = utf8_bin;