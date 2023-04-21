CREATE TABLE post
(
  `post_id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,  
  `subtitle` VARCHAR(255) NOT NULL,
  `tag` VARCHAR(255) NOT NULL,
  `img_modifier` VARCHAR(255),
  `author` VARCHAR(255),
  `author_url` VARCHAR(255),
  `publish_date` VARCHAR(255),
  `img_url` VARCHAR(255),
  `img_alt` VARCHAR(255),
  `featured` TINYINT(1) DEFAULT 0,
  PRIMARY KEY (`post_id`)
);