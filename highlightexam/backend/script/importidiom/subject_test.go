package main

import (
	"testing"

	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"github.com/stretchr/testify/assert"
)

func TestUpsertSubject(t *testing.T) {
	_, err := testdata.InitServer()
	assert.Nil(t, err)

	task := &Task{
		SubjectName:     "xxx",
		SubjectDescribe: "xxxxx",
	}
	subject, err := upsertSubject(task)
	assert.Nil(t, err)
	assert.Equal(t, subject.ID > 100, true)
}

func TestUpdateSubjectTotal(t *testing.T) {
	_, err := testdata.InitServer()
	assert.Nil(t, err)

	err = updateSubjectTotal(1)
	assert.Nil(t, err)
}
