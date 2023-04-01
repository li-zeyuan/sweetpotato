package model

type SubjectListResp struct {
	Studying *SubjectItem   `json:"studying"`
	Others   []*SubjectItem `json:"others"`
}

type SubjectItem struct {
	ID          int64  `json:"id"`
	UpdatedAt   int64  `json:"updated_at"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Total       int    `json:"total"`
}

type SubjectDetailReq struct {
	ID    int64 `form:"id" validate:"required"`
	Start int   `form:"start"`
	Limit int   `form:"limit"`
}
