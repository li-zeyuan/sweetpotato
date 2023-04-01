package service

import (
	"context"
	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/dao"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
)

var Profile = profileService{}

type profileService struct{}

func (l *profileService) Detail(ctx context.Context, uid int64) (*model.ProfileApiDetailResp, error) {
	userProfile, err := dao.D.User.GetOne(ctx, uid)
	if err != nil {
		return nil, err
	}

	resp := new(model.ProfileApiDetailResp)
	resp.UpdatedAt = utils.Time2TimeStamp(userProfile.UpdatedAt)
	resp.Uid = userProfile.ID
	resp.Name = userProfile.Name
	resp.Gender = userProfile.Gender
	resp.Portrait = userProfile.Portrait
	resp.CurrentSubjectId = userProfile.CurrentSubjectId
	resp.StudyTotalDay = userProfile.StudyTotalDay
	resp.StudyNum = userProfile.StudyNum

	return resp, nil
}

func (l *profileService) StudyNumEdit(ctx context.Context, uid int64, studyNum int) error {
	err := dao.D.User.Update(ctx, uid, map[string]interface{}{
		"study_num": studyNum,
	})
	if err != nil {
		return  err
	}

	return nil
}
