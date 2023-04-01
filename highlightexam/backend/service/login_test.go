package service

import (
	"testing"

	"github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"github.com/stretchr/testify/assert"
)

func TestWeChatLogin(t *testing.T) {
	ctx, err := testdata.InitServer()
	assert.Nil(t, err)

	req := &model.WeChatLoginReq{
		Code: "063E6lFa1APAVE0wfBIa12tXBP1E6lFF",
	}

	l := loginService{}
	token, err := l.WeChatLogin(ctx, req)
	assert.Nil(t, err)
	assert.NotEqual(t, len(token), 0)
}
