package server

import (
	"context"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"gorm.io/gorm"
)

type ApprenticeServiceServer struct {
	lancrv1.ApprenticeServiceDefaultServer

	lancrv1.UnimplementedApprenticeServiceServer
}

func NewApprenticeServiceServer(db *gorm.DB) *ApprenticeServiceServer {
	return &ApprenticeServiceServer{
		ApprenticeServiceDefaultServer: lancrv1.ApprenticeServiceDefaultServer{
			DB: db,
		},
	}
}

func (a *ApprenticeServiceServer) ReadApprentice(ctx context.Context, req *lancrv1.ReadApprenticeRequest) (*lancrv1.ReadApprenticeResponse, error) {
	return a.ApprenticeServiceDefaultServer.ReadApprentice(ctx, req)
}

// TODO implement a pager for this method
func (a *ApprenticeServiceServer) ListApprentices(ctx context.Context, req *lancrv1.ListApprenticesRequest) (*lancrv1.ListApprenticesResponse, error) {
	return a.ApprenticeServiceDefaultServer.ListApprentices(ctx, req)
}

func (a *ApprenticeServiceServer) CreateApprentice(ctx context.Context, req *lancrv1.CreateApprenticeRequest) (*lancrv1.CreateApprenticeResponse, error) {
	return a.ApprenticeServiceDefaultServer.CreateApprentice(ctx, req)
}

func (a *ApprenticeServiceServer) UpdateApprentice(ctx context.Context, req *lancrv1.UpdateApprenticeRequest) (*lancrv1.UpdateApprenticeResponse, error) {
	apprentice := req.GetPayload()
	apprenticeORM, err := apprentice.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	err = a.DB.Save(&apprenticeORM).Error
	if err != nil {
		return nil, err
	}
	var res *lancrv1.ApprenticeORM
	err = a.DB.First(&res).Error
	if err != nil {
		return nil, err
	}
	pbRes, err := res.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &lancrv1.UpdateApprenticeResponse{Result: &pbRes}, nil
}

func (a *ApprenticeServiceServer) DeleteApprentice(ctx context.Context, req *lancrv1.DeleteApprenticeRequest) (*lancrv1.DeleteApprenticeResponse, error) {
	return a.ApprenticeServiceDefaultServer.DeleteApprentice(ctx, req)
}
