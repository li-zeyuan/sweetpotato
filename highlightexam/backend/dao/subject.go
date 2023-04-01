package dao

import (
	"context"
	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"go.uber.org/zap"
)

type Subject struct{}

func (s *Subject) List(ctx context.Context) ([]*model.SubjectTable, error) {
	subjects := make([]*model.SubjectTable, 0)
	err := mysqlstore.Db.Table(model.TableNameSubject).
		WithContext(ctx).
		Where(comModel.DefaultDelCondition).
		Order(comModel.UpdatedAtDESCCondition).
		Find(&subjects).Error
	if err != nil {
		mylogger.Error(ctx ,"get all subject error: ", zap.Error(err))
		return nil, err
	}

	return subjects, nil
}

func (s *Subject) MapByIds(ctx context.Context, ids []int64) (map[int64]*model.SubjectTable, error) {
	if len(ids) == 0 {
		return nil ,nil
	}

	subjects := make([]*model.SubjectTable, 0)
	err := mysqlstore.Db.Table(model.TableNameSubject).
		WithContext(ctx).
		Where("id in ?", ids).
		Where(comModel.DefaultDelCondition).
		Find(&subjects).Error
	if err != nil {
		mylogger.Error(ctx, "get all subject by id error: ", zap.Error(err))
		return nil, err
	}

	subjectMap := make(map[int64]*model.SubjectTable)
	for _, subject := range subjects {
		subjectMap[subject.ID] = subject
	}

	return subjectMap, nil
}