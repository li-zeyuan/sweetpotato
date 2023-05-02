package main

import (
	"context"
	"fmt"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
)

func getKnowledge(idioms []string, subjectID int64) ([]string, error) {
	ctx := context.Background()
	existIdioms := make([]string, 0)
	err := mysqlstore.Db.Table(model.TableNameKnowledge).
		WithContext(ctx).
		Where("subject_id = ?", subjectID).
		Where(comModel.DefaultDelCondition).
		Where("name on ?", idioms).
		Scan(&existIdioms).Error
	if err != nil {
		fmt.Println("get will study knowledge error: ", err)
		return nil, err
	}

	return utils.ExcludeSlice(idioms, existIdioms), nil
}
