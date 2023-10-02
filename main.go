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

func main() {
	cmd.Execute()
}
