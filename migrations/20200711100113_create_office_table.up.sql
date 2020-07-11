-- ----------------------------
-- Table structure for office 科室表  
-- ----------------------------
CREATE TABLE IF NOT EXISTS party_office(
   id INT(11),
   user_id INT(11) NOT NULL COMMENT '用户id',
   branch INT(11) NOT NULL COMMENT '第几支部',
   role INT(11) NOT NULL COMMENT '角色 1:支部书记,2:委员会成员,3:所以成员',
   is_secretary TINYINT(2) NOT NULL COMMENT '是否是支部书记 0:不是,1:是',
   is_committee TINYINT(2) NOT NULL COMMENT '是否是委员会成员 0:不是,1:是',
   is_member TINYINT(2) NOT NULL COMMENT '是否是所有成员 0:不是,1:是',
   create_time INT(11) unsigned DEFAULT 0 COMMENT '添加时间',
   PRIMARY KEY ( id )
);