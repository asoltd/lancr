package server

import (
	"context"
	"fmt"
	"log"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (b *Backend) GetHero(ctx context.Context, req *lancrv1.GetHeroRequest) (*lancrv1.GetHeroResponse, error) {
	ref := b.firestore.Collection("heroes").Doc(req.GetId())
	snapshot, err := ref.Get(ctx)
	if err != nil {
		return nil, err
	}

	if status.Code(err) == codes.NotFound {
		return nil, fmt.Errorf("hero with id %s does not exist", req.GetId())
	}
	if err != nil {
		return nil, err
	}
	hero := &lancrv1.Hero{}
	err = b.snapshotToMessage(snapshot, hero)
	if err != nil {
		log.Printf("failed to serialize firestore snapshot to proto, ref ID: %s", snapshot.Ref.ID)
		return nil, err
	}
	return &lancrv1.GetHeroResponse{Hero: hero}, nil
}

// TODO implement a pager for this method
func (b *Backend) ListHeroes(ctx context.Context, req *lancrv1.ListHeroesRequest) (*lancrv1.ListHeroesResponse, error) {
	res := []*lancrv1.Hero{}
	iter := b.firestore.Collection("heroes").Documents(ctx)
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		hero := &lancrv1.Hero{}
		err = b.snapshotToMessage(snapshot, hero)
		if err != nil {
			log.Printf("failed to parse hero into pb, ref ID: %s", snapshot.Ref.ID)
		}
		res = append(res, hero)
		if len(res) == int(req.GetPageSize()) {
			break
		}
	}
	return &lancrv1.ListHeroesResponse{Heroes: res}, nil
}

func (b *Backend) CreateHero(ctx context.Context, req *lancrv1.CreateHeroRequest) (*lancrv1.CreateHeroResponse, error) {
	hero := req.GetHero()
	if hero == nil {
		return nil, fmt.Errorf("hero is required")
	}

	herostruct, err := messageToStruct(hero)
	if err != nil {
		return nil, fmt.Errorf("failed to convert hero to struct: %v", err)
	}

	// idempotency is important here, just implementing quick to add test-stub hero
	ref := b.firestore.Collection("heroes").Doc(hero.GetId())
	if _, err := ref.Get(ctx); err == nil {
		return nil, fmt.Errorf("hero with id %s already exists", hero.GetId())
	}

	_, err = ref.Create(ctx, herostruct)
	if err != nil {
		return nil, err
	}
	return &lancrv1.CreateHeroResponse{Hero: hero}, nil
}

func (b *Backend) DeleteHero(ctx context.Context, req *lancrv1.DeleteHeroRequest) (*lancrv1.DeleteHeroResponse, error) {
	ref := b.firestore.Collection("heroes").Doc(req.GetId())
	if _, err := ref.Get(ctx); status.Code(err) == codes.NotFound {
		return nil, fmt.Errorf("hero with id %s does not exist", req.GetId())
	}

	_, err := ref.Delete(ctx)
	if err != nil {
		return nil, err
	}

	return &lancrv1.DeleteHeroResponse{}, nil
}
