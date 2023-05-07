package main

import (
	"context"
	"encoding/json"
	"fmt"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mysqlstore"
	"github.com/li-zeyuan/common/sequence"
	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sun/highlightexam/onrequest"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
)

func getKnowledge(idioms []string, subjectID int64) ([]string, error) {
	ctx := context.Background()
	existIdioms := make([]string, 0)
	err := mysqlstore.Db.Table(model.TableNameKnowledge).
		WithContext(ctx).
		Where("subject_id = ?", subjectID).
		Where(comModel.DefaultDelCondition).
		Where("name in ?", idioms).
		Scan(&existIdioms).Error
	if err != nil {
		fmt.Println("get will study knowledge error: ", err)
		return nil, err
	}

	return utils.ExcludeSlice(idioms, existIdioms), nil
}

func createKnowledge(subjectID int64, idiomList []*onrequest.Idiom) error {
	kModels := make([]*model.KnowledgeTable, 0, len(idiomList))
	for _, idi := range idiomList {
		otherField := model.OtherField{
			Pinyin: idi.Pinyin,
		}

		jsonB, err := json.Marshal(otherField)
		if err != nil {
			fmt.Println("json marshal other field error: ", err)
			return err
		}

		kModel := &model.KnowledgeTable{
			BaseModel: comModel.BaseModel{
				ID: sequence.NewID(),
			},
			SubjectID:   subjectID,
			Name:        idi.Title,
			Description: idi.Description,
			Other:       string(jsonB),
		}

		kModels = append(kModels, kModel)
	}

	ctx := context.Background()
	err := mysqlstore.Db.Table(model.TableNameKnowledge).
		WithContext(ctx).
		Create(&kModels).Error
	if err != nil {
		fmt.Println("create knowledge error: ", err)
		return err
	}

	return nil
}
