-- ----------------------------
-- Table structure for activity 活动 
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_activity(
   id INT(11),
   theme VARCHAR(100) NOT NULL COMMENT '活动主题',
   time INT(11) NOT NULL COMMENT '活动时间',
   address VARCHAR(30) NOT NULL COMMENT '活动地址',
   content TEXT NOT NULL COMMENT '活动内容',
   totol INT(11) NOT NULL COMMENT '活动人数',
   status TINYINT(2) NOT NULL COMMENT '1正常显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);