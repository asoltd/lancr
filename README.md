## Backend

In order to deploy, use `./hack/deploy.sh` script

Development of the backend requires buf.build and some plugins

```sh
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

To update stubs:

```sh
buf mod update  # initially
buf generate
go mod tidy
```

After any updates, be sure to

```sh
SWAGGER_UI_VERSION=v5.0.0 ./hack/generate-swagger-ui.sh
```

to regenerate Swagger UI docs

## License

MIT
