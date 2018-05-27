package main

var routes = Routes{
  Route{
    "Index",
    "GET",
    "/",
    Index,
  },
  Route{
    "TodoIndex",
    "GET",
    "/currencies",
    CurrenciesIndex,
  },
  Route{
    "TodoShow",
    "GET",
    "/currencies/{code}",
    CurrencyShow,
  },
}