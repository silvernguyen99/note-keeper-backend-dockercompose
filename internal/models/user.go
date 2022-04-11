package models

type ModelUser struct {
	UserId      uint32 `json:"user_id" gorm:"primaryKey"`
	AccessToken string `json:"access_token"`
	UserName    string `json:"user_name"`
}

func (m *ModelUser) BeforeCreate() {

}

func (m *ModelUser) TableName() string {
	return "tb_users"
}
