package v1

import (
	context "context"
	gorm "gorm.io/gorm"
)

type HeroServiceDefaultServer struct {
	DB *gorm.DB
}

// ReadHero ...
func (m *HeroServiceDefaultServer) ReadHero(ctx context.Context, in *ReadHeroRequest) (*ReadHeroResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(HeroServiceHeroWithBeforeReadHero); ok {
		var err error
		if db, err = custom.BeforeReadHero(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadHero(ctx, &Hero{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadHeroResponse{Result: res}
	if custom, ok := interface{}(in).(HeroServiceHeroWithAfterReadHero); ok {
		var err error
		if err = custom.AfterReadHero(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// HeroServiceHeroWithBeforeReadHero called before DefaultReadHeroHero in the default ReadHero handler
type HeroServiceHeroWithBeforeReadHero interface {
	BeforeReadHero(context.Context, *gorm.DB) (*gorm.DB, error)
}

// HeroServiceHeroWithAfterReadHero called before DefaultReadHeroHero in the default ReadHero handler
type HeroServiceHeroWithAfterReadHero interface {
	AfterReadHero(context.Context, *ReadHeroResponse, *gorm.DB) error
}

// ListHeroes ...
func (m *HeroServiceDefaultServer) ListHeroes(ctx context.Context, in *ListHeroesRequest) (*ListHeroesResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(HeroServiceHeroWithBeforeListHeroes); ok {
		var err error
		if db, err = custom.BeforeListHeroes(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListHero(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListHeroesResponse{Results: res}
	if custom, ok := interface{}(in).(HeroServiceHeroWithAfterListHeroes); ok {
		var err error
		if err = custom.AfterListHeroes(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// HeroServiceHeroWithBeforeListHeroes called before DefaultListHeroesHero in the default ListHeroes handler
type HeroServiceHeroWithBeforeListHeroes interface {
	BeforeListHeroes(context.Context, *gorm.DB) (*gorm.DB, error)
}

// HeroServiceHeroWithAfterListHeroes called before DefaultListHeroesHero in the default ListHeroes handler
type HeroServiceHeroWithAfterListHeroes interface {
	AfterListHeroes(context.Context, *ListHeroesResponse, *gorm.DB) error
}

// CreateHero ...
func (m *HeroServiceDefaultServer) CreateHero(ctx context.Context, in *CreateHeroRequest) (*CreateHeroResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(HeroServiceHeroWithBeforeCreateHero); ok {
		var err error
		if db, err = custom.BeforeCreateHero(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateHero(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateHeroResponse{Result: res}
	if custom, ok := interface{}(in).(HeroServiceHeroWithAfterCreateHero); ok {
		var err error
		if err = custom.AfterCreateHero(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// HeroServiceHeroWithBeforeCreateHero called before DefaultCreateHeroHero in the default CreateHero handler
type HeroServiceHeroWithBeforeCreateHero interface {
	BeforeCreateHero(context.Context, *gorm.DB) (*gorm.DB, error)
}

// HeroServiceHeroWithAfterCreateHero called before DefaultCreateHeroHero in the default CreateHero handler
type HeroServiceHeroWithAfterCreateHero interface {
	AfterCreateHero(context.Context, *CreateHeroResponse, *gorm.DB) error
}

// UpdateHero ...
func (m *HeroServiceDefaultServer) UpdateHero(ctx context.Context, in *UpdateHeroRequest) (*UpdateHeroResponse, error) {
	out := &UpdateHeroResponse{}
	return out, nil
}

// DeleteHero ...
func (m *HeroServiceDefaultServer) DeleteHero(ctx context.Context, in *DeleteHeroRequest) (*DeleteHeroResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(HeroServiceHeroWithBeforeDeleteHero); ok {
		var err error
		if db, err = custom.BeforeDeleteHero(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteHero(ctx, &Hero{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteHeroResponse{}
	if custom, ok := interface{}(in).(HeroServiceHeroWithAfterDeleteHero); ok {
		var err error
		if err = custom.AfterDeleteHero(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// HeroServiceHeroWithBeforeDeleteHero called before DefaultDeleteHeroHero in the default DeleteHero handler
type HeroServiceHeroWithBeforeDeleteHero interface {
	BeforeDeleteHero(context.Context, *gorm.DB) (*gorm.DB, error)
}

// HeroServiceHeroWithAfterDeleteHero called before DefaultDeleteHeroHero in the default DeleteHero handler
type HeroServiceHeroWithAfterDeleteHero interface {
	AfterDeleteHero(context.Context, *DeleteHeroResponse, *gorm.DB) error
}
