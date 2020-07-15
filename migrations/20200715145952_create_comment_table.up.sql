-- ----------------------------
-- Table structure for comment  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_comment(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   mien_id INT(11) NOT NULL COMMENT '关联活动风采的id',
   type INT(11) NOT NULL COMMENT '专题评论类型1活动风采',
   content TEXT COMMENT '风采内容',
   img_url TEXT COMMENT '评论带的图片',
   status TINYINT(2) NOT NULL COMMENT '1正常显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id ) USING BTREE
);