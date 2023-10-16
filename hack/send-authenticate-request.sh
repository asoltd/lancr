#!/bin/bash

TOKEN=$(bun run ./ts_client/get-token.ts)

grpcurl \
	-plaintext \
	-d '{ "id_token": "'$TOKEN'" }' \
	localhost:42069 \
	lancr.v1.AuthService/Authenticate 