CREATE TABLE `admin` (
  `admin_id` int(11) NOT NULL AUTO_INCREMENT,
  `admin_name` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `admin_password` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  `role_code` varchar(45) CHARACTER SET latin1 DEFAULT NULL,
  PRIMARY KEY (`admin_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;