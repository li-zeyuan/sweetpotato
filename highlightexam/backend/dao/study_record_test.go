package dao

import (
	"testing"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"github.com/stretchr/testify/assert"
)

func TestUpsert(t *testing.T) {
	ctx, err := testdata.InitServer()
	assert.Nil(t, err)

	cases := []struct {
		Name  string
		Model *model.StudyRecordTable
	}{
		{"1",
			&model.StudyRecordTable{
				BaseModel:   comModel.BaseModel{ID: 1, CreatedAt: utils.NowUTC()},
				Uid:         100000000000000000,
				SubjectID:   1,
				KnowledgeID: 1,
			},
		},
	}

	s := StudyRecord{}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			err = s.Upsert(ctx, c.Model)
			assert.Nil(t, err)
		})
	}
}

func TestTodayStudyNum(t *testing.T) {
	ctx, err := testdata.InitServer()
	assert.Nil(t, err)

	s := StudyRecord{}
	num, err := s.TodayStudyNum(ctx, 100000000000000000)
	assert.Nil(t, err)
	assert.Equal(t, num > 0, true)
}
