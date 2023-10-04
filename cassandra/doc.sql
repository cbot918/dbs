CREATE KEYSPACE IF NOT EXISTS history
WITH replication = {
    'class': 'SimpleStrategy',
    'replication_factor': 1
};

CREATE TABLE record (
   id int, 
   to_user int, 
   body text, 
   PRIMARY KEY (id)
);

-- 
DESCRIBE keyspace;

-- 
insert into record (id, to_user, body) values (1,1,'hihi');

-- 
select * from record;