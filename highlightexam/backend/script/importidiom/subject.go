package main

import (
	"context"
	"fmt"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
)

func getSubject(task *Task) (*model.SubjectTable, error) {
	ctx := context.Background()
	subject := new(model.SubjectTable)
	err := mysqlstore.Db.Table(model.TableNameSubject).
		WithContext(ctx).
		Where("name = ?", task.SubjectName).
		Where(comModel.DefaultDelCondition).
		Find(&subject).Error
	if err != nil {
		fmt.Println("get one subject by name error: ", err)
		return nil, err
	}

	return subject, nil
}
