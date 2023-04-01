package dao

import (
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/config"
)

var D *Dao

type Dao struct {
	cfg       *config.Config
	User      *User
	Subject   *Subject
	Knowledge *Knowledge
	StudyRecord *StudyRecord
}

func New(cfg *config.Config) {
	D = &Dao{
		cfg:       cfg,
		User:      &User{},
		Subject:   &Subject{},
		Knowledge: &Knowledge{},
		StudyRecord: &StudyRecord{},
	}
}
