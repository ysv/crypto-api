package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type Token struct {
  Token string `json:"token"`
  Address string  `json:"address"`
  Amount string `json:"amount"`
}

func setContentTypeJSON(w http.ResponseWriter){
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request) (Token,error){
  token := Token{}
  _ = json.NewDecoder(r.Body).Decode(&token)
  _, err := ValidateJWT(token.Token)
  if err != nil{
    w.WriteHeader(http.StatusForbidden)
    return Token{},err
  }
  return token, nil
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

func CurrencyGetNewAddress(w http.ResponseWriter, r *http.Request){
  setContentTypeJSON(w)
  _, err := AuthMiddleware(w, r)
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
  _, err := AuthMiddleware(w, r)
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
  params, err := AuthMiddleware(w, r)
  if err != nil{
    return
  }
  vars := mux.Vars(r)
  currencyId := vars["code"]
  currency := CurrencyFindByCode(currencyId)
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(SendToAddress(currency, params.Address, params.Amount))
}