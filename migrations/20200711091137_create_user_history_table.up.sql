-- ----------------------------
-- Table structure for user_history 用户历史记录表  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_user_history(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   type INT(11) NOT NULL COMMENT '1：党建工作室 2：活动报名 3：理论分享 4：主题党日 5：党建动态',
   user_id INT(11) NOT NULL COMMENT '操作用户id',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);