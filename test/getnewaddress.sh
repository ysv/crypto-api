set -xe
curl -X POST -H "Content-Type: application/json" -d "@auth.json" localhost:8080/currencies/btc/getnewaddress | jq
