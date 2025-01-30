-- Создаем базу данных
CREATE DATABASE db_auth;

-- Назначаем существующего пользователя postgres владельцем базы данных
GRANT ALL PRIVILEGES ON DATABASE db_auth TO postgres;
