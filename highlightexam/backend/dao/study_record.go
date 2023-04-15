package dao

import (
	"context"
	"fmt"

	"github.com/li-zeyuan/common/utils"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

type StudyRecord struct{}

func (s *StudyRecord) ListByUid(ctx context.Context, uid int64, start, limit int) ([]*model.StudyRecord, error) {
	sql := fmt.Sprintf("select subject_id, count from (select subject_id, count(id) as count, max(updated_at) as updated_at from `study_record` where uid = %d and deleted_at = 0 group by subject_id) as temp order by temp.updated_at desc limit %d,%d", uid, start, limit)
	studyRecords := make([]*model.StudyRecord, 0)
	err := mysqlstore.Db.Table(model.TableNameStudyRecord).
		WithContext(ctx).
		Raw(sql).
		Scan(&studyRecords).
		Error
	if err != nil {
		mylogger.Error(ctx, "get study record by uid error: ", zap.Error(err))
		return nil, err
	}

	return studyRecords, nil
}

func (s *StudyRecord) HasStudiedKids(ctx context.Context, uid, subjectID int64) ([]int64, error) {
	kids := make([]int64, 0)
	err := mysqlstore.Db.Table(model.TableNameStudyRecord).
		WithContext(ctx).
		Select("knowledge_id").
		Where("uid = ?", uid).
		Where("subject_id = ?", subjectID).
		Where(comModel.DefaultDelCondition).
		Scan(&kids).Error
	if err != nil {
		mylogger.Error(ctx, "get study record num error: ", zap.Error(err))
		return nil, err
	}

	return kids, nil
}

func (s *StudyRecord) Upsert(ctx context.Context, m *model.StudyRecordTable) error {
	if m == nil {
		return nil
	}

	err := mysqlstore.Db.Table(model.TableNameStudyRecord).
		WithContext(ctx).
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(m).Error
	if err != nil {
		mylogger.Error(ctx, "save record num error: ", zap.Error(err))
		return err
	}

	return nil
}

func (s *StudyRecord) TodayStudyNum(ctx context.Context, uid int64) (int64, error) {
	var num int64
	err := mysqlstore.Db.Table(model.TableNameStudyRecord).
		WithContext(ctx).
		Where("uid = ?", uid).
		Where(comModel.DefaultDelCondition).
		Where(comModel.ColumnUpdatedAt+" > ?", utils.TodayStartUTC()).
		Count(&num).Error
	if err != nil {
		mylogger.Error(ctx, "get today study record num error: ", zap.Error(err))
		return 0, err
	}

	return num, nil
}
