package main

import (
  "net/url"
)

type Currency struct {
  Code            string    `json:"code"`
  Symbol          string    `json:"symbol"`
  JSONRPCEndpoint *url.URL  `json:"-"`
}

type Currencies []Currency

var currencies Currencies

// currency.go
func CurrencyFindByCode(code string) Currency{
  for i := range currencies {
    if currencies[i].Code == code {
      return currencies[i]
    }
  }
  return Currency{}
}

// currency.go
func LoadCurrencies()  {
  currencies = append(currencies, Currency{Code: "BTC", Symbol: "B"})
  currencies = append(currencies, Currency{Code: "ETH", Symbol: "E"})
}

