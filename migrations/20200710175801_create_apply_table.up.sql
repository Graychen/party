-- ----------------------------
-- Table structure for apply 报名表  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_apply(
   id INT(11),
   activity_id INT(11) NOT NULL COMMENT '活动id',
   user_id INT(11) NOT NULL COMMENT '用户id',
   status TINYINT(2) NOT NULL COMMENT '1通过,0不通过',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);