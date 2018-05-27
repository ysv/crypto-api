package main

var routes = Routes{
  Route{
    "CurrenciesIndex",
    "GET",
    "/currencies",
    CurrenciesIndex,
  },
  Route{
    "CurrencyShow",
    "GET",
    "/currencies/{code}",
    CurrencyShow,
  },
}