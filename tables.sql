CREATE TABLE "nfd" (
  "id" bigint unsigned not null,
  "asa_id" bigint unsigned not null,
  "state" varchar(32) not null,
  "current_as_of_block" bigint unsigned not null,
  "name" varchar(128) not null,
  "version" varchar(4) not null,
  "username" varchar(128) default null,
  "bio" text default null,
  "email" varchar(128) default null,
  "avatar" varchar(256) default null,
  "avatar_id" bigint unsigned default null,
  "banner" varchar(256) default null,
  "banner_id" bigint unsigned default null,
  "discord_id" varchar(24) default null,
  "twitter_id" varchar(16) default null,
  "telegram_id" varchar(24) default null,
  "github_id" varchar(40) default null,
  "domain" varchar(128) default null,
  "last_changed_block" bigint unsigned default null,
  "created_block" bigint unsigned default null,
  "purchased_block" bigint unsigned default null,
  PRIMARY KEY ("id"),
  UNIQUE KEY "asa_id" ("asa_id"),
  UNIQUE KEY "name" ("name"),
  KEY "state" ("state"),
  KEY "discord_id" ("discord_id"),
  KEY "twitter_id" ("twitter_id"),
  KEY "telegram_id" ("telegram_id")
);

CREATE TABLE "nfd_json" (
  "id" bigint unsigned not null,
  "data" json not null,
  PRIMARY KEY ("id")
);

CREATE TABLE "nfd_wallet" (
  "id" bigint unsigned not null,
  "address" varchar(58) not null,
  "verified" tinyint(1) not null DEFAULT '0',
  "deposit" tinyint(1) not null DEFAULT '0',
  PRIMARY KEY ("id","address")
);

create table "community" (
    "id" bigint unsigned not null,
    "version" varchar(6) not null,
    "akta" bigint unsigned not null default 0,
    "tab" enum('collections', 'staking', 'social', 'raffles', 'shuffles', 'gallery', 'about') not null default 'collections',
    primary key ("id"),
    index ("akta")
);

CREATE TABLE "community_json" (
    "id" bigint unsigned not null,
    "data" json not null,
    "malformed" boolean not null default 0,
    PRIMARY KEY ("id")
);

type CommunityAssociate struct {
	Address   string  `structs:"address,omitempty" db:"address" json:"address,omitempty"`
	Role      string  `structs:"role,omitempty" db:"role" json:"role,omitempty"`
	Confirmed *bool   `structs:"confirmed,omitempty" db:"confirmed" json:"confirmed,omitempty"`
	Txn       *string `structs:"txn,omitempty" db:"txn" json:"txn,omitempty"`
}

create table "community_associate" (
    "id" bigint unsigned not null,
    "address" varchar(58) not null,
    "role" varchar(64) not null,
    "confirmed" boolean not null default 0,
    "txn" varchar(64) default null,
    primary key ("id", "address"),
    index ("address"),
    index ("confirmed")
);


create table "community_token" (
    "id" bigint unsigned not null,
    "asset_id" bigint unsigned not null

    primary key ("id", "asset_id")
);

create table "community_faq" (
    "id" bigint unsigned not null,
    "q" varchar(256) not null,
    "a" text not null,
    "ordering" int unsigned not null,
    primary key ("id", "q")
);

create table "community_extras" (
    "id" bigint unsigned not null,
    "mkey" varchar(128) not null,
    "mvalue" text not null,
    primary key ("id", "mkey"),
    index ("mkey")
);

create table "collection" (
    "id" varchar(24) not null,
    "nfd_id" bigint unsigned not null,
    "name" varchar(128) not null,
    "description" text default null,
    "banner" bigint unsigned,
    "avatar" bigint unsigned,
    "network" varchar(128),
    "explicit" boolean not null default false,
    primary key ("id"),
    index ("nfd_id"),
    unique index ("nfd_id", "name")
);

create table "property" (
    "id" varchar(24) not null,
    "collection_id" varchar(24) not null,
    "name" varchar(128) not null,
    primary key ("id", "name")
);

