# Backend

There is an issue with the `github.com/infobloxopen/protoc-gen-gorm` package,
it uses deprecated version of `gorm.io/gorm` and there is a nil pointer ref
, using my patch under `github.com/piotrostr/protoc-gen-gorm`

In order to generate GORM models from the protobuf definition, the
`github.com/infobloxopen/protoc-gen-gorm` package is used with the dependency
protobufs hosted on `buf.build/piotrostr/protoc-gen-gorm`, the installation is

```sh
go install github.com/infobloxopen/protoc-gen-gorm@latest
```

but as of `v1.1.2` requires a change in one of the files:
`github.com/infobloxopen/protoc-gen-gorm/plugin/plugin.go:2830`

```go
if field == nil || field.Desc == nil || field.Desc.Message() == nil {
    return false, "", ""
}
```

Backend will be migrating to PostreSQL to start with and then Spanner using the
PostgreSQL adapter. For the PoC will be using a single PostgreSQL service in
the cluster or even in-memory, the important part is that you can generate GORM
models based on .proto definitions

No vendor lock-in, compatible with any database system

And is secure against injections of any sort

Firestore is cool for a few reasons:

- fast, simple to use and works in real-time (updates to documents are visible
  to the user instantly)
- noSQL

Firestore sucks for a number of reasons:

- NoSQL
- not a lot of querying possibilties, search sucks on firestore - you can use
  Typesense plugins etc but this costs and is not very configurable
- not suitable for backend usage, is a frontend serverless database with authentication rules
- costly: each read costs, multiple reads required for nested collections
- mappings are not supported, the relationships are hard to implement
- if someone changes something all of the objects storing that thing have a hard
  time, e.g. changing a username or a team name would require performing this
  change for all of the documents referencing the data

Currently, authentication is trivial, the token is required to POST/PUT/DELETE,
but the ownerID is not checked against the clientID

All of the scripts required to set up the infrastructure except for service
account creation and applying the required permissions can be found in `./hack`

The CI/CD uses Github Actions and Google Cloud Workload Identity to run tests,
builds, push and deploy to a GKE cluster.

Service accounts required are the `firestore-editor@` and `cloud-deploy-sa@`, one
is for running tests, the other for building and pushing (Cloud Build) and
creating releases (Cloud Deploy)

The `.github/workflows/*.yaml` files are based around `./hack/build.sh`,
`./hack/test.sh`, `./hack/coverage.sh` and `./hack/deploy.sh` which can be used
manually.

The repo also works well with `skaffold`, if you `skaffold run` with the
`kubectl` authenticated, you will roll out a deployment.

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
