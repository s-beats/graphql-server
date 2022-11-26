#!/bin/bash

MYSQL_PWD=password
export MYSQL_PWD

mysql -u root <<EOS
DROP DATABASE IF EXISTS \`graphql-todo\`;
CREATE DATABASE IF NOT EXISTS \`graphql-todo\` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOS
mysql -u root graphql-todo < /var/ddl/schema.sql
