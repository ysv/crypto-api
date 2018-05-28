package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "fmt"
)

func setContentTypeJSON(w http.ResponseWriter){
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func CurrenciesIndex(w http.ResponseWriter, r *http.Request){
  setContentTypeJSON(w)
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(currencies)
}

func CurrencyShow(w http.ResponseWriter, r *http.Request) {
  setContentTypeJSON(w)

  vars := mux.Vars(r)
  currencyId := vars["code"]
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(CurrencyFindByCode(currencyId))
}

func AuthCreate(w http.ResponseWriter, r *http.Request)  {
  setContentTypeJSON(w)
  var auth User
  _ = json.NewDecoder(r.Body).Decode(&auth)

}