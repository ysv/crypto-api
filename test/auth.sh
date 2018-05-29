set -xe
curl -X POST -H "Content-Type: application/json" -d '{"name":"yaroslav","password":"changeme"}' localhost:8080/auth/login > auth.json && cat auth.json | jq
