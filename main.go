/*
Copyright © 2023 piotr.jp.ostrowski@gmail.com
*/
package main

import "github.com/asoltd/lancr/cmd"

// TODO's
// - [ ] Add TLS certs for gateway to server (probably using a service-mesh)
// - [ ] Perfect the schema mapping between the Firestore existing types and the
// protobufs declarations TODO(oliwierost)
// - [ ] Add Firebase Authentication and Authorization
// - [ ] Implement the remainder of the CRUD operations for all of the items
// - [ ] Ensure idempotency of the application (no double withdrawals),
// transactional support should be enabled by default in Firestore since it is
// based in Spanner
// - [ ] add a .cloudignore file
// - [ ] add a test-case for passing on the X-Firebase-ID-Token header as grpc metadata
// - [ ] set up auth tests
// - [ ] include information about the X-Firebase-ID-Token header in the OpenAPI
// spec (grpc-gateway allows customization like that)
// - [ ] some of the fields in the protobuf might be marked as private? (e.g. the
// transactions for a hero)
// more im thinking if say there should not be a need-to-know parameter for example
// /v1/heroes/{hero_id}?include_transactions=true
// this would then say require MFA or something
// - [ ] write an integration test case to see if the gateway + server work together
// TODO create a client here and retrieve the docs before authorizing any reqeuest
// - [ ] when updating make sure that update mask by default prevents writing
// sensitive fields, like the transactions - OUTPUT_ONLY field from import "google/api/field_behavior.proto";
// - [ ] https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/#hiding-fields-methods-services-and-enum-values
//- [ ] ensure that the server is not susceptible to a common pitfalls, as DoS
//attack, memory leaks, man in the middle, xss, sql injection, etc.
//- [ ] separate the service account firestore-editor@ that currently is used by
//both services (and testing locally as well as during remote CI/CD)
// - [ ] path params like:
// option (google.api.http) = {
//  get: "/v1/{name=shelves/*/books/*}"
// };
// - [ ] firestore on clientside SDK like in TypeScript updates the data live on website when document updates,
// see if this can be achieved with the grpc-gateway and streaming?
// - [ ]
// There is an issue with the `github.com/infobloxopen/protoc-gen-gorm` package, it
// uses deprecated version of `gorm.io/gorm` that comes from another git repository
// that has been moved to a new API
// this has to be done soon since otherwise buf generate fails and it is brittle, we need a custom fork of that repo with the issues fixed

func main() {
	cmd.Execute()
}
