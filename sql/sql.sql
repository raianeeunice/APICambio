CREATE DATABASE IF NOT EXISTS cambio;
USE cambio;

DROP TABLE IF EXISTS depositos;
CREATE TABLE depositos(
    id int auto_increment primary Key,
    valorDeposito float not null,
    feitoEm timestamp default current_timestamp()
) ENGINE=INNODB;