create table "property_value" (
    "id" varchar(24) not null,
    "name" varchar(128) not null,
    "image" varchar(256) default null,
    "image_integrity" varchar(256),
    "image_mimetype" varchar(32),
    "animation_url" varchar(256),
    "animation_url_integrity" varchar(256),
    "animation_url_mimetype" varchar(32),
    primary key ("id", "name"),
    index ("image")
);

create table "property_value_extras" (
    -- id is the id of the collection the creator is attached to
    "id" varchar(24) not null,
    "name" varchar(128) not null,
    "mkey" varchar(128) not null,
    "mvalue" text not null,
    primary key ("id", "name", "mkey"),
    index ("name"),
    index ("mkey")
);

create table "collection_asset" (
    -- id is the id of the collection the creator is attached to
    "id" varchar(24) not null,
    "asa_id" bigint unsigned not null,
    primary key ("id", "asa_id")
);

create table "collection_excluded_asset" (
    -- id is the id of the collection the creator is attached to
    "id" varchar(24) not null,
    "asa_id" bigint unsigned not null,
    primary key ("id", "asa_id")
);

create table "collection_prefix" (
    -- id is the id of the collection the creator is attached to
    "id" varchar(24) not null,
    -- prefix is the prefix of the collection unit names
    "prefix" varchar(256) not null,
    primary key ("id", "prefix"),
    index ("id"),
    index ("prefix")
);

create table "collection_address" (
    -- id is the id of the collection the creator is attached to
    "id" varchar(24) not null,
    -- address is the algorand public address of a user
    "address" varchar(58) not null,
    primary key ("id", "address"),
    index ("id"),
    index ("address")
);

create table "collection_artist" (
    -- id is the id of the collection the creator is attached to
    "id" varchar(24) not null,
    -- address is the algorand public address of a user
    "address" varchar(58) not null,
    primary key ("id", "address"),
    index ("id"),
    index ("address")
);

create table "collection_extras" (
    -- id is the id of the collection the creator is attached to
    "id" varchar(24) not null,
    "mkey" varchar(128) not null,
    "mvalue" text not null,
    primary key ("id", "mkey"),
    index ("mkey")
);

CREATE TABLE "asset" (
  "id" bigint unsigned NOT NULL,
  "name" varchar(256) DEFAULT NULL,
  "url" varchar(256) DEFAULT NULL,
  "image_url" varchar(256) DEFAULT NULL,
  "standard" enum('asa','nfd','nft-arc69','nft-arc3','nft-arc19','nft-none','lp') DEFAULT NULL,
  "mime" varchar(32) DEFAULT NULL,
  "unit_name" varchar(256) DEFAULT NULL,
  "clawback" varchar(58) DEFAULT NULL,
  "freeze" varchar(58) DEFAULT NULL,
  "manager" varchar(58) DEFAULT NULL,
  "reserve" varchar(58) DEFAULT NULL,
  "creator" varchar(58) NOT NULL,
  "created_at" int unsigned NOT NULL,
  "deleted" tinyint(1) NOT NULL DEFAULT '0',
  "destroyed_at" int unsigned DEFAULT NULL,
  "decimals" bigint unsigned NOT NULL,
  "total" bigint unsigned NOT NULL,

  PRIMARY KEY ("id"),
  KEY "name" ("name"),
  KEY "url" ("url"),
  KEY "standard" ("standard"),
  KEY "unit_name" ("unit_name"),
  KEY "creator" ("creator"),
  KEY "reserve" ("reserve")
);

create table "asset_property" (
    "id" bigint unsigned NOT NULL,
    "meta_key" varchar(128) NOT NULL,
    "meta_value" text NOT NULL,

    primary key ("id", "meta_key"),
    index ("meta_key")
);

