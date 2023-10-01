/*
Copyright © 2023 piotr.jp.ostrowski@gmail.com
*/
package main

import "github.com/asoltd/lancr/cmd"

// TODO's
// - [ ] Add TLS for load balancer + gateway, get domain etc
// - [ ] Add TLS certs for gateway to server
// - [ ] Move config to environment variables / YAML + viper
// - [ ] Move command-line flags to cobra
// - [ ] Separate the backend and frontend and while sticking to monorepo
// - [ ] Perfect the schema mapping between the Firestore existing types and the
// protobufs declarations TODO(oliwierost)
// - [ ] Add Firebase Authentication and Authorization
// - [ ] Add CI
// - [ ] Deploy, add CD
// - [ ] Implement the remainder of the CRUD operations for all of the items
// - [ ] Ensure idempotency of the application (no double withdrawals),
// transactional support should be enabled by default in Firestore since it is
// based in Spanner
// - [ ] add a .cloudignore file

func main() {
	cmd.Execute()
}
