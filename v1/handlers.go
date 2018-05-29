package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type SendToAddressParams struct {
  Address string  `json:"address"`
  Amount  float32 `json:"amount"`
}

func setContentTypeJSON(w http.ResponseWriter){
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request) (error){
  token := struct {
    Token string `json:"token"`
  }{}
  _ = json.NewDecoder(r.Body).Decode(&token)
  _, err := ValidateJWT(token.Token)
  if err != nil{
    w.WriteHeader(http.StatusForbidden)
    return err
  }
  return nil
}

func AuthCreate(w http.ResponseWriter, r *http.Request)  {
  setContentTypeJSON(w)
  var user UserProfile
  _ = json.NewDecoder(r.Body).Decode(&user)
  if err := ValidateUser(user); err != nil{
    w.WriteHeader(http.StatusForbidden)
  } else {
    w.WriteHeader(http.StatusOK)
    var sessionJWT = GenerateSessionJWT(user)
    json.NewEncoder(w).Encode(struct {
      Token string `json:"token"`
    }{sessionJWT})
  }
}

func CurrenciesIndex(w http.ResponseWriter, r *http.Request){
  setContentTypeJSON(w)
  err := AuthMiddleware(w, r)
  if err != nil{
    return
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(currencies)
}

func CurrencyShow(w http.ResponseWriter, r *http.Request) {
  setContentTypeJSON(w)
  err := AuthMiddleware(w, r)
  if err != nil{
    return
  }
  vars := mux.Vars(r)
  currencyId := vars["code"]
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(CurrencyFindByCode(currencyId))
}

func CurrencyGetNewAddress(w http.ResponseWriter, r *http.Request){
  setContentTypeJSON(w)
  err := AuthMiddleware(w, r)
  if err != nil{
    return
  }
  vars := mux.Vars(r)
  currencyId := vars["code"]
  currency := CurrencyFindByCode(currencyId)
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(GetNewAddress(currency))
}


func CurrencyGetBalance(w http.ResponseWriter, r *http.Request){
  setContentTypeJSON(w)
  err := AuthMiddleware(w, r)
  if err != nil{
    return
  }
  vars := mux.Vars(r)
  currencyId := vars["code"]
  currency := CurrencyFindByCode(currencyId)
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(GetBalance(currency))
}

func CurrencySendToAddress(w http.ResponseWriter, r *http.Request){
  setContentTypeJSON(w)
  err := AuthMiddleware(w, r)
  if err != nil{
    return
  }
  vars := mux.Vars(r)
  currencyId := vars["code"]
  currency := CurrencyFindByCode(currencyId)
  params := SendToAddressParams{}
  _ = json.NewDecoder(r.Body).Decode(&params)
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(SendToAddress(currency, params.Address, params.Amount))
}