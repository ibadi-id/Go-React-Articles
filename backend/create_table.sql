DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts`  (
  `id` int(11) NOT null AUTO_INCREMENT,
  `title` varchar(200) NOT null,
  `content` text NOT null,
  `category` varchar(100) NOT null,
  `created_date` timestamp NOT null DEFAULT NOW(),
  `updated_date` timestamp NOT null DEFAULT NOW(),
  `status` ENUM('Publish', 'Draft', 'Trash') NOT null default 'Draft',
  PRIMARY KEY (`id`)
)