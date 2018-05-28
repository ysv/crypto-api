package main

var routes = Routes{
  Route{
    "CurrenciesIndex",
    "POST",
    "/currencies",
    CurrenciesIndex,
  },
  Route{
    "CurrencyShow",
    "POST",
    "/currencies/{code}",
    CurrencyShow,
  },
  Route{
    "UserProfile",
    "POST",
    "/auth/login",
    AuthCreate,
  },
}