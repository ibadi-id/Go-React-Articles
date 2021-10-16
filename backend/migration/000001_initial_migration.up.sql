-- ----------------------------
-- Table structure for posts
-- ----------------------------

CREATE TABLE IF NOT EXISTS `posts`  (
  `id` int(11) NOT null AUTO_INCREMENT,
  `title` varchar(200) NOT null,
  `content` text NOT null,
  `category` varchar(100) NOT null,
  `created_date` timestamp NOT null DEFAULT NOW(),
  `updated_date` timestamp NOT null DEFAULT NOW(),
  `status` ENUM('Publish', 'Draft', 'Trash') NOT null default 'Draft',
  PRIMARY KEY (`id`)
);

-- ----------------------------
-- Initial of posts
-- ----------------------------

INSERT INTO posts (id,title,content,category,status) VALUES
(1,'Article 1','Contrary to popular belief, Lorem Ipsum is not simply random text','berita', 'Publish'),
(2,'Article 2','Contrary to popular belief, Lorem Ipsum is not simply random text','olahraga', 'Publish'),
(3,'Article 3','Contrary to popular belief, Lorem Ipsum is not simply random text','berita', 'Publish'),
(4,'Article 4','Contrary to popular belief, Lorem Ipsum is not simply random text','kesehatan', 'Publish'),
(5,'Article 5','Contrary to popular belief, Lorem Ipsum is not simply random text','hobi', 'Publish'),
(6,'Article 6','Contrary to popular belief, Lorem Ipsum is not simply random text','automotif', 'Publish'),
(7,'Article 7','Contrary to popular belief, Lorem Ipsum is not simply random text','kesehatan', 'Publish'),
(8,'Article 8','Contrary to popular belief, Lorem Ipsum is not simply random text','hobi', 'Draft'),
(9,'Article 9','Contrary to popular belief, Lorem Ipsum is not simply random text','berita', 'Trash'),
(10,'Article 10','Contrary to popular belief, Lorem Ipsum is not simply random text','kesehatan', 'Draft');