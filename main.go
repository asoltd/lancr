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

func main() {
	cmd.Execute()
}
