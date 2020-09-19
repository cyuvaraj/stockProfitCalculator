package main

import (
	"context"
	"flag"
)

func main() {

	var (
		inputFilePath       string
		skipCorporateAction bool
		startFinancialMonth string
		endFinancialMonth   string
	)
	{
		flag.StringVar(&inputFilePath, "input-csv", "TradeHistory.csv", "csv input file path")
		flag.BoolVar(&skipCorporateAction, "skip-corp-action", true, "skip corporate action")
		flag.StringVar(&startFinancialMonth, "start-financial-month", "july", "month when the financial year starts")
		flag.StringVar(&endFinancialMonth, "end-financial-month", "jun", "month when the financial year ends")

	}

	flag.Parse()

	ctx := context.Background()

	config := Config{
		InputFilePath:       inputFilePath,
		SkipCorporateAction: skipCorporateAction,
	}
	calculateProfits(ctx, config)

}
