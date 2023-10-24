-- new db

-- subscription
-- subscription_test

create table subscription_factory (
    -- id is the onchain index of the subscription contract
    id bigint unsigned not null,
    -- version represents the current factory version
    version int(2) not null,
    -- current_as_of_block represents the last time this subscription was checked on chain
    current_as_of_block bigint unsigned not null,

    primary key (id),
    index (current_as_of_block)
);

create table subscription (
    -- id is the onchain index of the subscription contract
    id bigint unsigned not null,
    -- creation date of the contract
    start_date datetime not null,
    -- box_index representing the offering index of the merchant
    box_index int(10) unsigned not null,
    -- interval that the subscription runs for (how often they should be charged)
    intervl int(10) unsigned not null,
    -- last_payment is a timestamp of the last successful payment
    last_payment datetime not null,
    -- subscriber is the address of the subscriber
    subscriber varchar(58) not null,
    -- recipient is the merchant or receiver of the payments
    recipient varchar(58) not null, 
    -- asa is the asset id of the payments
    asa bigint unsigned not null,
    -- amount is the quantity being sent of the asa in the lowest unit
    amount int(10) unsigned not null,
    -- pass_count is number of passes allowed for the subscription
    pass_count int(1) unsigned not null,
    -- last payout attempt is a timestamp of the last payout attempt that failed
    lpa datetime,
    -- current_as_of_block represents the last time this subscription was checked on chain
    current_as_of_block bigint unsigned not null,
    -- deleted represents whether the subscription has been deleted
    deleted boolean not null default 0,
    -- deleted date tracks when it was deleted to evaluate if theres any remaining valid subscription time
    deleted_date datetime,

    primary key (id),
    index (start_date),
    index (intervl),
    index (last_payment),
    index (subscriber),
    index (recipient),
    index (asa),
    index (amount),
    index (lpa),
    index (current_as_of_block)
);

create table subscription_pass (
    -- id is the id of the subscription contract
    id int(10) unsigned not null,
    address varchar(58) not null,

    primary key (id, address),
    index (id),
    index (address)
);

create table subscription_offering (
    -- recipient is the address of the merchant
    recipient varchar(58) not null,
    -- box_index representing the offering index
    box_index int(10) unsigned not null,
    -- interval
    intervl int(10) unsigned not null,
    -- asa
    asa bigint unsigned not null,
    -- amount
    amount int(10) unsigned not null,
    -- pass_count
    pass_count int(1) unsigned not null,
    -- cid
    cid varchar(46) not null,
    -- active - whether or not you can subscribe currently
    active boolean not null default 1,
    -- shutdown - whether or not the offering is permanently discontinued
    shutdown boolean not null default 0,

    primary key (recipient, box_index),
    index (recipient)
);

create table subscription_blacklist (
    -- recipient is the address of the merchant
    recipient varchar(58) not null,
    -- address thats blacklisted
    address varchar(58) not null,

    primary key (recipient, address)
);

create table subscription_payment (
    -- id is the transaction id of the payout
    id varchar(52) not null,
    -- sub_id is the onchain index of the subscription contract
    sub_id bigint unsigned not null,

    primary key (id),
    index (sub_id)
);