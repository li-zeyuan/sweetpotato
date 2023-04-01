package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/service"
)

var UserHandler = new(userProfile)

type userProfile struct{}

func (u *userProfile) Detail(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	profile, err := service.Profile.Detail(c.Request.Context(), uid)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, profile)
}

func (u *userProfile) StudyNumEdit(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	req := new(model.StudyNumEditReq)
	err := httptransfer.ParseBody(c, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	err = service.Profile.StudyNumEdit(c.Request.Context(), uid, req.Num)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, struct{}{})
}