-- shf_
create table "shuffle" (
    "id" varchar(24) not null,
    "nfd_id" bigint unsigned not null,
    -- "type" enum("slate", "token") not null default "token",
    "title" varchar(128) not null,
    "about" text NOT NULL,
    "avatar" varchar(256) DEFAULT NULL,
    "avatar_id" int unsigned DEFAULT NULL,
    "banner" varchar(256) DEFAULT NULL,
    "banner_id" int unsigned DEFAULT NULL,
    "max_per_wallet" int unsigned default null,
    "max_sold" int unsigned default null,
    "shuffle_begins_date" datetime NOT NULL,
    "creator" varchar(58) NOT NULL,
    "created_date" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "address" varchar(58) NOT NULL,

    primary key ("id"),
    index ("nfd_id"),
    index ("shuffle_begins_date")
);

-- sas_
create table "shuffle_asset" (
    "shuffle_id" varchar(24) not null,
    "asset_id" bigint unsigned not null,

    primary key ("shuffle_id", "asset_id")
);

-- swl_
create table "shuffle_whitelist_token" (
    "id" varchar(24) not null,
    "shuffle_id" varchar(24) not null,
    "asset_id" bigint unsigned not null,
    -- the number of assets within the whitelist reserve_asset group to keep available
    "reserve_quantity" int unsigned not null default 0,

    primary key ("id", "shuffle_id", "asset_id")
);

create table "shuffle_reserve_asset" (
    "id" varchar(24) not null,
    "shuffle_id" varchar(24) not null,
    -- either address or whitelist_id is required
    "address" varchar(58),
    "whitelist_id" varchar(24),
    "asset_id" bigint unsigned not null,

    primary key ("id"),
    index ("shuffle_id"),
    index ("asset_id"),
    index ("address"),
    index ("whitelist_id")
);

create table "shuffle_payment_option" (
    "id" varchar(24) not null,
    "shuffle_id" varchar(24) not null,
    "whitelist_id" varchar(24) default null,
    "asset_id" bigint unsigned not null,
    -- amount or amount in algo is required
    "amount" bigint unsigned,
    "amount_in_algo" bigint unsigned,

    primary key ("id"),
    index ("shuffle_id"),
    index ("whitelist_id")
);

create table "shuffle_sale" (
    "txn" varchar(64) not null,
    "shuffle_id" varchar(24) not null,
    "asset_id" bigint unsigned not null,

    primary key ("txn", "shuffle_id", "asset_id"),
    index ("shuffle_id")
);


[x] insert into shuffle (id, nfd_id, title, about, avatar, banner, max_sold, shuffle_begins_date, creator, address) values ('shf_cik3g7qs1s440l1db4qg', 765902356, 'Akita Omnigems', 'Akita Omnigems is a second generation series of 5,000 unique NFTs minted on the unstoppable Algorand blockchain.', 'https://cdn.akita.community/shuffles/shf_cik3g7qs1s440l1db4qg/avatar.gif', '', 2750, '2023-07-08: 09:00:00', 'AKCTRDK4OWNWHTPH4XPKLNWNLZ333VE35SKQ4FGQK3ZJA4FIHCLTRG3PFI', 'TGRQ37EJYLS6Z7TQM5QTDJNIQHXRQ7IDBKDUTFL6EDGHSVA5UQAHSRHE3Q');

[x] insert into shuffle_payment_option (id, shuffle_id, asset_id, amount) values ('spo_cik7i12s1s45hkido6tg', 'shf_cik3g7qs1s440l1db4qg', 0, 100);
[x] insert into shuffle_payment_option (id, shuffle_id, asset_id, amount_in_algo) values ('spo_cik7iaas1s45hqhkr840', 'shf_cik3g7qs1s440l1db4qg', 793124631, 100);
[x] insert into shuffle_payment_option (id, shuffle_id, asset_id, amount_in_algo) values ('spo_cik7idqs1s45i7fk4bo0', 'shf_cik3g7qs1s440l1db4qg', 31566704, 100);

-- all 250 - 5000 assets
-- [-] insert into shuffle_asset (shuffle_id, asset_id) values ('shf_cik3g7qs1s440l1db4qg', );

