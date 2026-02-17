-- Create database
CREATE DATABASE IF NOT EXISTS sysadmin_portfolio CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Create user
CREATE USER IF NOT EXISTS 'sysadmin'@'%' IDENTIFIED BY 'sysadmin_password';

-- Grant privileges
GRANT ALL PRIVILEGES ON sysadmin_portfolio.* TO 'sysadmin'@'%';

-- Flush privileges
FLUSH PRIVILEGES;

-- Use database
USE sysadmin_portfolio;
