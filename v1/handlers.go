package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

func CurrenciesIndex(w http.ResponseWriter, r *http.Request){
  json.NewEncoder(w).Encode(currencies)
}

func CurrencyShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  currencyId := vars["code"]
  json.NewEncoder(w).Encode(CurrencyFindByCode(currencyId))
}

