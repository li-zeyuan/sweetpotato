package main

import (
	"testing"

	"github.com/li-zeyuan/sun/highlightexam/onrequest"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCreateKnowledge(t *testing.T) {
	_, err := testdata.InitServer()
	assert.Nil(t, err)

	idiomList := []*onrequest.Idiom{
		{Title: "title3", Description: "des", Pinyin: "pinyin"},
	}

	err = createKnowledge(1, idiomList)
	assert.Nil(t, err)
}
