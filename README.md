## Backend

All of the scripts required to set up the infrastructure except for service account creation and applying the required permissions can be found in `./hack`

The CI/CD uses Github Actions and Google Cloud Workload Identity to run tests,
builds, push and deploy to a GKE cluster.

Service accounts required are the `firestore-editor@` and `cloud-deploy-sa@`, one
is for running tests, the other for building and pushing (Cloud Build) and
creating releases (Cloud Deploy)

The `.github/workflows/*.yaml` files are based around `./hack/build.sh`, `./hack/test.sh`, `./hack/coverage.sh` and `./hack/deploy.sh` which can be used manually.

The repo also works well with `skaffold`, if you `skaffold run` with the `kubectl` authenticated, you will roll out a deployment.

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

The Swagger docs are dependant on an external repository, it was cloned with

```sh
SWAGGER_UI_VERSION=v5.0.0 ./hack/generate-swagger-ui.sh
```

to regenerate Swagger UI docs

## License

MIT
