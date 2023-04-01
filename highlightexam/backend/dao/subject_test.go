package dao

import (
	"testing"

	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"github.com/stretchr/testify/assert"
)

func TestSubjectList(t *testing.T) {
	ctx, err := testdata.InitServer()
	assert.Nil(t, err)

	s := Subject{}
	subjects, err := s.List(ctx)
	assert.Nil(t, err)
	t.Log(len(subjects))
}
