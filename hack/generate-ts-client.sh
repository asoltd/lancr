#!/bin/bash

# Generate TypeScript client
bun install -g openapi-typescript-codegen

openapi \
  --input ./gen/openapiv2/lancr/v1/hero_service.swagger.json \
  --output ./ts_client \
  --name HeroServiceClient
