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

func FindUser(user UserProfile) (UserProfile, error) {
  for i := range users {
    if users[i].Name == user.Name && users[i].Password == user.Password {
      return users[i], nil
    }
  }
  return UserProfile{}, errors.New("user doesn't exist")
}

func ValidateUser(user UserProfile) error{
  _, err := FindUser(user)
  return err
}

func LoadUsers(){
  users = append(users, UserProfile{"yaroslav", "changeme"})
}