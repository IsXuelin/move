use move;

DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
     `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
     `driver_id` int(11) NOT NULL COMMENT '司机id',
     `start_lat` decimal(10,6) NOT NULL COMMENT  '起点维度',
     `start_lng` decimal(10,6) NOT NULL COMMENT  '起点经度',
     `end_lat` decimal(10,6) NOT NULL COMMENT  '终点维度',
     `end_lng` decimal(10,6) NOT NULL COMMENT  '终点经度',
     `create_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
     `end_time` TIMESTAMP NULL COMMENT '结束时间',
     `status` int(5) COMMENT '状态',
     PRIMARY KEY (`id`),
     INDEX (passenger_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
