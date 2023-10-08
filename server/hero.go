package server

import (
	"context"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"gorm.io/gorm"
)

type HeroServiceServer struct {
	lancrv1.HeroServiceDefaultServer

	lancrv1.UnimplementedHeroServiceServer
}

func NewHeroServiceServer(db *gorm.DB) *HeroServiceServer {
	return &HeroServiceServer{
		HeroServiceDefaultServer: lancrv1.HeroServiceDefaultServer{
			DB: db,
		},
	}
}

func (h *HeroServiceServer) ReadHero(ctx context.Context, req *lancrv1.ReadHeroRequest) (*lancrv1.ReadHeroResponse, error) {
	return h.HeroServiceDefaultServer.ReadHero(ctx, req)
}

// TODO implement a pager for this method
func (h *HeroServiceServer) ListHeroes(ctx context.Context, req *lancrv1.ListHeroesRequest) (*lancrv1.ListHeroesResponse, error) {
	return h.HeroServiceDefaultServer.ListHeroes(ctx, req)
}

func (h *HeroServiceServer) CreateHero(ctx context.Context, req *lancrv1.CreateHeroRequest) (*lancrv1.CreateHeroResponse, error) {
	return h.HeroServiceDefaultServer.CreateHero(ctx, req)
}

func (h *HeroServiceServer) UpdateHero(ctx context.Context, req *lancrv1.UpdateHeroRequest) (*lancrv1.UpdateHeroResponse, error) {
	return h.HeroServiceDefaultServer.UpdateHero(ctx, req)
}

func (h *HeroServiceServer) DeleteHero(ctx context.Context, req *lancrv1.DeleteHeroRequest) (*lancrv1.DeleteHeroResponse, error) {
	return h.HeroServiceDefaultServer.DeleteHero(ctx, req)
}
