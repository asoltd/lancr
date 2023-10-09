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
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	// Open a test database connection (SQLite in-memory for this example)
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}

	err = db.AutoMigrate(&lancrv1.HeroORM{})
	if err != nil {
		t.Fatalf("Failed to auto-migrate: %v", err)
	}

	return db
}

func TestCreateHero(t *testing.T) {
	db := setupTestDB(t)

	h := server.NewHeroServiceServer(db)

	req := &lancrv1.CreateHeroRequest{
		Payload: &lancrv1.Hero{
			Id:          "test-stub",
			DisplayName: "test guy",
		},
	}
	_, err := h.CreateHero(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestListHeroes(t *testing.T) {
	db := setupTestDB(t)

	h := server.NewHeroServiceServer(db)

	req := &lancrv1.ListHeroesRequest{
		PageSize: 10,
	}
	_, err := h.ListHeroes(context.Background(), req)
	if err != nil {
		t.Fatalf("ListHeroes() failed: %v", err)
	}
}

func TestReadHero(t *testing.T) {
	db := setupTestDB(t)

	h := server.NewHeroServiceServer(db)

	req := &lancrv1.ReadHeroRequest{
		Id: "test-stub",
	}
	_, err := h.ReadHero(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteHero(t *testing.T) {
	db := setupTestDB(t)

	h := server.NewHeroServiceServer(db)

	_, err := h.DeleteHero(context.Background(), &lancrv1.DeleteHeroRequest{
		Id: "test-stub",
	})
	if err != nil {
		t.Error(err)
	}
}
