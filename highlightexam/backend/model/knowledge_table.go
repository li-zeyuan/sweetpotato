package model

import "github.com/li-zeyuan/common/model"

const (
	TableNameKnowledge = "knowledge"
	DefaultStudyBatch = 100
)

// 知识点表
type KnowledgeTable struct {
	model.BaseModel
	SubjectID   int64  // 题库id
	Name        string // 名称
	Description string // 描述
	Other       string // 其他字段；json
}

type OtherField struct {
	Pinyin string `json:"pinyin"`
}
