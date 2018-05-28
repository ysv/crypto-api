package main

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

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