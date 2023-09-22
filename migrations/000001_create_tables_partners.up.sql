CREATE TABLE `partners` (
    `id` VARCHAR(255) NOT NULL ,
    `trading_name` TEXT NOT NULL ,
    `document` VARCHAR(150) NOT NULL ,
    `currency` VARCHAR(3) NOT NULL ,
    PRIMARY KEY (`id`), UNIQUE (`document`)
) ENGINE = InnoDB;
