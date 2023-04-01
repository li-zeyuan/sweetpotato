package dao

import (
	"context"
	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"go.uber.org/zap"
)

type Knowledge struct{}

func (k *Knowledge) ListBySubjectID(ctx context.Context, subjectID int64, start, limit int) ([]*model.KnowledgeTable, error) {
	knowledgeList := make([]*model.KnowledgeTable, 0)
	err := mysqlstore.Db.Table(model.TableNameKnowledge).
		WithContext(ctx).
		Where("subject_id = ?", subjectID).
		Where(comModel.DefaultDelCondition).
		Order(comModel.IDASCCondition).
		Offset(start).
		Limit(limit).
		Find(&knowledgeList).Error
	if err != nil {
		mylogger.Error(ctx, "get knowledge by subject error: ", zap.Error(err))
		return nil, err
	}

	return knowledgeList, nil
}


func (k *Knowledge) WillStudiedList(ctx context.Context,subjectID int64,  kids []int64) ([]*model.KnowledgeTable, error) {
	knowledgeList := make([]*model.KnowledgeTable, 0)
	err := mysqlstore.Db.Table(model.TableNameKnowledge).
		WithContext(ctx).
		Where("subject_id = ?", subjectID).
		Where(comModel.DefaultDelCondition).
		Not(map[string]interface{}{"id": kids}).
		Limit(model.DefaultStudyBatch).
		Scan(&knowledgeList).Error
	if err != nil {
		mylogger.Error(ctx, "get will study knowledge error: ", zap.Error(err))
		return nil, err
	}

	return knowledgeList, nil
}