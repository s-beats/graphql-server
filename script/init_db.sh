#!/bin/bash

mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf <<EOS
DROP DATABASE IF EXISTS \`database\`;
CREATE DATABASE IF NOT EXISTS \`database\` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOS
mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf database < /var/ddl/task.ddl
