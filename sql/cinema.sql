CREATE TABLE cinema (
    cinema_id INT(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
    cinema_name CHAR(100) NOT NULL DEFAULT '',
    cinema_add CHAR(150) NOT NULL DEFAULT '',
    location_id INT(11) NOT NULL DEFAULT 0,
    cinema_types CHAR(100) NOT NULL DEFAULT '',
    cinema_card INT(11) NOT NULL DEFAULT 0 COMMENT '影城卡',
    cinema_min_price INT(11) NOT NULL DEFAULT 0 COMMENT '几元起',
    cinema_support CHAR(200) NOT NULL DEFAULT '' COMMENT '影院提供的支持，包括退签等',
    cinema_add_id INT(11) NOT NULL DEFAULT 0 COMMENT '影院城市对应的位置',
    cinema_discount INT(11) NOT NULL DEFAULT 0 COMMENT '影城卡最低减价多少元',
    PRIMARY KEY(cinema_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
