-- ----------------------------
-- Table structure for user_info 用户信息表  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_user_info(
   id INT(11),
   user_id INT(11) NOT NULL COMMENT '用户id',
   tel INT(11) NOT NULL COMMENT '分机号（仅限企业内部开发调用)',
   work_place VARCHAR(50) NOT NULL COMMENT '办公地点',
   mobile INT(11) NOT NULL COMMENT '手机号码',
   email varchar(30) NOT NULL COMMENT '员工的电子邮箱',
   jobnumber varchar(80) NOT NULL COMMENT '员工工号',
   is_senior TINYINT(2) NOT NULL COMMENT '是否是高管模式 0:不是,1:是',
   hiredDate INT(11) NOT NULL COMMENT '入职时间',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);