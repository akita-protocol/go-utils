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

create table "asset_properties" (
    "id" bigint unsigned NOT NULL,
    "meta_key" varchar(128) NOT NULL,
    "meta_value" text NOT NULL,

    primary key ("id", "meta_key"),
    index ("meta_key")
);