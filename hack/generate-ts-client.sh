#!/bin/bash

# Generate TypeScript client
bun install -g openapi-typescript-codegen

openapi \
--input ./gen/openapiv2/api.swagger.json \
--output ./ts_client \
--name ApiClient
