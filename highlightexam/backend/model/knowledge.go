package model

type KnowledgeListResp struct {
	List    []*KnowledgeListRespItem `json:"list"`
	HasMore bool                     `json:"has_more"`
}

type KnowledgeListRespItem struct {
	ID          int64      `json:"id"`
	SubjectID   int64      `json:"subject_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Other       OtherField `json:"other"`
}
