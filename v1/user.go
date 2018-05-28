package main

import (
  "errors"
)

type UserProfile struct {
  Name     string `json:"name"`
  Password string `json:"password"`
}

type Users []UserProfile

var users Users

func ValidateUser(user UserProfile) error{
  for i := range users {
    if users[i].Name == user.Name && users[i].Password == user.Password {
      return nil
    }
  }
  return errors.New("wrong user or password")
}

func LoadUsers(){
  users = append(users, UserProfile{"yaroslav", "changeme"})
}