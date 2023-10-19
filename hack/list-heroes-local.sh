curl $(minikube ip)/v1/heroes \
  -H "Authorization: Bearer $(bun run ./ts_client/get-token.ts)" \
  -H "accept: application/json"

