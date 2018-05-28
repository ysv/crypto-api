package main

import (
  "net/http"
  "bytes"
  "fmt"
  "encoding/json"
)

type RPCResponse struct {
  Result string `json:"result"`
}

func rpcCall(method string) *http.Response{
  var jsonStr = []byte(`{"method":"` + method + `"}`)
  req, _ := http.NewRequest("POST", "http://yaroslav:changeme@127.0.0.1:18332", bytes.NewBuffer(jsonStr))
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  return resp
}

func getNewAddress() RPCResponse{
  response := rpcCall("getnewaddress")
  defer response.Body.Close()
  rpcResponse := RPCResponse{}
  //body, _ := ioutil.ReadAll(response.Body)
  json.NewDecoder(response.Body).Decode(&rpcResponse)
  //fmt.Println("response Body2:", string(body))
  fmt.Println("genNewAddress: ", rpcResponse.Result)
  return rpcResponse
}

func getBlockchainInfo()  RPCResponse{
  response := rpcCall("getbestblockhash")
  defer response.Body.Close()
  rpcResponse := RPCResponse{}
  //body, _ := ioutil.ReadAll(response.Body)
  json.NewDecoder(response.Body).Decode(&rpcResponse)
  //fmt.Println("response Body2:", string(body))
  fmt.Println("getbestblockhash: ", rpcResponse.Result)
  return rpcResponse
}

func main()  {
  res, _ := json.Marshal(getNewAddress())
  fmt.Println(string(res))

  res, _ = json.Marshal(getBlockchainInfo())
  fmt.Println(string(res))
}