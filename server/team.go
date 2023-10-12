package server

import (
	"context"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"gorm.io/gorm"
)

type TeamsServiceServer struct {
	lancrv1.TeamsServiceDefaultServer
	lancrv1.TeamsServiceTeamWithBeforeCreateTeam

	lancrv1.UnimplementedTeamsServiceServer
}

func NewTeamsServiceServer(db *gorm.DB) *TeamsServiceServer {
	return &TeamsServiceServer{
		TeamsServiceDefaultServer: lancrv1.TeamsServiceDefaultServer{
			DB: db,
		},
	}
}

func (h *TeamsServiceServer) ReadTeam(ctx context.Context, req *lancrv1.ReadTeamRequest) (*lancrv1.ReadTeamResponse, error) {
	return h.TeamsServiceDefaultServer.ReadTeam(ctx, req)
}

// TODO implement a pager for this method
func (h *TeamsServiceServer) ListTeams(ctx context.Context, req *lancrv1.ListTeamsRequest) (*lancrv1.ListTeamsResponse, error) {
	return h.TeamsServiceDefaultServer.ListTeams(ctx, req)
}

func (h *TeamsServiceServer) CreateTeam(ctx context.Context, req *lancrv1.CreateTeamRequest) (*lancrv1.CreateTeamResponse, error) {
	return h.TeamsServiceDefaultServer.CreateTeam(ctx, req)
}

func (h *TeamsServiceServer) UpdateTeam(ctx context.Context, req *lancrv1.UpdateTeamRequest) (*lancrv1.UpdateTeamResponse, error) {
	hero := req.GetPayload()
	heroORM, err := hero.ToORM(ctx)
	if err != nil {
		return nil, err
	}

	err = h.DB.Save(&heroORM).Error
	if err != nil {
		return nil, err
	}

	var res *lancrv1.TeamORM
	err = h.DB.First(&res).Error
	if err != nil {
		return nil, err
	}

	pbRes, err := res.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &lancrv1.UpdateTeamResponse{Result: &pbRes}, nil
}

func (h *TeamsServiceServer) DeleteTeam(ctx context.Context, req *lancrv1.DeleteTeamRequest) (*lancrv1.DeleteTeamResponse, error) {
	return h.TeamsServiceDefaultServer.DeleteTeam(ctx, req)
}
