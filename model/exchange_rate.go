package model

import "encoding/xml"

// ExchangeRate 汇率
type (
	ExchangeRate struct {
		// ID 主键
		ID int64 `db:"id" json:"id"`
		// CurrencySrc 源货币
		CurrencySrc string `db:"currency_src" json:"currencySrc"`

		// CurrencyDst 目标货币
		CurrencyDst string `db:"currency_dst" json:"currencyDst"`

		// CurrencyDstDescription 目标货币描述
		CurrencyDstDescription string `db:"currency_dst_description" json:"currencyDstDescription"`

		// Rate 汇率
		Rate float64 `db:"rate" json:"rate"`
		// ValidMonth 有效日期
		ValidMonth string `db:"valid_month" json:"validMonth"`
		// GmtCreate 创建时间
		GmtCreate string `db:"gmt_create" json:"gmtCreate"`
		// GmtModified 修改时间
		GmtModified string `db:"gmt_modified" json:"gmtModified"`
	}

	ExchangeRateXml struct {
		XmlName    xml.Name  `xml:"publicerenKoersen"`
		UpdateInfo UpdateXml `xml:"update"`
		Rates      []RateXml `xml:"douaneMaandwisselkoers"`
	}

	UpdateXml struct {
		Year  string `xml:"jaar"`
		Month string `xml:"maand"`
	}

	RateXml struct {
		Currency    string `xml:"muntCode"`
		Description string `xml:"muntSoort"`
		Rate        string `xml:"tariefInVreemdeValuta"`
	}
)
