CREATE TABLE IF NOT EXISTS party_users(
   id INT,
   name VARCHAR(100) NOT NULL,
   password VARCHAR(40) NOT NULL,
   unionid INT(11) NOT NULL COMMENT '员工在当前开发者企业账号范围内的唯一标识',
   PRIMARY KEY ( id )
);