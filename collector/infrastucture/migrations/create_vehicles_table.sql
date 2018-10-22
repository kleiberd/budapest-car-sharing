-- +migrate Up
CREATE TABLE vehicles (
	`external_id` VARCHAR (255) NOT NULL,
	`provider` VARCHAR (255) NOT NULL,
	`latitude` DECIMAL(14, 12) NOT NULL,
	`longitude` DECIMAL(15, 12) NOT NULL,
	`type` VARCHAR (255) NOT NULL,
	`fuel_type` VARCHAR (255) NOT NULL,
	`brand` VARCHAR (255) NOT NULL,
	`model` VARCHAR (255) NOT NULL,
	`plate` VARCHAR (255) NOT NULL,
	`range` INT(10) NOT NULL,
	`created_at` DATETIME NOT NULL,
	PRIMARY KEY (`external_id`, `provider`)
);

-- +migrate Down
DROP TABLE vehicles;