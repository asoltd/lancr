package server

import (
	"context"

	lancrv1 "github.com/asoltd/lancr/gen/go/lancr/v1"
	"gorm.io/gorm"
)

type QuestServiceServer struct {
	lancrv1.QuestServiceDefaultServer

	lancrv1.UnimplementedQuestServiceServer
}

func NewQuestServiceServer(db *gorm.DB) *QuestServiceServer {
	return &QuestServiceServer{
		QuestServiceDefaultServer: lancrv1.QuestServiceDefaultServer{
			DB: db,
		},
	}
}

func (q *QuestServiceServer) ReadQuest(ctx context.Context, req *lancrv1.ReadQuestRequest) (*lancrv1.ReadQuestResponse, error) {
	return q.QuestServiceDefaultServer.ReadQuest(ctx, req)
}

// TODO implement a pager for this method
func (q *QuestServiceServer) ListQuests(ctx context.Context, req *lancrv1.ListQuestsRequest) (*lancrv1.ListQuestsResponse, error) {
	return q.QuestServiceDefaultServer.ListQuests(ctx, req)
}

func (q *QuestServiceServer) CreateQuest(ctx context.Context, req *lancrv1.CreateQuestRequest) (*lancrv1.CreateQuestResponse, error) {
	return q.QuestServiceDefaultServer.CreateQuest(ctx, req)
}

func (q *QuestServiceServer) UpdateQuest(ctx context.Context, req *lancrv1.UpdateQuestRequest) (*lancrv1.UpdateQuestResponse, error) {
	return q.QuestServiceDefaultServer.UpdateQuest(ctx, req)
}

func (q *QuestServiceServer) DeleteQuest(ctx context.Context, req *lancrv1.DeleteQuestRequest) (*lancrv1.DeleteQuestResponse, error) {
	return q.QuestServiceDefaultServer.DeleteQuest(ctx, req)
}
