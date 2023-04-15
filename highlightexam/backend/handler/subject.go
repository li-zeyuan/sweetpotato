package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common/httptransfer"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/service"
)

var SubjectHandler = subject{}

type subject struct{}

func (s *subject) List(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	subjects, err := service.Subject.List(c.Request.Context(), uid)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, subjects)
}

func (s *subject) Detail(c *gin.Context) {
	req := new(model.SubjectDetailReq)
	err := httptransfer.ParseQuery(c, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	resp, err := service.Subject.Detail(c.Request.Context(), req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, resp)
}

func (s *subject) Study(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	req := new(model.SubjectStudyReq)
	err := httptransfer.ParseBody(c, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	err = service.Subject.Study(c.Request.Context(), uid, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, struct{}{})
}
