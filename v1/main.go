package main

import (
  "log"
  "net/http"
)


func main() {
  // Initialize router.
  router := NewRouter()

  // Load currencies.
  LoadCurrencies()

  // Start server.
  log.Fatal(http.ListenAndServe(":8080", router))
}
