package main

import "time"

type Config struct {
	InputFilePath       string
	SkipCorporateAction bool
}

type Transaction struct {
	Market    string
	Direction string
	Cost      float32
	Price     float32
	Quantity  int
	Date      time.Time
	Activity  string
}
