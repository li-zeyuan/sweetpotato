package dao

import (
	"context"
	"time"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"go.uber.org/zap"
)

type User struct{}

func (u *User) GetByOpenid(ctx context.Context, openid string) (*model.UserProfileTable, error) {
	m := new(model.UserProfileTable)
	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		WithContext(ctx).
		Where("openid = ?", openid).
		First(m).Error
	if err != nil {
		mylogger.Error(ctx, "get user by openid error: ", zap.Error(err))
		return nil, err
	}

	return m, nil
}

func (u *User) Save(ctx context.Context, models []*model.UserProfileTable) error {
	if len(models) == 0 {
		return nil
	}

	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		WithContext(ctx).
		Create(&models).Error
	if err != nil {
		mylogger.Error(ctx, "create users error: ", zap.Error(err))
		return err
	}

	return nil
}

func (u *User) GetOne(ctx context.Context, uid int64) (*model.UserProfileTable, error) {
	m := new(model.UserProfileTable)
	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		WithContext(ctx).
		Where("id = ?", uid).
		Where(comModel.DefaultDelCondition).
		First(m).Error
	if err != nil {
		mylogger.Error(ctx, "get user by uid error: ", zap.Error(err))
		return nil, err
	}

	return m, nil
}

func (u *User) Update(ctx context.Context, uid int64, fieldMap map[string]interface{}) error {
	if len(fieldMap) == 0 {
		return nil
	}

	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		WithContext(ctx).
		Where("id = ?", uid).
		Where(comModel.DefaultDelCondition).
		Updates(fieldMap).Error
	if err != nil {
		mylogger.Error(ctx, "update user error: ", zap.Error(err))
		return err
	}

	return nil
}

func (u *User) StudyLastTime(ctx context.Context, uid int64) (time.Time, error) {
	lastTime := time.Time{}
	err := mysqlstore.Db.Table(model.TableNameUserProfile).
		WithContext(ctx).
		Select("study_last_time").
		Where("id = ?", uid).
		Where(comModel.DefaultDelCondition).
		Scan(&lastTime).
		Error
	if err != nil {
		mylogger.Error(ctx, "select study last time error: ", zap.Error(err))
		return time.Time{}, err
	}

	return lastTime, nil
}
