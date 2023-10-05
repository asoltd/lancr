// the methods order is important, parallel won't work here until some proper
// mocking is in place
//
// ensure that the tests are running on the 'test' database and not the 'prod',
// there is no mocks here
package server_test

import (
	"context"
	"fmt"
	"testing"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/asoltd/lancr/server"
)

func TestListHeroes(t *testing.T) {
	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	req := &lancrv1.ListHeroesRequest{
		PageSize: 10,
	}
	_, err = b.ListHeroes(context.Background(), req)
	if err != nil {
		t.Fatalf("ListHeroes() failed: %v", err)
	}
}

func TestCreateHero(t *testing.T) {
	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	req := &lancrv1.CreateHeroRequest{
		Payload: &lancrv1.Hero{
			Id: "test-stub",
		},
	}
	_, err = b.CreateHero(context.Background(), req)
	if err != nil {
		if err.Error() != fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'PRIMARY'", req.Payload.Id) {
			t.Error(err)
		}
	}
}

func TestReadHero(t *testing.T) {
	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	req := &lancrv1.ReadHeroRequest{
		Id: "test-stub",
	}
	_, err = b.ReadHero(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHero(t *testing.T) {
	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	_, err = b.DeleteHero(context.Background(), &lancrv1.DeleteHeroRequest{
		Id: "test-stub",
	})
	if err != nil {
		t.Error(err)
	}

	_, err = b.ReadHero(context.Background(), &lancrv1.ReadHeroRequest{Id: "test-stub"})
	if err == nil {
		t.Errorf("hero was not deleted %v", err)
	}
}
