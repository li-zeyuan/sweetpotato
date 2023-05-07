package service

import (
	"context"
	"fmt"
	"time"

	"github.com/li-zeyuan/common/external"
	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/sequence"
	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/config"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/dao"
	model2 "github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
)

var Login = loginService{}

type loginService struct{}

func (l *loginService) WeChatLogin(ctx context.Context, req *model.WeChatLoginReq) (string, error) {
	wx := external.NewWechat(config.AppCfg.WxAppId, config.AppCfg.WxSecret)
	wxSession, err := wx.QueryWxSession(ctx, req.Code)
	if err != nil {
		return "", err
	}

	userProfile, err := dao.D.User.GetByOpenid(ctx, wxSession.OpenId)
	if err != nil {
		return "", err
	}
	if userProfile != nil && userProfile.DeletedAt != 0 {
		return "", httptransfer.ErrorLoginForbid
	}
	if userProfile.ID == 0 {
		userProfile = new(model2.UserProfileTable)
		userProfile.ID = sequence.NewID()
		userProfile.Name = fmt.Sprintf("official_%d", userProfile.ID%1000)
		userProfile.Openid = wxSession.OpenId
		userProfile.SessionKey = wxSession.SessionKey
		userProfile.StudyNum = model2.DefaultStudyNum
		userProfile.StudyLastTime = time.Date(1970, 1, 1, 0, 0, 1, 0, time.UTC)
		err = dao.D.User.Save(ctx, []*model2.UserProfileTable{userProfile})
		if err != nil {
			return "", err
		}
	}

	// todo redis token
	token, err := utils.GenerateToken(userProfile.ID, time.Hour*24*30, config.AppCfg.JwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
