/*
Copyright © 2023 piotr.jp.ostrowski@gmail.com
*/
package main

import "github.com/asoltd/lancr/cmd"

// TODO's
// - [ ] Add TLS certs for proxy to server
// - [ ] Move config to environment variables / YAML + viper
// - [ ] Move command-line flags to cobra
// - [ ] Separate the backend and frontend and while sticking to monorepo
// - [ ] Perfect the schema mapping between the Firestore existing types and the
// protobufs declarations TODO(oliwierost)
// - [ ] Add Authentication
// - [ ] Add Authorization
// - [ ] Add CI
// - [ ] Deploy, add CD
// - [ ] Implement the remainder of the CRUD operations for all of the items
// - [ ] Ensure idempotency of the application (no double withdrawals),
// transactional support should be enabled by default in Firestore since it is
// based in Spanner
// - [ ] add a .cloudignore file
// - [x] Workload Identity is the preferred way to authenticate to GCP services,
// I will be using Kubernetes secrets for now
// - [ ] Make the DNS resolve properly, still putting localhost:4201 in the
// gateway dialAddress, which is not gonna connect

func main() {
	cmd.Execute()
}
