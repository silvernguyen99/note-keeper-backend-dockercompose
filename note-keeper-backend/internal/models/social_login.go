package models

type ModelLoginSocial struct {
	UserId       uint32 `json:"user_id"`
	TypeProvider string `json:"type_provider"`
	SocialId     string `json:"social_id"`
}

func init() {

}

func (m *ModelLoginSocial) BeforeCreate() {

}

func (m ModelLoginSocial) TableName() string {
	return "tb_login_social"
}
