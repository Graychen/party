-- ----------------------------
-- Table structure for dynamic 党建动态
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_dynamic(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   title VARCHAR(100) NOT NULL COMMENT '党建标题',
   content TEXT NOT NULL COMMENT '党建内容',
   num INT(11) NOT NULL COMMENT '阅读数量',
   list_img_url VARCHAR(30) COMMENT '封面图片url',
   status TINYINT(2) NOT NULL COMMENT '1正常显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);