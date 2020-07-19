-- ----------------------------
-- Table structure for apply 封面图表  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_banner(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   img_url varchar(11) NOT NULL COMMENT '图片url',
   user_id INT(11) NOT NULL COMMENT '操作用户id',
   status TINYINT(2) NOT NULL COMMENT '1显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);