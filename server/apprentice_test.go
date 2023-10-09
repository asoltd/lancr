package server_test

import (
	"context"
	"fmt"
	"testing"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"github.com/asoltd/lancr/server"
)

func TestCreateApprentice(t *testing.T) {
	db := server.SetupTestDB(t)

	a := server.NewApprenticeServiceServer(db)

	req := &lancrv1.CreateApprenticeRequest{
		Payload: &lancrv1.Apprentice{
			Id: "test-apprentice",
		},
	}
	_, err := a.CreateApprentice(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

func TestListApprentices(t *testing.T) {
	db := server.SetupTestDB(t)

	a := server.NewApprenticeServiceServer(db)

	req := &lancrv1.ListApprenticesRequest{
		PageSize: 10,
	}
	_, err := a.ListApprentices(context.Background(), req)
	if err != nil {
		t.Fatalf("ListApprentices() failed: %v", err)
	}
}

// TestReadApprentice Depends on CreateApprentice, will fail if ran on its own or CreateApprentice fails
func TestReadApprentice(t *testing.T) {
	db := server.SetupTestDB(t)

	a := server.NewApprenticeServiceServer(db)

	req := &lancrv1.ReadApprenticeRequest{
		Id: "test-apprentice",
	}
	_, err := a.ReadApprentice(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
}

// TestUpdateApprentice Depends on CreateApprentice, will fail if ran on its own or CreateApprentice fails
func TestUpdateApprentice(t *testing.T) {
	db := server.SetupTestDB(t)

	a := server.NewApprenticeServiceServer(db)

	res, err := a.UpdateApprentice(context.Background(), &lancrv1.UpdateApprenticeRequest{
		Payload: &lancrv1.Apprentice{
			Id:     "test-apprentice",
			Mentor: "test-mentor",
		},
	})
	fmt.Println(res)
	if err != nil {
		t.Error(err)
	}

	apprenticeInDB := &lancrv1.ApprenticeORM{}
	err = db.First(&lancrv1.ApprenticeORM{}, "id = ?", "test-apprentice").Scan(&apprenticeInDB).Error
	if err != nil {
		t.Error(err)
	}

	want := "test-mentor"
	pb, err := apprenticeInDB.ToPB(context.Background())
	if err != nil {
		t.Errorf("failed to convert to pb: %v", err)
	}
	if pb.GetMentor() != "test-mentor" {
		if pb.GetMentor() != want {
			t.Errorf("expected %s, got %v", want, pb.GetMentor())
		}
	}
}

func TestDeleteApprentice(t *testing.T) {
	db := server.SetupTestDB(t)

	a := server.NewApprenticeServiceServer(db)

	_, err := a.DeleteApprentice(context.Background(), &lancrv1.DeleteApprenticeRequest{
		Id: "test-apprentice",
	})
	if err != nil {
		t.Error(err)
	}
}
