CREATE DATABASE IF NOT EXISTS social_network;
USE social_network;

DROP WITH EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;