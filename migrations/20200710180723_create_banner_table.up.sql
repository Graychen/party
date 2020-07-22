-- ----------------------------
-- Table structure for apply 封面图表  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_banner(
   id INT(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
   img_url varchar(255) NOT NULL COMMENT '图片url',
   user_id INT(11) NOT NULL COMMENT '操作用户id',
   type INT(11) NOT NULL COMMENT '1：党建工作室 2：活动报名 3：理论分享 4：主题党日 5：党建动态',
   status TINYINT(2) NOT NULL COMMENT '1显示,0不显示',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);