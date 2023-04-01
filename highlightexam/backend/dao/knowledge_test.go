package dao

import (
	"testing"

	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"github.com/stretchr/testify/assert"
)

func TestListBySubjectID(t *testing.T) {
	ctx, err := testdata.InitServer()
	assert.Nil(t, err)

	k := Knowledge{}
	knowledgeList, err := k.ListBySubjectID(ctx, 1, 0, 2)
	assert.Nil(t, err)
	t.Log(len(knowledgeList))
}
