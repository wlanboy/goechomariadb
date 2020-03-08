CREATE DATABASE 'goevent';
CREATE USER 'goevent' IDENTIFIED BY 'goevent';
GRANT USAGE ON *.* TO 'goevent'@'%' IDENTIFIED BY 'goevent';
GRANT ALL privileges ON `goevent`.* TO 'goevent'@'%';
FLUSH PRIVILEGES;
SHOW GRANTS FOR 'goevent'@'%'; 