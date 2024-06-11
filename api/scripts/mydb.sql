
-- run this to test the database connection
-- this is where you put your dumps (/api/scripts)

-- Remember that,

CREATE TABLE users (
    username VARCHAR(50) PRIMARY KEY ,
    password VARCHAR(20)
 );


create table notification (
    notification_type VARCHAR(50),
    notification_content VARCHAR(50),
    timestamp VARCHAR(50),
    username VARCHAR(50)
);


insert into notification values
('normal','You need to code Flutter today','2024-05-23','masterplan'),
('normal','You finished your website today, Congrats!','2024-06-12','masterplan');


insert into users values
('masterplan','master123'),
('vincent','Tanzania');



