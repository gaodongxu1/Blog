CREATE DATABASE IF NOT EXISTS `Blog`;
USE `Blog`;

CREATE TABLE IF NOT EXISTS `article` (
	`Id` int(12) unsigned NOT NULL AUTO_INCREMENT,
	`Title` varchar(200) NOT NULL,
	`Content` text NOT NULL,
	`Status`  int(12) NOT NULL,
	`PublishedTime` varchar(60) NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
  


  CREATE TABLE IF NOT EXISTS `admin` (
  `id` int(12) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL UNIQUE ,
  `password` varchar(128) NOT NULL DEFAULT '',
  `created` datetime NOT NULL DEFAULT current_timestamp,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
 