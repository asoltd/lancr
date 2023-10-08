// the methods order is important, parallel won't work here until some proper
// mocking is in place
//
// ensure that the tests are running on the 'test' database and not the 'prod',
// there is no mocks here
package server_test

import (
	"context"
	"log"
	"testing"

	"github.com/asoltd/lancr/db"
	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/asoltd/lancr/server"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Open a test database connection (SQLite in-memory for this example)
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate your database schema here
	// db.AutoMigrate(&YourModel{})

	return db
}

func TestListHeroes(t *testing.T) {
	db := setupTestDB()

	h := server.NewHeroServiceServer(db)

	req := &lancrv1.ListHeroesRequest{
		PageSize: 10,
	}
	_, err := h.ListHeroes(context.Background(), req)
	if err != nil {
		t.Fatalf("ListHeroes() failed: %v", err)
	}
}

func TestCreateHero(t *testing.T) {
	db := setupTestDB()

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

func TestReadHero(t *testing.T) {
	db := setupTestDB()

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
	db := setupTestDB()

	h := server.NewHeroServiceServer(db)

	_, err := h.DeleteHero(context.Background(), &lancrv1.DeleteHeroRequest{
		Id: "test-stub",
	})
	if err != nil {
		t.Error(err)
	}
}
