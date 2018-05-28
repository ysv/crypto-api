package main

import (
  "net/url"
  "strings"
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
    if strings.ToLower(currencies[i].Code) == strings.ToLower(code) {
      return currencies[i]
    }
  }
  return Currency{}
}

// currency.go
func LoadCurrencies()  {
  btcUrl, _ := url.Parse("http://yaroslav:changeme@127.0.0.1:18332")
  currencies = append(currencies, Currency{Code: "BTC", Symbol: "B", JSONRPCEndpoint: btcUrl})})
  currencies = append(currencies, Currency{Code: "ETH", Symbol: "E"})
}

