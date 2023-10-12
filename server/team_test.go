package server_test

import (
	"context"
	"testing"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/asoltd/lancr/server"
)

func TestCreateTeam(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewTeamsServiceServer(db)

	req := &lancrv1.CreateTeamRequest{
		Payload: &lancrv1.Team{
			Id: "test-stub",
		},
	}
	_, err := h.CreateTeam(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestListTeams(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewTeamsServiceServer(db)

	req := &lancrv1.ListTeamsRequest{
		PageSize: 10,
	}
	_, err := h.ListTeams(context.Background(), req)
	if err != nil {
		t.Fatalf("ListTeams() failed: %v", err)
	}
}

// TestReadTeam Depends on CreateTeam, will fail if ran on its own or CreateTeam fails
func TestReadTeam(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewTeamsServiceServer(db)

	req := &lancrv1.ReadTeamRequest{
		Id: "test-stub",
	}
	_, err := h.ReadTeam(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateTeam Depends on CreateTeam, will fail if ran on its own or CreateTeam fails
func TestUpdateTeam(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewTeamsServiceServer(db)

	_, err := h.UpdateTeam(context.Background(), &lancrv1.UpdateTeamRequest{
		Payload: &lancrv1.Team{
			Id:       "test-stub",
			Industry: lancrv1.Industry_INDUSTRY_ACCOUNTING,
		},
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTeam(t *testing.T) {
	db := server.SetupTestDB(t)

	h := server.NewTeamsServiceServer(db)

	_, err := h.DeleteTeam(context.Background(), &lancrv1.DeleteTeamRequest{
		Id: "test-stub",
	})
	if err != nil {
		t.Error(err)
	}
}
