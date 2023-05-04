/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"exchange-rate-gather/service"
	"fmt"
	"github.com/spf13/cobra"
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

		rates, err := service.GatherExchangeRates()
		if err != nil {
			fmt.Println("GatherExchangeRates error:", err)
			return
		}

		if currency == "" {
			first, _ := json.Marshal(rates[0])
			fmt.Println("Currency did not input, default print the first rate info: ", string(first))
		} else {
			for _, rate := range rates {
				if rate.CurrencyDst == currency {
					fmt.Printf("EUR to %s's current exchange rate = %f \n", currency, rate.Rate)
					break
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	queryCmd.PersistentFlags().StringVar(&currency, "currency", "", "指定货币，如：USD")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
