package model

const (
	// QueryExchangeRateExists 查询汇率是否存在
	QueryExchangeRateExists = `select count(*) from config_exchange_rate where currency_src = ? and currency_dst = ? and valid_month = ?`

	// InsertExchangeRate 插入汇率
	InsertExchangeRate = `insert into config_exchange_rate(currency_src, currency_dst, currency_dst_description, rate, valid_month) 
values(:currency_src, :currency_dst, :currency_dst_description, :rate, :valid_month)`
)
