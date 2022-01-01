#!/bin/bash

mysql --defaults-extra-file=/docker-entrypoint-initdb.d/my.conf <<EOS
DROP DATABASE IF EXISTS \`database\`;
CREATE DATABASE IF NOT EXISTS \`database\` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOS
mysql --defaults-extra-file=/docker-entrypoint-initdb.d/my.conf database < /docker-entrypoint-initdb.d/create_table.ddl
