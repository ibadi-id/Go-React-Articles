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
(1,'Artikel dengan urutan 1','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','berita', 'Publish'),
(2,'Artikel dengan urutan 2','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','olahraga', 'Publish'),
(3,'Artikel dengan urutan 3','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','berita', 'Publish'),
(4,'Artikel dengan urutan 4','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','kesehatan', 'Publish'),
(5,'Artikel dengan urutan 5','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','hobi', 'Publish'),
(6,'Artikel dengan urutan 6','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','automotif', 'Publish'),
(7,'Artikel dengan urutan 7','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','kesehatan', 'Publish'),
(8,'Artikel dengan urutan 8','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','hobi', 'Draft'),
(9,'Artikel dengan urutan 9','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','berita', 'Trash'),
(10,'Artikel dengan urutan 10','Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi utul.','kesehatan', 'Draft');