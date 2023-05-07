package main

import (
	"context"
	"fmt"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/common/sequence"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
)

func upsertSubject(task *Task) (*model.SubjectTable, error) {
	subject, err := getSubject(task)
	if err != nil {
		return nil, err
	}

	if subject.ID > 0 {
		return subject, nil
	}

	subject = &model.SubjectTable{
		BaseModel: comModel.BaseModel{
			ID: sequence.NewID(),
		},
		Name:        task.SubjectName,
		Description: task.SubjectDescribe,
	}

	err = mysqlstore.Db.Table(model.TableNameSubject).
		Create(subject).Error
	if err != nil {
		fmt.Println("create one subject by name error: ", err)
		return nil, err
	}

	return subject, nil
}

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

func updateSubjectTotal(subjectID int64) error {
	var kNum int64
	err := mysqlstore.Db.Table(model.TableNameKnowledge).
		Where("subject_id = ?", subjectID).
		Where(comModel.DefaultDelCondition).
		Count(&kNum).Error
	if err != nil {
		fmt.Println("get knowledge num by name error: ", err)
		return err
	}

	err = mysqlstore.Db.Table(model.TableNameSubject).
		Where("id = ?", subjectID).
		Where(comModel.DefaultDelCondition).
		Update("total", kNum).Error
	if err != nil {
		fmt.Println("update subject num by name error: ", err)
		return err
	}

	return nil
}
