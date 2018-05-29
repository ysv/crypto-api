set -xe
curl -X GET -H "Content-Type: application/json" localhost:8080/currencies/btc | jq
