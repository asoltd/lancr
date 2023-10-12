// the methods order is important, parallel won't work here until some proper
// mocking is in place
//
// ensure that the tests are running on the 'test' database and not the 'prod',
// there is no mocks here
package server_test

import (
	"context"
	"testing"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/asoltd/lancr/server"
)

func TestCreateHero(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewHeroServiceServer(db)

	req := &lancrv1.CreateHeroRequest{
		Payload: &lancrv1.Hero{
			Id: "test-stub",
		},
	}
	_, err := h.CreateHero(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestListHeroes(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewHeroServiceServer(db)

	req := &lancrv1.ListHeroesRequest{
		PageSize: 10,
	}
	_, err := h.ListHeroes(context.Background(), req)
	if err != nil {
		t.Fatalf("ListHeroes() failed: %v", err)
	}
}

// TestReadHero Depends on CreateHero, will fail if ran on its own or CreateHero fails
func TestReadHero(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewHeroServiceServer(db)

	req := &lancrv1.ReadHeroRequest{
		Id: "test-stub",
	}
	_, err := h.ReadHero(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateHero Depends on CreateHero, will fail if ran on its own or CreateHero fails
func TestUpdateHero(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewHeroServiceServer(db)

	_, err := h.UpdateHero(context.Background(), &lancrv1.UpdateHeroRequest{
		Payload: &lancrv1.Hero{
			Id: "test-stub",
			Name: &lancrv1.Name{
				First: "test",
				Last:  "stub",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHero(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewHeroServiceServer(db)

	_, err := h.DeleteHero(context.Background(), &lancrv1.DeleteHeroRequest{
		Id: "test-stub",
	})
	if err != nil {
		t.Error(err)
	}
}