[x] insert into shuffle_reserve_asset (id, shuffle_id, address, asset_id) values ('sra_cik97a2s1s474fu3aa1g', 'shf_cik3g7qs1s440l1db4qg', 'X5KD3VTDUN3H2FFYJX2NDIGQ2V2SS44XUSYGJWHI4KH6ZOXF44SE7KEXVU', 1140534100);
[x] insert into shuffle_reserve_asset (id, shuffle_id, address, asset_id) values ('sra_cik97bas1s474ki60g4g', 'shf_cik3g7qs1s440l1db4qg', 'X5KD3VTDUN3H2FFYJX2NDIGQ2V2SS44XUSYGJWHI4KH6ZOXF44SE7KEXVU', 1140774299);
[x] insert into shuffle_reserve_asset (id, shuffle_id, address, asset_id) values ('sra_cik97cqs1s474pgmboj0', 'shf_cik3g7qs1s440l1db4qg', 'X5KD3VTDUN3H2FFYJX2NDIGQ2V2SS44XUSYGJWHI4KH6ZOXF44SE7KEXVU', 1140770623);

[x] insert into shuffle_whitelist_token (id, shuffle_id, asset_id, reserve_quantity) values ('swl_cik56o2s1s44bhlkmm10', 'shf_cik3g7qs1s440l1db4qg', 955393925, 152); -- regular

[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount) values ('spo_cik7j1qs1s45iduku5j0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik56o2s1s44bhlkmm10', 0, 89); -- algo
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7j6is1s45iip9m34g', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik56o2s1s44bhlkmm10', 793124631, 89); -- governance algo
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7j7qs1s45inarl16g', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik56o2s1s44bhlkmm10', 31566704, 89); -- usdc
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7j8qs1s45irsfuqb0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik56o2s1s44bhlkmm10', 523683256, 89); -- Akita

-- [-] insert into shuffle_reserve_asset (shuffle_id, whitelist_id, asset_id) values ('shf_cik3g7qs1s440l1db4qg', 'swl_cik56o2s1s44bhlkmm10,', xyz);


[x] insert into shuffle_whitelist_token (id, shuffle_id, asset_id, reserve_quantity) values ('swl_cik57ois1s44bri26bhg', 'shf_cik3g7qs1s440l1db4qg', 955399472, 76); -- emerald

