create table wallet (
	address varchar(58) not null primary key,
	type enum('foundation', 'cex', 'inc') not null,

    index (type)
);