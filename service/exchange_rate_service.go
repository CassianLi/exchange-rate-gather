package service

import (
	"exchange-rate-gather/model"
	"exchange-rate-gather/utils"
	"fmt"
	"strings"
)

// ExchangeRateForNlService 创建一个类用来获取NL汇率信息
type ExchangeRateForNlService struct {
	// 年份
	Year string
	// 月份
	Month string
}

// GetExchangeRates 获取汇率
func (e *ExchangeRateForNlService) GetExchangeRates() (rates []model.ExchangeRate, err error) {
	// 获取汇率API
	api := e.getExchangeRateApi()
	// 获取汇率信息
	body, err := utils.HttpGet(api)
	if err != nil {
		return nil, err
	}
	// 解析汇率信息
	rates, err = e.parseExchangeRateFromXmlBody(body)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

// getExchangeRateApi	获取汇率API
func (e *ExchangeRateForNlService) getExchangeRateApi() (api string) {
	return fmt.Sprintf("https://www.belastingdienst.nl/data/douane_wisselkoersen/wks.douane.wisselkoersen.dd%s%s.xml", e.Year, e.Month)
}

// 从xml Body中解析汇率信息
func (e *ExchangeRateForNlService) parseExchangeRateFromXmlBody(body string) (rates []model.ExchangeRate, err error) {
	// str解析为xml对象，并指定encoding为:ISO-8859-1
	var xmlObj model.ExchangeRateXml
	err = utils.XmlToStruct(body, &xmlObj, "ISO-8859-1")
	if err != nil {
		return nil, err
	}
	ratesXml := xmlObj.Rates
	for _, rateXml := range ratesXml {
		rate := rateXml.Rate
		// rate: 1.406,43 调整格式为: 1406.43
		if strings.Contains(rate, ".") {
			rate = strings.ReplaceAll(rate, ".", "")
		}

		if strings.Contains(rate, ",") {
			rate = strings.ReplaceAll(rate, ",", ".")
		}
		exchangeRate := model.ExchangeRate{
			CurrencySrc:            "EUR",
			CurrencyDst:            rateXml.Currency,
			CurrencyDstDescription: rateXml.Description,
			Rate:                   utils.StrToFloat64(rate),
			ValidMonth:             fmt.Sprintf("%s-%s", xmlObj.UpdateInfo.Year, xmlObj.UpdateInfo.Month),
		}
		rates = append(rates, exchangeRate)
	}
	return rates, nil
}
