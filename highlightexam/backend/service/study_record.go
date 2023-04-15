package service

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"gorm.io/gorm"

	comModel "github.com/li-zeyuan/common/model"
	"github.com/li-zeyuan/common/mylogger"
	"github.com/li-zeyuan/common/sequence"
	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/dao"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/model"
	"go.uber.org/zap"
)

var Study = studyService{}

type studyService struct{}

func (s *studyService) RecordList(ctx context.Context, uid int64, req *comModel.StartAndLimit) (*model.StudyRecordListResp, error) {
	studyRecords, err := dao.D.StudyRecord.ListByUid(ctx, uid, req.Start, req.Limit)
	if err != nil {
		return nil, err
	}

	resp := new(model.StudyRecordListResp)
	resp.List = make([]*model.StudyRecordItem, 0)
	if len(studyRecords) == 0 {
		return resp, nil
	}

	subjectIDs := make([]int64, 0, len(studyRecords))
	for _, sr := range studyRecords {
		subjectIDs = append(subjectIDs, sr.SubjectID)
	}

	subjectMap, err := dao.D.Subject.MapByIds(ctx, subjectIDs)
	if err != nil {
		return nil, err
	}

	for _, sr := range studyRecords {
		sub, ok := subjectMap[sr.SubjectID]
		if !ok {
			continue
		}

		resp.List = append(resp.List, &model.StudyRecordItem{
			SubjectID:   sub.ID,
			SubjectName: sub.Name,
			Studied:     sr.Count,
			Total:       sub.Total,
		})
	}

	if len(resp.List) == req.Limit {
		resp.HasMore = true
	}

	return resp, nil
}

func (s *studyService) KnowledgeList(ctx context.Context, uid int64, params *model.StudyKnowledgeReq) (*model.StudyKnowledgeResp, error) {
	hasStudiedKids, err := dao.D.StudyRecord.HasStudiedKids(ctx, uid, params.SubjectID)
	if err != nil {
		return nil, err
	}

	batchKnowledge, err := dao.D.Knowledge.WillStudiedList(ctx, params.SubjectID, hasStudiedKids)
	if err != nil {
		return nil, err
	}

	resp := new(model.StudyKnowledgeResp)
	resp.HasStudied = len(hasStudiedKids)
	resp.List = make([]*model.StudyKnowledgeItem, 0, model.DefaultStudyBatch)
	for _, k := range batchKnowledge {
		other := model.OtherField{}
		err = json.Unmarshal([]byte(k.Other), &other)
		if err != nil {
			mylogger.Error(ctx, "json decode other field error: ", zap.Error(err))
			return nil, err
		}

		resp.List = append(resp.List, &model.StudyKnowledgeItem{
			ID:          k.ID,
			Name:        k.Name,
			Description: k.Description,
			Other:       other,
		})
	}

	if len(resp.List) == model.DefaultStudyBatch {
		resp.HasMore = true
	}
	shuffleKnowledgeList(resp.List)
	return resp, nil
}

func shuffleKnowledgeList(list []*model.StudyKnowledgeItem) {
	if len(list) == 0 {
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}

func (s *studyService) Doing(ctx context.Context, uid int64, params *model.StudyDoingReq) (*model.StudyDoingResp, error) {
	record := new(model.StudyRecordTable)
	record.CreatedAt = utils.NowUTC()
	record.ID = sequence.NewID()
	record.Uid = uid
	record.SubjectID = params.SubjectID
	record.KnowledgeID = params.KnowledgeID

	err := dao.D.StudyRecord.Upsert(ctx, record)
	if err != nil {
		return nil, err
	}

	user, err := dao.D.User.GetOne(ctx, uid)
	if err != nil {
		return nil, err
	}

	fieldMap := make(map[string]interface{}, 2)
	fieldMap["study_last_time"] = utils.NowUTC()
	if user.StudyLastTime.Before(utils.TodayStartUTC()) {
		fieldMap["study_total_day"] = gorm.Expr("study_total_day + ?", 1)
	}

	err = dao.D.User.Update(ctx, uid, fieldMap)
	if err != nil {
		return nil, err
	}

	todayStudyNum, err := dao.D.StudyRecord.TodayStudyNum(ctx, uid)
	if err != nil {
		return nil, err
	}

	resp := new(model.StudyDoingResp)
	resp.IsCompletedToday = todayStudyNum >= int64(user.StudyNum)

	return resp, nil
}
