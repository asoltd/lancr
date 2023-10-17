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
// - [ ] Ensure idempotency of the application (no double withdrawals), use request_id parameter
// - [ ] add a .cloudignore file
// - [ ] include information about the X-Firebase-ID-Token header in the OpenAPI
// spec (grpc-gateway allows customization like that)
// - [ ] write an integration test case to see if the gateway + server work together
// TODO create a client here and retrieve the docs before authorizing any reqeuest
// - [ ] when updating make sure that update mask by default prevents writing
// sensitive fields, like the transactions - OUTPUT_ONLY field from import
// "google/api/field_behavior.proto";
// - [ ] https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/customizing_openapi_output/#hiding-fields-methods-services-and-enum-values
//- [ ] ensure that the server is not susceptible to a common pitfalls, as DoS
//attack, memory leaks, man in the middle, xss, sql injection, etc.
// - [ ] create and apply the gateway-sa service account (firebase auth)
// - [ ] look for an open-source example of a gateway with auth integrated,
// something like twitter or discord for reference
// - [ ] path params like:
// option (google.api.http) = {
//  get: "/v1/{name=shelves/*/books/*}"
// };
// - [ ] firestore on clientside SDK like in TypeScript updates the data live on website when document updates,
// see if this can be achieved with the grpc-gateway and streaming?
// - [ ] manaage secrets like CLICKHOUSE_PASSWORD in production, i want to use sth like spiffe
// - [ ] set up proper logs (zap?)
// - [ ] add healthchecks
// - [ ] set up MFA for changing certain settings
// - [ ] handling of create Hero, creation of User automatically and Hero being
// customizable based on the UX, Hero is public-readable, User is private
// - [ ] the one-to-many, many-to-many, belongs-to, ..., relationships - based on the MVP functionality
// - [ ] handling of the errors and returning the errors which can be displayed nicely
// - [ ] asynchronous notifications service
// - [ ] is there a need for having a message queue running? some of the updates
// around bids etc would be very nice to work live
// - [ ] the specific services should not have access to all tables, every
// service should only be able to access its own corresponding table
// - [ ] implement paging for list methods, implement search
// - [ ] do you need to ensure that
// 1. the table is going to be the same as the ORM object
// 2. the payload is not nil, is the parameters in request gonna be required
// - [ ] set up cloud armor for dddos protection
// = [ ] set up static files hosting
// - [ ] set up custom create so that there are say no quests that have the same
// name
// - [ ] the client has to re-generate automatically, same as buf generate
// - [ ] implement automatic retries

func main() {
	cmd.Execute()
}
