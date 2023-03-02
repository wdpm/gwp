drop database if exists gwp;
create database gwp;
drop user if exists gwp;
create user gwp with password '123456';
grant all privileges on database gwp to gwp;
alter role gwp superuser;