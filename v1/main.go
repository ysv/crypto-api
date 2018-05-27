package main

import (
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "net/url"
)

type Currency struct {
  Code            string `json:"code"`
  Symbol          string `json:"symbol"`
  JSONRPCEndpoint *url.URL
}

func main() {

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index).Methods("GET")
  router.HandleFunc("/currencies", CurrenciesIndex).Methods("GET")
  router.HandleFunc("/currencies/{code}", CurrencyShow).Methods("GET")

  log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func CurrenciesIndex(w http.ResponseWriter, r *http.Request){
  fmt.Fprintln(w, "Currency Index")
}

func CurrencyShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  currencyId := vars["code"]
  fmt.Fprintln(w, "Currency Show: ", currencyId)
}