[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount) values ('spo_cik7jcqs1s45j0g9p4d0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik57ois1s44bri26bhg', 0, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7jgas1s45j5irf760', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik57ois1s44bri26bhg', 793124631, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7jhas1s45j9g75bu0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik57ois1s44bri26bhg', 31566704, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7jiis1s45je688ivg', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik57ois1s44bri26bhg', 523683256, 89);

-- [-] insert into shuffle_reserve_asset (shuffle_id, whitelist_id, asset_id) values ('shf_cik3g7qs1s440l1db4qg', 'swl_cik57ois1s44bri26bhg,', xyz);


[x] insert into shuffle_whitelist_token (id, shuffle_id, asset_id, reserve_quantity) values ('swl_cik58bqs1s44c3jhejgg', 'shf_cik3g7qs1s440l1db4qg', 955400972, 69); -- ruby

[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount) values ('spo_cik7jqas1s45jkhedip0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58bqs1s44c3jhejgg', 0, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7jrqs1s45jp1ii5n0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58bqs1s44c3jhejgg', 793124631, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7jsqs1s45jtn9m18g', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58bqs1s44c3jhejgg', 31566704, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7ju2s1s45k2h55idg', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58bqs1s44c3jhejgg', 523683256, 89);

-- [-] insert into shuffle_reserve_asset (shuffle_id, whitelist_id, asset_id) values ('shf_cik3g7qs1s440l1db4qg', 'swl_cik58bqs1s44c3jhejgg,', xyz);


[x] insert into shuffle_whitelist_token (id, shuffle_id, asset_id, reserve_quantity) values ('swl_cik58k2s1s44c8jh6q60', 'shf_cik3g7qs1s440l1db4qg', 955403494, 65); -- sapphire

[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount) values ('spo_cik7kgqs1s45ka7ggm50', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58k2s1s44c8jh6q60', 0, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7khqs1s45keumosa0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58k2s1s44c8jh6q60', 793124631, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7kj2s1s45kjvs99n0', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58k2s1s44c8jh6q60', 31566704, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7kkas1s45kp0q0hkg', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58k2s1s44c8jh6q60', 523683256, 89);

-- [-] insert into shuffle_reserve_asset (shuffle_id, whitelist_id, asset_id) values ('shf_cik3g7qs1s440l1db4qg', 'swl_cik58k2s1s44c8jh6q60,', xyz);



[x] insert into shuffle_whitelist_token (id, shuffle_id, asset_id, reserve_quantity) values ('swl_cik58pis1s44cd6e79og', 'shf_cik3g7qs1s440l1db4qg', 955404324, 61); -- topaz

[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount) values ('spo_cik7kpis1s45kv9kj320', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58pis1s44cd6e79og', 0, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7ksis1s45l4b57esg', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58pis1s44cd6e79og', 793124631, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7ktqs1s45l9aqhepg', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58pis1s44cd6e79og', 31566704, 89);
[x] insert into shuffle_payment_option (id, shuffle_id, whitelist_id, asset_id, amount_in_algo) values ('spo_cik7kuqs1s45lduc3mpg', 'shf_cik3g7qs1s440l1db4qg', 'swl_cik58pis1s44cd6e79og', 523683256, 89);

-- [-] insert into shuffle_reserve_asset (shuffle_id, whitelist_id, asset_id) values ('shf_cik3g7qs1s440l1db4qg', 'swl_cik58pis1s44cd6e79og,', xyz);

create table "shuffle_asset_property" (
    "id" bigint unsigned NOT NULL,
    "meta_key" varchar(128) NOT NULL,
    "meta_value" text NOT NULL,

    primary key ("id", "meta_key"),
    index ("meta_key")
);


insert into server_role (id, server_id, role_id, name) values ('rol_ck316s7skmp00gsinfdg', 904883897224032256, 1152707761390485545, 'Bronze');
insert into server_role (id, server_id, role_id, name) values ('rol_ck316vfskmp00prfohb0', 904883897224032256, 1152708756438798456, 'Silver');
insert into server_role (id, server_id, role_id, name) values ('rol_ck3172nskmp013ma2pi0', 904883897224032256, 1152709087168041010, 'Gold');
insert into server_role (id, server_id, role_id, name) values ('rol_ck3174fskmp01ct5gr5g', 904883897224032256, 1152710715287818372, 'Platinum');
insert into server_role (id, server_id, role_id, name) values ('rol_ck31767skmp01lr7tsv0', 904883897224032256, 1152710874046410874, 'Diamond');


insert into role_req (role_id, group_id, greater_than, less_than) values ('rol_ck316s7skmp00gsinfdg', 'grp_cipq34qs1s485123oebg',       999999,          100000000);
insert into role_req (role_id, group_id, greater_than, less_than) values ('rol_ck316vfskmp00prfohb0', 'grp_cipq34qs1s485123oebg',     99999999,         1000000000);
insert into role_req (role_id, group_id, greater_than, less_than) values ('rol_ck3172nskmp013ma2pi0', 'grp_cipq34qs1s485123oebg',    999999999,       100000000000);
insert into role_req (role_id, group_id, greater_than, less_than) values ('rol_ck3174fskmp01ct5gr5g', 'grp_cipq34qs1s485123oebg',  99999999999,      1000000000000);
insert into role_req (role_id, group_id, greater_than, less_than) values ('rol_ck31767skmp01lr7tsv0', 'grp_cipq34qs1s485123oebg', 999999999999, 999999999999999999);


create table "lp" (
    "id" bigint unsigned not null, -- the id of the lp token
    "type" enum('tinyman_1', 'tinyman_1.1', 'tinyman_2', 'humble', 'pact') not null,
    "asset_one" bigint unsigned not null, -- the id of the first asset
    "asset_two" bigint unsigned not null, -- the id of the second asset
    "address" varchar(58) not null, -- the address of pool
    "issued" bigint unsigned not null, -- the amount of lp tokens issued

    primary key ("id"),
    index ("type"),
    index ("asset_one"),
    index ("asset_two"),
    index ("address"),
    index ("issued")
)