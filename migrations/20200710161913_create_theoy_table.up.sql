-- ----------------------------
-- Table structure for theoy 理论分享
-- ----------------------------
CREATE TABLE IF NOT EXISTS theoy(
   id INT(11),
   title VARCHAR(100) NOT NULL COMMENT '党课标题',
   content TEXT NOT NULL COMMENT '党课内容',
   num INT(11) NOT NULL COMMENT '阅读数量',
   status TINYINT(2) NOT NULL COMMENT '1正常显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);