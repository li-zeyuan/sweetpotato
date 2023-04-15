package model

type StudyRecordListResp struct {
	List    []*StudyRecordItem `json:"list"`
	HasMore bool               `json:"has_more"`
}

type StudyRecordItem struct {
	SubjectID   int64  `json:"subject_id"`
	SubjectName string `json:"subject_name"`
	Studied     int    `json:"studied"`
	Total       int    `json:"total"`
}

type StudyRecord struct {
	SubjectID int64
	Count     int
}

type StudyKnowledgeResp struct {
	List       []*StudyKnowledgeItem `json:"list"`
	HasStudied int                   `json:"has_studied"`
	HasMore    bool                  `json:"has_more"`
}

type StudyKnowledgeItem struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Other       OtherField `json:"other"`
}

type StudyKnowledgeReq struct {
	SubjectID int64 `form:"subject_id" validate:"gt=0"`
}

type StudyDoingReq struct {
	SubjectID   int64 `json:"subject_id" validate:"gt=0"`
	KnowledgeID int64 `json:"knowledge_id" validate:"gt=0"`
}

type StudyDoingResp struct {
	IsCompletedToday bool `json:"is_completed_today"`
}
