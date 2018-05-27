package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)


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
