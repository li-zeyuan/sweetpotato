package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/li-zeyuan/common/httptransfer"
	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/service"
)

var StudyHandler = study{}

type study struct{}

func (s *study) RecordList(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	req := new(comModel.StartAndLimit)
	err := httptransfer.ParseQuery(c, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	resp, err := service.Study.RecordList(c.Request.Context(), uid, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, resp)
}

func (s *study) KnowledgeList(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	req := new(model.StudyKnowledgeReq)
	err := httptransfer.ParseQuery(c, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	resp, err := service.Study.KnowledgeList(c.Request.Context(), uid, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, resp)
}

func (s *study) Doing(c *gin.Context) {
	uid := httptransfer.GetUid(c)
	req := new(model.StudyDoingReq)
	err := httptransfer.ParseBody(c, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	resp, err := service.Study.Doing(c.Request.Context(), uid, req)
	if err != nil {
		httptransfer.ErrJSONResp(c, http.StatusInternalServerError, err)
		return
	}

	httptransfer.SuccJSONResp(c, resp)
}
