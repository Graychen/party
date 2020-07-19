-- ----------------------------
-- Table structure for user_history 用户历史记录表  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_user_history(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   url varchar(50) NOT NULL COMMENT '阅读url',
   user_id INT(11) NOT NULL COMMENT '操作用户id',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);