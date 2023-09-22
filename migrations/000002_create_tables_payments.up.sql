CREATE TABLE `payments` (
    `id` VARCHAR(28) NOT NULL ,
    `partner_id` VARCHAR(255) NOT NULL ,
    `amount` FLOAT NOT NULL ,
    `foreign_amount` FLOAT NOT NULL ,
    `consumer_name` VARCHAR(255) NOT NULL ,
    `consumer_national_id` VARCHAR(13) NOT NULL ,
    `created_at` DATETIME NOT NULL ,
    PRIMARY KEY (`id`),
    INDEX (`partner_id`)
) ENGINE = InnoDB;
