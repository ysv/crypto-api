package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "net/url"
  "encoding/json"
)

type Currency struct {
  Code            string    `json:"code"`
  Symbol          string    `json:"symbol"`
  JSONRPCEndpoint *url.URL  `json:"-"`
}

type Currencies []Currency

var currencies Currencies

func main() {

  // Init router.
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index).Methods("GET")
  router.HandleFunc("/currencies", CurrenciesIndex).Methods("GET")
  router.HandleFunc("/currencies/{code}", CurrencyShow).Methods("GET")

  // Load currencies.
  LoadCurrencies()

  // Start server.
  log.Fatal(http.ListenAndServe(":8080", router))
}

// handlers.go
func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// handlers.go
func CurrenciesIndex(w http.ResponseWriter, r *http.Request){
  json.NewEncoder(w).Encode(currencies)
}

// handlers.go
func CurrencyShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  currencyId := vars["code"]
  json.NewEncoder(w).Encode(CurrencyFindByCode(currencyId))
}

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