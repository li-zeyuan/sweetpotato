package model

type ProfileApiDetailResp struct {
	UpdatedAt        int64  `json:"updated_at"`
	Uid              int64  `json:"uid"`      // 用户ID
	Name             string `json:"name"`     // 用户昵称
	Gender           int    `json:"gender"`   // 性别 1-男；2-女
	Portrait         string `json:"portrait"` // 头像
	CurrentSubjectId int64  `json:"current_subject_id"`
	StudyTotalDay    int    `json:"study_total_day"`
	StudyNum         int    `json:"study_num"`
}

type StudyNumEditReq struct {
	Num int `json:"num" validate:"required"`
}
