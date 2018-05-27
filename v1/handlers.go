package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "fmt"
  "html"
)

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

