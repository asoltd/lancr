package server

import (
	"context"
	"fmt"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
)

func (b *Backend) ReadHero(ctx context.Context, req *lancrv1.ReadHeroRequest) (*lancrv1.ReadHeroResponse, error) {
	hero, err := lancrv1.DefaultReadHero(ctx, &lancrv1.Hero{Id: req.GetId()}, b.db)
	if err != nil {
		return nil, err
	}
	return &lancrv1.ReadHeroResponse{Result: hero}, nil
}

// TODO implement a pager for this method
func (b *Backend) ListHeroes(ctx context.Context, req *lancrv1.ListHeroesRequest) (*lancrv1.ListHeroesResponse, error) {
	res, err := lancrv1.DefaultListHero(ctx, b.db)
	if err != nil {
		return nil, err
	}
	return &lancrv1.ListHeroesResponse{Results: res}, nil
}

func (b *Backend) CreateHero(ctx context.Context, req *lancrv1.CreateHeroRequest) (*lancrv1.CreateHeroResponse, error) {
	hero := req.GetPayload()
	if hero == nil {
		fmt.Printf("CreateHeroRequest is missing hero parameter, %+v, %+v", hero, req)
		return nil, fmt.Errorf("CreateHeroRequest is missing hero parameter")
	}

	result, err := lancrv1.DefaultCreateHero(ctx, hero, b.db)
	if err != nil {
		return nil, err
	}
	return &lancrv1.CreateHeroResponse{Result: result}, nil
}

func (b *Backend) UpdateHero(ctx context.Context, req *lancrv1.UpdateHeroRequest) (*lancrv1.UpdateHeroResponse, error) {
	hero := req.GetPayload()
	result, err := lancrv1.DefaultStrictUpdateHero(ctx, hero, b.db)
	if err != nil {
		return nil, err
	}
	return &lancrv1.UpdateHeroResponse{Result: result}, nil
}

func (b *Backend) DeleteHero(ctx context.Context, req *lancrv1.DeleteHeroRequest) (*lancrv1.DeleteHeroResponse, error) {
	lancrv1.DefaultDeleteHero(ctx, &lancrv1.Hero{Id: req.GetId()}, b.db)
	return &lancrv1.DeleteHeroResponse{}, nil
}
