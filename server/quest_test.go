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

func TestCreateQuest(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewQuestServiceServer(db)

	req := &lancrv1.CreateQuestRequest{
		Payload: &lancrv1.Quest{
			Id: "test-stub",
		},
	}
	_, err := h.CreateQuest(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestListQuests(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewQuestServiceServer(db)

	req := &lancrv1.ListQuestsRequest{
		PageSize: 10,
	}
	_, err := h.ListQuests(context.Background(), req)
	if err != nil {
		t.Fatalf("ListQuests() failed: %v", err)
	}
}

// TestReadQuest Depends on CreateQuest, will fail if ran on its own or CreateQuest fails
func TestReadQuest(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewQuestServiceServer(db)

	req := &lancrv1.ReadQuestRequest{
		Id: "test-stub",
	}
	_, err := h.ReadQuest(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateQuest Depends on CreateQuest, will fail if ran on its own or CreateQuest fails
func TestUpdateQuest(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewQuestServiceServer(db)

	_, err := h.UpdateQuest(context.Background(), &lancrv1.UpdateQuestRequest{
		Payload: &lancrv1.Quest{
			Id:    "test-stub",
			Title: "updated test quest",
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteQuest(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewQuestServiceServer(db)

	_, err := h.DeleteQuest(context.Background(), &lancrv1.DeleteQuestRequest{
		Id: "test-stub",
	})
	if err != nil {
		t.Error(err)
	}
}
