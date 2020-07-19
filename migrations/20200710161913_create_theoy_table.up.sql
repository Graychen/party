-- ----------------------------
-- Table structure for theoy 理论分享
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_theoy(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   title VARCHAR(100) NOT NULL COMMENT '理论标题',
   content TEXT NOT NULL COMMENT '理论内容',
   num INT(11) NOT NULL COMMENT '阅读数量',
   status TINYINT(2) NOT NULL COMMENT '1正常显示,0不显示',
   list_img_url VARCHAR(30) COMMENT '封面图片url',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);