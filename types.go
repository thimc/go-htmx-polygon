package main

import "time"

type Stock struct {
	Active          bool      `json:"active"`
	Cik             string    `json:"cik"`
	CompositeFigi   string    `json:"composite_figi"`
	CurrencyName    string    `json:"currency_name"`
	LastUpdatedUtc  time.Time `json:"last_updated_utc"`
	Locale          string    `json:"locale"`
	Market          string    `json:"market"`
	Name            string    `json:"name"`
	PrimaryExchange string    `json:"primary_exchange"`
	ShareClassFigi  string    `json:"share_class_figi"`
	Ticker          string    `json:"ticker"`
}

type SearchResult struct {
	Count     int     `json:"count"`
	NextURL   string  `json:"next_url"`
	RequestID string  `json:"request_id"`
	Results   []Stock `json:"results"`
	Status    string  `json:"status"`
}

type Values struct {
	AfterHours float64 `json:"afterHours"`
	Close      float64 `json:"close"`
	From       string  `json:"from"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	Open       float64 `json:"open"`
	PreMarket  float64 `json:"preMarket"`
	Status     string  `json:"status"`
	Symbol     string  `json:"symbol"`
	Volume     float64   `json:"volume"`
}
