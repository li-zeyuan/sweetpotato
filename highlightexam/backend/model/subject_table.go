package model

import "github.com/li-zeyuan/common/model"

const (
	TableNameSubject = "subject"
)

// 题库表
type SubjectTable struct {
	model.BaseModel
	Name        string // 题库名称
	Description string // 描述
	Total       int    // 题库总数量
}
