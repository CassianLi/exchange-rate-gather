/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"exchange-rate-gather/service"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var currency string

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "查询欧元与指定货币的汇率",
	Long: `查询当前欧元与指定货币的汇率. For example:

exchange-rate-gather query --currency=USD`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("query called")

		fmt.Printf("Get EUR to %s's exchage rate ...\n", currency)

		service := service.ExchangeRateForNlService{
			Year:  year,
			Month: month,
		}
		rates, err := service.GetExchangeRates()
		if err != nil {
			log.Println("Gather exchange rates error: ", err)
		}
		log.Println("EUR\tCurrency\tCurrency Description\tRate\tValid Month")
		for _, rate := range rates {
			// 想表格一样打印汇率信息
			log.Printf("%s\t%s\t%s\t%f\t%s\n", rate.CurrencySrc, rate.CurrencyDst, rate.CurrencyDstDescription, rate.Rate, rate.ValidMonth)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//queryCmd.PersistentFlags().StringVar(&currency, "currency", "", "指定货币，如：USD")

	queryCmd.PersistentFlags().StringVar(&currency, "currency", "", "货币,(如：USD, 不指定返回所有货币的汇率)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
