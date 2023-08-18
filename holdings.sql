create database holdings;

create table holdings (
    "address" varchar(58) not null,
    "asa_id" bigint unsigned not null,
    "amount" bigint unsigned,

    primary key (address, asa_id),
    index (address),
    index (asa_id),
    index (amount)
);