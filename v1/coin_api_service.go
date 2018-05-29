package main

import (
  "net/http"
  "bytes"
  "fmt"
  "encoding/json"
  "strings"
  "strconv"
)

type CoinAPIService interface {
  rpcCall(method string, optParams ...string) *http.Response
  getNewAddress() RPCResponseStringRes
  getBalance() RPCResponseFloatRes
  sendToAddress(address string, amount int)  RPCResponseStringRes
}

type BTC struct{
  Currency
}

type RPCResponseStringRes struct {
  Result string `json:"result"`
}

type RPCResponseFloatRes struct {
  Result float32 `json:"result"`
}

func GetNewAddress(currency Currency) RPCResponseStringRes{
  if currency.Code == "BTC" {
    service := BTC{currency}
    return service.getNewAddress()
  }
  return RPCResponseStringRes{}
}

func GetBalance(currency Currency) RPCResponseFloatRes {
  if currency.Code == "BTC" {
    service := BTC{currency}
    return service.getBalance()
  }
  return RPCResponseFloatRes{}
}

func SendToAddress(currency Currency, address string, amount int) RPCResponseStringRes{
  if currency.Code == "BTC" {
    service := BTC{currency}
    return service.sendToAddress(address, amount)
  }
  return RPCResponseStringRes{}
}

func (btc *BTC) rpcCall(method string, optParams ...string) *http.Response{
  for i, el := range optParams {
    optParams[i] = strconv.Quote(el)
  }
  optParamsString := strings.Join(optParams, ",")
  jsonParamsSlice := []string{ `"method":"` + method + `"`}
  if len(optParamsString) > 0 {
    jsonParamsSlice = append(jsonParamsSlice, `"params":[` + strings.Join(optParams, ",") + `]`)
  }
  var jsonStr = []byte("{" + strings.Join(jsonParamsSlice, ",") + "}")
  fmt.Println("{" + strings.Join(jsonParamsSlice, ",") + "}")
  req, _ := http.NewRequest("POST", btc.JSONRPCEndpoint.String(), bytes.NewBuffer(jsonStr))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  return resp
}

func (btc *BTC) getNewAddress() RPCResponseStringRes{
  response := btc.rpcCall("getnewaddress")
  defer response.Body.Close()
  rpcResponse := RPCResponseStringRes{}
  json.NewDecoder(response.Body).Decode(&rpcResponse)
  fmt.Println("genNewAddress: ", rpcResponse.Result)
  return rpcResponse
}

func (btc *BTC) getBalance() RPCResponseFloatRes {
  response := btc.rpcCall("getbalance")
  defer response.Body.Close()
  rpcResponse := RPCResponseFloatRes{}
  json.NewDecoder(response.Body).Decode(&rpcResponse)
  fmt.Println("getbalance: ", rpcResponse.Result)
  return rpcResponse
}

func (btc *BTC) sendToAddress(address string, amount int)  RPCResponseStringRes{
  response := btc.rpcCall("sendtoaddress", address, strconv.Itoa(amount))
  defer response.Body.Close()
  rpcResponse := RPCResponseStringRes{}
  json.NewDecoder(response.Body).Decode(&rpcResponse)
  fmt.Println("sendtoaddress: ", rpcResponse.Result)
  return rpcResponse
}

//func main()  {
//
//  btcUrl, _ := url.Parse("http://yaroslav:changeme@127.0.0.1:18332")
//  fmt.Println(btcUrl.String())
//  btc := &BTC{ Currency{"BTC", "B", btcUrl}}
//
//  res, _ := json.Marshal(btc.getNewAddress())
//  fmt.Println(string(res))
//
//  res, _ = json.Marshal(btc.getBalance())
//  fmt.Println(string(res))
//
//  res, _ = json.Marshal(btc.sendToAddress("2MtKeBvWttU36TfJARKcPbZgoLQ7KXwB7fT", 10))
//  fmt.Println(string(res))
//}