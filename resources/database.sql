CREATE DATABASE cwh;
USE cwh;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `role` varchar(20),
  `email` varchar(100) NOT NULL,
  `attributes` JSON,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;
INSERT INTO `users` VALUES
		(1,"Rob", STR_TO_DATE('18/02/2019 11:15:45','%d/%m/%Y %H:%i:%s'), "rob@test.com", "dfsdf,dsfds",NULL, 1),
		(2,"Rob", STR_TO_DATE('18/02/2019 11:15:45','%d/%m/%Y %H:%i:%s'), "rob@test.com", "dfsdf,dsfds",NULL, 1),
		(3,"Rob", STR_TO_DATE('18/02/2019 11:15:45','%d/%m/%Y %H:%i:%s'), "rob@test.com", "dfsdf,dsfds",NULL, 1),
		(4,"Rob", STR_TO_DATE('18/02/2019 11:15:45','%d/%m/%Y %H:%i:%s'), "rob@test.com", "dfsdf,dsfds",NULL, 1),
		(5,"Rob", STR_TO_DATE('18/02/2019 11:15:45','%d/%m/%Y %H:%i:%s'), "rob@test.com", "dfsdf,dsfds",NULL, 1),
		(6,"Rob", STR_TO_DATE('18/02/2019 11:15:45','%d/%m/%Y %H:%i:%s'), "rob@test.com", "dfsdf,dsfds",NULL, 1),
		(7,"Rob", STR_TO_DATE('18/02/2019 11:15:45','%d/%m/%Y %H:%i:%s'), "rob@test.com", "dfsdf,dsfds",NULL, 1);

DROP TABLE IF EXISTS `groups`;
CREATE TABLE `groups` (
  `group_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `members` JSON,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `attributes` JSON ,
  PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `groups` VALUES
	(2, "test-2",NULL,STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'),NULL),
  (3, "test-3",NULL,STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'),NULL),
  (4, "test-4",NULL,STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'), NULL),
	(5, "test-5",NULL,STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'),NULL);

DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects` (
  `project_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `created_by` varchar(100) NOT NULL,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `group`  int(11),
  `repo_url` varchar(100) NOT NULL,
  `attributes` JSON,
  `activities` JSON,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`project_id`),
  KEY `projects_FK` (`group`),
  CONSTRAINT `projects_FK` FOREIGN KEY (`group`) REFERENCES `groups` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;
INSERT INTO `projects` VALUES
	(1, "test-2", "stan",STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'), 2, "http://www.bictbucket.com/opda/test",NULL,NULL,1),
  (2, "test-2", "stan",STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'), 3, "http://www.bictbucket.com/opda/test",NULL,NULL,1),
  (3, "test-2", "stan",STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'), 3, "http://www.bictbucket.com/opda/test",NULL,NULL,1),
	(4, "test-2", "stan",STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'), 4, "http://www.bictbucket.com/opda/test",NULL,NULL,1);
    
DROP TABLE IF EXISTS `environments`;
CREATE TABLE `environments` (
  `environment_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `project` int(11) NOT NULL ,
  `created_on` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `attributes` JSON,
  PRIMARY KEY (`environment_id`),
  KEY `environment_FK` (`project`),
  CONSTRAINT `environment_FK` FOREIGN KEY (`project`) REFERENCES `projects` (`project_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `environments` VALUES
  (1,"master",2,STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'), NULL),
  (2,"dev",3,STR_TO_DATE('01/01/2021','%d/%m/%Y %H:%i:%s'),NULL);