package model

const (
	// InsertExchangeRate 插入汇率
	InsertExchangeRate = `insert into config_exchange_rate(currency_src, currency_dst, currency_dst_description, rate, valid_date) 
values(:currency_src, :currency_dst, :currency_dst_description, :rate, :valid_date)`
)
