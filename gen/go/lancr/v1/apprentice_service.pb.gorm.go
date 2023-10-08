package v1

import (
	context "context"
	gorm "gorm.io/gorm"
)

type ApprenticeServiceDefaultServer struct {
	DB *gorm.DB
}

// ReadApprentice ...
func (m *ApprenticeServiceDefaultServer) ReadApprentice(ctx context.Context, in *ReadApprenticeRequest) (*ReadApprenticeResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithBeforeReadApprentice); ok {
		var err error
		if db, err = custom.BeforeReadApprentice(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadApprentice(ctx, &Apprentice{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadApprenticeResponse{Result: res}
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithAfterReadApprentice); ok {
		var err error
		if err = custom.AfterReadApprentice(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ApprenticeServiceApprenticeWithBeforeReadApprentice called before DefaultReadApprenticeApprentice in the default ReadApprentice handler
type ApprenticeServiceApprenticeWithBeforeReadApprentice interface {
	BeforeReadApprentice(context.Context, *gorm.DB) (*gorm.DB, error)
}

// ApprenticeServiceApprenticeWithAfterReadApprentice called before DefaultReadApprenticeApprentice in the default ReadApprentice handler
type ApprenticeServiceApprenticeWithAfterReadApprentice interface {
	AfterReadApprentice(context.Context, *ReadApprenticeResponse, *gorm.DB) error
}

// ListApprentices ...
func (m *ApprenticeServiceDefaultServer) ListApprentices(ctx context.Context, in *ListApprenticesRequest) (*ListApprenticesResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithBeforeListApprentices); ok {
		var err error
		if db, err = custom.BeforeListApprentices(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListApprentice(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListApprenticesResponse{Results: res}
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithAfterListApprentices); ok {
		var err error
		if err = custom.AfterListApprentices(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ApprenticeServiceApprenticeWithBeforeListApprentices called before DefaultListApprenticesApprentice in the default ListApprentices handler
type ApprenticeServiceApprenticeWithBeforeListApprentices interface {
	BeforeListApprentices(context.Context, *gorm.DB) (*gorm.DB, error)
}

// ApprenticeServiceApprenticeWithAfterListApprentices called before DefaultListApprenticesApprentice in the default ListApprentices handler
type ApprenticeServiceApprenticeWithAfterListApprentices interface {
	AfterListApprentices(context.Context, *ListApprenticesResponse, *gorm.DB) error
}

// CreateApprentice ...
func (m *ApprenticeServiceDefaultServer) CreateApprentice(ctx context.Context, in *CreateApprenticeRequest) (*CreateApprenticeResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithBeforeCreateApprentice); ok {
		var err error
		if db, err = custom.BeforeCreateApprentice(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateApprentice(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateApprenticeResponse{Result: res}
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithAfterCreateApprentice); ok {
		var err error
		if err = custom.AfterCreateApprentice(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ApprenticeServiceApprenticeWithBeforeCreateApprentice called before DefaultCreateApprenticeApprentice in the default CreateApprentice handler
type ApprenticeServiceApprenticeWithBeforeCreateApprentice interface {
	BeforeCreateApprentice(context.Context, *gorm.DB) (*gorm.DB, error)
}

// ApprenticeServiceApprenticeWithAfterCreateApprentice called before DefaultCreateApprenticeApprentice in the default CreateApprentice handler
type ApprenticeServiceApprenticeWithAfterCreateApprentice interface {
	AfterCreateApprentice(context.Context, *CreateApprenticeResponse, *gorm.DB) error
}

// UpdateApprentice ...
func (m *ApprenticeServiceDefaultServer) UpdateApprentice(ctx context.Context, in *UpdateApprenticeRequest) (*UpdateApprenticeResponse, error) {
	out := &UpdateApprenticeResponse{}
	return out, nil
}

// DeleteApprentice ...
func (m *ApprenticeServiceDefaultServer) DeleteApprentice(ctx context.Context, in *DeleteApprenticeRequest) (*DeleteApprenticeResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithBeforeDeleteApprentice); ok {
		var err error
		if db, err = custom.BeforeDeleteApprentice(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteApprentice(ctx, &Apprentice{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteApprenticeResponse{}
	if custom, ok := interface{}(in).(ApprenticeServiceApprenticeWithAfterDeleteApprentice); ok {
		var err error
		if err = custom.AfterDeleteApprentice(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ApprenticeServiceApprenticeWithBeforeDeleteApprentice called before DefaultDeleteApprenticeApprentice in the default DeleteApprentice handler
type ApprenticeServiceApprenticeWithBeforeDeleteApprentice interface {
	BeforeDeleteApprentice(context.Context, *gorm.DB) (*gorm.DB, error)
}

// ApprenticeServiceApprenticeWithAfterDeleteApprentice called before DefaultDeleteApprenticeApprentice in the default DeleteApprentice handler
type ApprenticeServiceApprenticeWithAfterDeleteApprentice interface {
	AfterDeleteApprentice(context.Context, *DeleteApprenticeResponse, *gorm.DB) error
}
