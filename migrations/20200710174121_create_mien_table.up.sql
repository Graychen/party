-- ----------------------------
-- Table structure for team  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_mien(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   type INT(11) NOT NULL COMMENT '类型1团队风采,2活动风采',
   title VARCHAR(100) NOT NULL COMMENT '风采标题',
   content TEXT COMMENT '风采内容',
   num INT(11) NOT NULL COMMENT '阅读数量',
   list_img_url VARCHAR(255) COMMENT '封面图片url',
   status TINYINT(2) NOT NULL COMMENT '1正常显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);