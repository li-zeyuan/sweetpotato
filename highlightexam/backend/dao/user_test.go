package dao

import (
	"context"
	"fmt"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"testing"

	"github.com/li-zeyuan/common/sequence"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetByOpenid(t *testing.T) {
	ctx, err := testdata.InitServer()
	assert.Nil(t, err)

	u := User{}
	user, err := u.GetByOpenid(ctx, "open_id")
	assert.Equal(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, user)
}

func TestSave(t *testing.T) {
	err := InitDao()
	assert.Nil(t, err)

	userProfile := new(model.UserProfileTable)
	userProfile.ID = sequence.NewID()
	userProfile.Name = fmt.Sprintf("husband_%d", userProfile.ID%1000000)
	userProfile.Openid = "test"
	userProfile.SessionKey = "deef"

	u := User{}
	err = u.Save(context.Background(), []*model.UserProfileTable{userProfile})
	assert.Nil(t, err)
}


func TestStudyLastTime(t *testing.T) {
	ctx, err := testdata.InitServer()
	if err != nil {
		t.Fatal(err)
	}

	u := User{}
	lTime, err := u.StudyLastTime(ctx, 100000000000000000)
	assert.Nil(t, err)
	t.Log(lTime)
}