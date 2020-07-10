-- ----------------------------
-- Table structure for day 主题党课
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_day(
   id INT(11),
   time INT(11) NOT NULL COMMENT '时间',
   address VARCHAR(50) NOT NULL COMMENT '地点',
   name VARCHAR(50) NOT NULL COMMENT '活动名称',
   content TEXT NOT NULL COMMENT '活动内容',
   joined VARCHAR(50) NOT NULL COMMENT '参加对象',
   record VARCHAR(30) NOT NULL COMMENT '记录人',
   status TINYINT(2) NOT NULL COMMENT '1正常显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);