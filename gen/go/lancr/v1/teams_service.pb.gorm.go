package v1

import (
	context "context"
	gorm "gorm.io/gorm"
)

type TeamsServiceDefaultServer struct {
	DB *gorm.DB
}

// ReadTeam ...
func (m *TeamsServiceDefaultServer) ReadTeam(ctx context.Context, in *ReadTeamRequest) (*ReadTeamResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(TeamsServiceTeamWithBeforeReadTeam); ok {
		var err error
		if db, err = custom.BeforeReadTeam(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadTeam(ctx, &Team{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadTeamResponse{Result: res}
	if custom, ok := interface{}(in).(TeamsServiceTeamWithAfterReadTeam); ok {
		var err error
		if err = custom.AfterReadTeam(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// TeamsServiceTeamWithBeforeReadTeam called before DefaultReadTeamTeam in the default ReadTeam handler
type TeamsServiceTeamWithBeforeReadTeam interface {
	BeforeReadTeam(context.Context, *gorm.DB) (*gorm.DB, error)
}

// TeamsServiceTeamWithAfterReadTeam called before DefaultReadTeamTeam in the default ReadTeam handler
type TeamsServiceTeamWithAfterReadTeam interface {
	AfterReadTeam(context.Context, *ReadTeamResponse, *gorm.DB) error
}

// ListTeams ...
func (m *TeamsServiceDefaultServer) ListTeams(ctx context.Context, in *ListTeamsRequest) (*ListTeamsResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(TeamsServiceTeamWithBeforeListTeams); ok {
		var err error
		if db, err = custom.BeforeListTeams(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListTeam(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListTeamsResponse{Results: res}
	if custom, ok := interface{}(in).(TeamsServiceTeamWithAfterListTeams); ok {
		var err error
		if err = custom.AfterListTeams(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// TeamsServiceTeamWithBeforeListTeams called before DefaultListTeamsTeam in the default ListTeams handler
type TeamsServiceTeamWithBeforeListTeams interface {
	BeforeListTeams(context.Context, *gorm.DB) (*gorm.DB, error)
}

// TeamsServiceTeamWithAfterListTeams called before DefaultListTeamsTeam in the default ListTeams handler
type TeamsServiceTeamWithAfterListTeams interface {
	AfterListTeams(context.Context, *ListTeamsResponse, *gorm.DB) error
}

// CreateTeam ...
func (m *TeamsServiceDefaultServer) CreateTeam(ctx context.Context, in *CreateTeamRequest) (*CreateTeamResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(TeamsServiceTeamWithBeforeCreateTeam); ok {
		var err error
		if db, err = custom.BeforeCreateTeam(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateTeam(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateTeamResponse{Result: res}
	if custom, ok := interface{}(in).(TeamsServiceTeamWithAfterCreateTeam); ok {
		var err error
		if err = custom.AfterCreateTeam(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// TeamsServiceTeamWithBeforeCreateTeam called before DefaultCreateTeamTeam in the default CreateTeam handler
type TeamsServiceTeamWithBeforeCreateTeam interface {
	BeforeCreateTeam(context.Context, *gorm.DB) (*gorm.DB, error)
}

// TeamsServiceTeamWithAfterCreateTeam called before DefaultCreateTeamTeam in the default CreateTeam handler
type TeamsServiceTeamWithAfterCreateTeam interface {
	AfterCreateTeam(context.Context, *CreateTeamResponse, *gorm.DB) error
}

// UpdateTeam ...
func (m *TeamsServiceDefaultServer) UpdateTeam(ctx context.Context, in *UpdateTeamRequest) (*UpdateTeamResponse, error) {
	var err error
	var res *Team
	db := m.DB
	if custom, ok := interface{}(in).(TeamsServiceTeamWithBeforeUpdateTeam); ok {
		var err error
		if db, err = custom.BeforeUpdateTeam(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err = DefaultStrictUpdateTeam(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &UpdateTeamResponse{Result: res}
	if custom, ok := interface{}(in).(TeamsServiceTeamWithAfterUpdateTeam); ok {
		var err error
		if err = custom.AfterUpdateTeam(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// TeamsServiceTeamWithBeforeUpdateTeam called before DefaultUpdateTeamTeam in the default UpdateTeam handler
type TeamsServiceTeamWithBeforeUpdateTeam interface {
	BeforeUpdateTeam(context.Context, *gorm.DB) (*gorm.DB, error)
}

// TeamsServiceTeamWithAfterUpdateTeam called before DefaultUpdateTeamTeam in the default UpdateTeam handler
type TeamsServiceTeamWithAfterUpdateTeam interface {
	AfterUpdateTeam(context.Context, *UpdateTeamResponse, *gorm.DB) error
}

// DeleteTeam ...
func (m *TeamsServiceDefaultServer) DeleteTeam(ctx context.Context, in *DeleteTeamRequest) (*DeleteTeamResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(TeamsServiceTeamWithBeforeDeleteTeam); ok {
		var err error
		if db, err = custom.BeforeDeleteTeam(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteTeam(ctx, &Team{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteTeamResponse{}
	if custom, ok := interface{}(in).(TeamsServiceTeamWithAfterDeleteTeam); ok {
		var err error
		if err = custom.AfterDeleteTeam(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// TeamsServiceTeamWithBeforeDeleteTeam called before DefaultDeleteTeamTeam in the default DeleteTeam handler
type TeamsServiceTeamWithBeforeDeleteTeam interface {
	BeforeDeleteTeam(context.Context, *gorm.DB) (*gorm.DB, error)
}

// TeamsServiceTeamWithAfterDeleteTeam called before DefaultDeleteTeamTeam in the default DeleteTeam handler
type TeamsServiceTeamWithAfterDeleteTeam interface {
	AfterDeleteTeam(context.Context, *DeleteTeamResponse, *gorm.DB) error
}
