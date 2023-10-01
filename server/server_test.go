// server_test package only tests the backend implementation, it does not test
// the gateway nor the gRPC transport layer.
// all of the tests run on the staging environment, classic >>> mockist
package server_test

import (
	"context"
	"testing"

	lancrv1 "github.com/asoltd/lancr/gen/go/proto/lancr/v1"
	"github.com/asoltd/lancr/server"
)

func TestBackendNew(t *testing.T) {
	t.Parallel()

	_, err := server.New()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetHero(t *testing.T) {
	t.Parallel()

	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	req := &lancrv1.GetHeroRequest{
		Id: "test-stub",
	}
	_, err = b.GetHero(context.TODO(), req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListHeroes(t *testing.T) {
	t.Parallel()

	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	req := &lancrv1.ListHeroesRequest{
		PageSize: 10,
	}
	_, err = b.ListHeroes(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}

// TestCreatehero is not relevant at the moment, used this to verify method
// works, Update will be preferred in the future
// In order to run this test, the ID field in the test-stub-hero.json file can be changed
// to a unique value, otherwise the test will fail due to the ID already existing
func TestCreateHero(t *testing.T) {
	t.Skip()

	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	hero := &lancrv1.Hero{}
	err = server.UnmarshalJSONFileToMessage("../test-stub-hero.json", hero)
	if err != nil {
		t.Fatal(err)
	}

	req := &lancrv1.CreateHeroRequest{
		Hero: hero,
	}
	_, err = b.CreateHero(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteHero(t *testing.T) {
	t.Skip()

	b, err := server.New()
	if err != nil {
		t.Fatal(err)
	}

	_, err = b.DeleteHero(context.Background(), &lancrv1.DeleteHeroRequest{
		Id: "test-stub",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = b.GetHero(context.Background(), &lancrv1.GetHeroRequest{Id: "test-stub"})
	if err == nil {
		t.Fatalf("hero was not deleted %v", err)
	}
}
