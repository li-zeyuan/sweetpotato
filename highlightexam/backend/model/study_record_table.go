package model

import (
	"github.com/li-zeyuan/common/model"
)

const (
	TableNameStudyRecord = "study_record"
)

// 学习记录表
type StudyRecordTable struct {
	model.BaseModel
	Uid         int64
	SubjectID   int64  // 题库id
	KnowledgeID int64 // 知识点id
}
