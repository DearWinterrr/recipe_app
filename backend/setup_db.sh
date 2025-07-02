#!/bin/bash

# 数据库配置
DB_USER="root"
DB_PASSWORD="password"
DB_NAME="recipe_app"

# 创建数据库
echo "Creating database $DB_NAME if not exists..."
mysql -u$DB_USER -p$DB_PASSWORD -e "CREATE DATABASE IF NOT EXISTS $DB_NAME DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 导入初始数据
echo "Importing initial data..."
mysql -u$DB_USER -p$DB_PASSWORD $DB_NAME < init_data.sql

echo "Database setup completed!" 