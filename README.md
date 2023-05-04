# Getting started

## Installation

```bash
git clone git@github.com:JokerLiAnother/exchange-rate-gather.git

cd exchange-rate-gather

go mod tidy

go run main.go 
```

## Usage

### 1. 获取欧元兑换指定货币的汇率信息，并保存到数据库中

需要创建数据表结构，并编写配置文件`.gather.yaml`。同一汇率信息`(currency_src, currency_dst, valid_date)`只会保存一次，如果已经存在，则不会重复保存。

- 数据库表结构

```sql
-- 创建汇率换算表,用于存储货币与其他货币的汇率关系，并记录有效时间，汇率每天都需要更新
-- 字段：id, currency_src, currency_tag, rate, valid_date
-- 说明：currency_src为货币源，currency_tag为货币目标，rate为汇率，valid_date
-- 例子：id=1, currency_src=USD, currency_tag=CNY, rate=6.5, valid_date=2017-01-01
-- 说明：id为主键，currency_src和currency_tag为联合主键
-- 说明：valid_month为月份，如2017-01，表示2017年1月份的汇率
-- 说明：rate为汇率，如6.5，表示1美元=6.5人民币
-- 说明：valid_month为有效时间，如2017-01，表示2017年1月份的汇率有效
DROP TABLE IF EXISTS config_exchange_rate;
CREATE TABLE config_exchange_rate
(
    id                       BIGINT AUTO_INCREMENT
        PRIMARY KEY COMMENT '使用雪花算法生成',
    currency_src             VARCHAR(10)    NOT NULL COMMENT '源货币',
    currency_dst             VARCHAR(10)    NOT NULL COMMENT '目标货币',
    currency_dst_description VARCHAR(40)    NOT NULL DEFAULT '' COMMENT '目标货币描述',
    rate                     DECIMAL(10, 6) NOT NULL DEFAULT 0.0 COMMENT 'rate为汇率，如6.5，表示1美元=6.5人民币',
    valid_date               VARCHAR(10)    NOT NULL COMMENT 'valid_month为月份，如2017-01-01，表示2017年1月1日的汇率',
    gmt_create               DATETIME       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    gmt_modified             DATETIME NULL ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE KEY `uk_currency_src_dst_date_idx` (`currency_src`, `currency_dst`, `valid_date`)
) COMMENT '汇率换算表,用于存储货币与其他货币的汇率关系，并记录有效时间'
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8;
```

- `.gather.yaml` 配置文件

```yaml
mysql:
  driver: mysql
  url: 'USER:PASSWORD(HOST:PORT)/DATABASE'
  # connection max life time: default 3 minutes
  max-life-time: 3
  # max open connections: default 10
  max-open-connections: 10
  # max idle connections: default 10
  max-idle-connections: 10

# 需要保存汇率的货币
currency-dst: CNY,USD,GBP
```

### 2. 查询欧元当前对其他货币的汇率信息

不需要配置文件，直接运行命令即可

```bash
exchange-rate-gather --currency USD
```
