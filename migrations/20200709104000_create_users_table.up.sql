CREATE TABLE IF NOT EXISTS party_users(
   id INT,
   name VARCHAR(100) NOT NULL,
   password VARCHAR(40) NOT NULL,
   PRIMARY KEY ( id )
);