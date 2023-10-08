package v1

import (
	context "context"
	gorm "gorm.io/gorm"
)

type QuestServiceDefaultServer struct {
	DB *gorm.DB
}

// ReadQuest ...
func (m *QuestServiceDefaultServer) ReadQuest(ctx context.Context, in *ReadQuestRequest) (*ReadQuestResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(QuestServiceQuestWithBeforeReadQuest); ok {
		var err error
		if db, err = custom.BeforeReadQuest(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultReadQuest(ctx, &Quest{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &ReadQuestResponse{Result: res}
	if custom, ok := interface{}(in).(QuestServiceQuestWithAfterReadQuest); ok {
		var err error
		if err = custom.AfterReadQuest(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// QuestServiceQuestWithBeforeReadQuest called before DefaultReadQuestQuest in the default ReadQuest handler
type QuestServiceQuestWithBeforeReadQuest interface {
	BeforeReadQuest(context.Context, *gorm.DB) (*gorm.DB, error)
}

// QuestServiceQuestWithAfterReadQuest called before DefaultReadQuestQuest in the default ReadQuest handler
type QuestServiceQuestWithAfterReadQuest interface {
	AfterReadQuest(context.Context, *ReadQuestResponse, *gorm.DB) error
}

// ListQuests ...
func (m *QuestServiceDefaultServer) ListQuests(ctx context.Context, in *ListQuestsRequest) (*ListQuestsResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(QuestServiceQuestWithBeforeListQuests); ok {
		var err error
		if db, err = custom.BeforeListQuests(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultListQuest(ctx, db)
	if err != nil {
		return nil, err
	}
	out := &ListQuestsResponse{Results: res}
	if custom, ok := interface{}(in).(QuestServiceQuestWithAfterListQuests); ok {
		var err error
		if err = custom.AfterListQuests(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// QuestServiceQuestWithBeforeListQuests called before DefaultListQuestsQuest in the default ListQuests handler
type QuestServiceQuestWithBeforeListQuests interface {
	BeforeListQuests(context.Context, *gorm.DB) (*gorm.DB, error)
}

// QuestServiceQuestWithAfterListQuests called before DefaultListQuestsQuest in the default ListQuests handler
type QuestServiceQuestWithAfterListQuests interface {
	AfterListQuests(context.Context, *ListQuestsResponse, *gorm.DB) error
}

// CreateQuest ...
func (m *QuestServiceDefaultServer) CreateQuest(ctx context.Context, in *CreateQuestRequest) (*CreateQuestResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(QuestServiceQuestWithBeforeCreateQuest); ok {
		var err error
		if db, err = custom.BeforeCreateQuest(ctx, db); err != nil {
			return nil, err
		}
	}
	res, err := DefaultCreateQuest(ctx, in.GetPayload(), db)
	if err != nil {
		return nil, err
	}
	out := &CreateQuestResponse{Result: res}
	if custom, ok := interface{}(in).(QuestServiceQuestWithAfterCreateQuest); ok {
		var err error
		if err = custom.AfterCreateQuest(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// QuestServiceQuestWithBeforeCreateQuest called before DefaultCreateQuestQuest in the default CreateQuest handler
type QuestServiceQuestWithBeforeCreateQuest interface {
	BeforeCreateQuest(context.Context, *gorm.DB) (*gorm.DB, error)
}

// QuestServiceQuestWithAfterCreateQuest called before DefaultCreateQuestQuest in the default CreateQuest handler
type QuestServiceQuestWithAfterCreateQuest interface {
	AfterCreateQuest(context.Context, *CreateQuestResponse, *gorm.DB) error
}

// UpdateQuest ...
func (m *QuestServiceDefaultServer) UpdateQuest(ctx context.Context, in *UpdateQuestRequest) (*UpdateQuestResponse, error) {
	out := &UpdateQuestResponse{}
	return out, nil
}

// DeleteQuest ...
func (m *QuestServiceDefaultServer) DeleteQuest(ctx context.Context, in *DeleteQuestRequest) (*DeleteQuestResponse, error) {
	db := m.DB
	if custom, ok := interface{}(in).(QuestServiceQuestWithBeforeDeleteQuest); ok {
		var err error
		if db, err = custom.BeforeDeleteQuest(ctx, db); err != nil {
			return nil, err
		}
	}
	err := DefaultDeleteQuest(ctx, &Quest{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	out := &DeleteQuestResponse{}
	if custom, ok := interface{}(in).(QuestServiceQuestWithAfterDeleteQuest); ok {
		var err error
		if err = custom.AfterDeleteQuest(ctx, out, db); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// QuestServiceQuestWithBeforeDeleteQuest called before DefaultDeleteQuestQuest in the default DeleteQuest handler
type QuestServiceQuestWithBeforeDeleteQuest interface {
	BeforeDeleteQuest(context.Context, *gorm.DB) (*gorm.DB, error)
}

// QuestServiceQuestWithAfterDeleteQuest called before DefaultDeleteQuestQuest in the default DeleteQuest handler
type QuestServiceQuestWithAfterDeleteQuest interface {
	AfterDeleteQuest(context.Context, *DeleteQuestResponse, *gorm.DB) error
}
