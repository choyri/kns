CREATE TABLE IF NOT EXISTS `import_records`
(
    `id`         INT UNSIGNED AUTO_INCREMENT,
    `start_date` DATE,
    `end_date`   DATE,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '导入记录';

CREATE TABLE IF NOT EXISTS `export_records`
(
    `id`         INT UNSIGNED AUTO_INCREMENT,
    `amount`     INT UNSIGNED COMMENT '数量',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '导出记录';

CREATE TABLE IF NOT EXISTS `orders`
(
    `id`                      INT UNSIGNED AUTO_INCREMENT,
    `import_record_id`        INT UNSIGNED NOT NULL COMMENT '关联 import_records.id',
    `customer_name`           VARCHAR(50)  NOT NULL COMMENT '客户名称',
    `salesman`                VARCHAR(50)  NOT NULL COMMENT '业务',
    `customer_order_number`   VARCHAR(50)  NOT NULL COMMENT '客户单号',
    `brand`                   VARCHAR(50)  NOT NULL COMMENT '品牌',
    `order_number`            VARCHAR(50)  NOT NULL COMMENT '订单号码',
    `serial_number`           INT          NOT NULL COMMENT '序号',
    `product_name_code`       VARCHAR(50)  NOT NULL COMMENT '品名代码',
    `product_name_chinese`    VARCHAR(50)  NOT NULL COMMENT '中文品名',
    `product_name_english`    VARCHAR(100) NOT NULL COMMENT '英文品名',
    `ingredient`              VARCHAR(100) NOT NULL COMMENT '成分',
    `specification`           VARCHAR(50)  NOT NULL COMMENT '规格',
    `color`                   VARCHAR(50)  NOT NULL COMMENT '颜色',
    `color_number`            VARCHAR(50) COMMENT '色号',
    `customer_version_number` VARCHAR(50) COMMENT '客户版号',
    `created_at`              TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`              TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FULLTEXT KEY `ft_customer_name` (`customer_name`),
    FULLTEXT KEY `ft_salesman` (`salesman`),
    FULLTEXT KEY `ft_customer_order_number` (`customer_order_number`),
    FULLTEXT KEY `ft_brand` (`brand`),
    FULLTEXT KEY `ft_order_number` (`order_number`),
    FULLTEXT KEY `ft_product_name_code` (`product_name_code`),
    FULLTEXT KEY `ft_product_name_chinese` (`product_name_chinese`),
    FULLTEXT KEY `ft_product_name_english` (`product_name_english`),
    FULLTEXT KEY `ft_ingredient` (`ingredient`),
    FULLTEXT KEY `ft_specification` (`specification`),
    FULLTEXT KEY `ft_color` (`color`),
    FULLTEXT KEY `ft_color_number` (`color_number`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '订单';
