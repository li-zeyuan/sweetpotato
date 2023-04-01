package service

import (
	"context"
	"encoding/json"

	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/dao"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"go.uber.org/zap"
)

var Subject = subjectService{}

type subjectService struct{}

func (s *subjectService) List(ctx context.Context, uid int64) (*model.SubjectListResp, error) {
	var currentSubjectId int64
	if uid > 0 {
		userProfile, err := dao.D.User.GetOne(ctx ,uid)
		if err != nil {
			return nil, err
		}
		currentSubjectId = userProfile.CurrentSubjectId
	}


	subjects, err := dao.D.Subject.List(ctx)
	if err != nil {
		return nil, err
	}

	resp := new(model.SubjectListResp)
	resp.Studying = new(model.SubjectItem)
	resp.Others = make([]*model.SubjectItem, 0)
	for _, subject := range subjects {
		if subject.ID == currentSubjectId {
			resp.Studying.ID = subject.ID
			resp.Studying.UpdatedAt = utils.Time2TimeStamp(subject.UpdatedAt)
			resp.Studying.Name = subject.Name
			resp.Studying.Description = subject.Description
			resp.Studying.Total = subject.Total
		} else {
			resp.Others = append(resp.Others, &model.SubjectItem{
				ID:          subject.ID,
				UpdatedAt:   utils.Time2TimeStamp(subject.UpdatedAt),
				Name:        subject.Name,
				Description: subject.Description,
				Total:       subject.Total,
			})
		}
	}

	return resp, nil
}

func (s *subjectService) Detail(ctx context.Context, req *model.SubjectDetailReq) (*model.KnowledgeListResp, error) {
	knowledgeList, err := dao.D.Knowledge.ListBySubjectID(ctx ,req.ID, req.Start, req.Limit)
	if err != nil {
		return nil, err
	}

	resp := new(model.KnowledgeListResp)
	resp.List = make([]*model.KnowledgeListRespItem, 0)
	for _, k := range knowledgeList {
		other := model.OtherField{}
		err = json.Unmarshal([]byte(k.Other), &other)
		if err != nil {
			mylogger.Error(ctx, "json decode other field error: ", zap.Error(err))
			return nil, err
		}
		resp.List = append(resp.List, &model.KnowledgeListRespItem{
			ID:          k.ID,
			SubjectID:   k.SubjectID,
			Name:        k.Name,
			Description: k.Description,
			Other:       other,
		})
	}

	if len(resp.List) == req.Limit {
		resp.HasMore = true
	}

	return resp, nil
}
