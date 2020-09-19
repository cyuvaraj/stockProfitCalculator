package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func calculateProfits(ctx context.Context, config Config) {
	transactionData := loadCsvFile(config.InputFilePath)
	transactions := mapToStruct(transactionData)
	getRemainingCount(transactions, config)

}

func getRemainingCount(transactions []Transaction, config Config) {

	shares := make(map[string]uint)

	for _, t := range transactions {

		if config.SkipCorporateAction {
			if strings.ToLower(t.Activity) != "trade" {
				continue
			}
		}

		var currentCount uint = 0
		if count, ok := shares[t.Market]; ok {
			currentCount = count
		}

		if t.Direction == "BUY" {
			currentCount += uint(Abs(t.Quantity))
		} else {
			currentCount -= uint(Abs(t.Quantity))
		}

		shares[t.Market] = currentCount
	}

	for k, s := range shares {
		fmt.Printf("%s - %d \n", k, s)
	}

}

func Abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func mapToStruct(transactionData map[int]map[string]string) []Transaction {

	transactions := make([]Transaction, 0)

	for _, t := range transactionData {

		transaction := Transaction{}

		//transaction.Date = getDate(t["date"])
		transaction.Market = t["market"]
		transaction.Cost = getFloat(t["cost/proceeds"])
		transaction.Direction = t["direction"]
		transaction.Price = getFloat(t["price"])
		transaction.Activity = t["activity"]

		transaction.Quantity = getInt(t["quantity"])

		transactions = append(transactions, transaction)
	}

	return transactions
}

func getInt(data string) int {
	value, err := strconv.ParseInt(data, 10, 0)

	if err != nil {
		panic(err)
	}
	return int(value)
}

func getFloat(data string) float32 {
	value, err := strconv.ParseFloat(data, 32)

	if err != nil {
		panic(err)
	}
	return float32(value)
}

func getDate(data string) time.Time {
	t, err := time.Parse(data, data)

	if err != nil {
		panic(err)
	}

	return t
}
