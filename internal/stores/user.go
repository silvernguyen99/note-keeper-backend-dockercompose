package stores

import (
	"note-keeper-backend/internal/models"

	"gorm.io/gorm"
)

type UserStore struct {
	*gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db}
}

func (m *UserStore) Save(object *models.ModelUser) error {
	object.BeforeCreate()
	return m.save(object)
}

func (m *UserStore) save(object *models.ModelUser) error {
	return m.Create(object).Error
}
func (m *UserStore) GetByAccessToken(accessToken string) (*models.ModelUser, bool, error) {
	var object = &models.ModelUser{}

	err := m.Model(models.ModelUser{}).Where("access_token = ?", accessToken).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *UserStore) GetByUserId(userId uint32) (*models.ModelUser, bool, error) {
	var object = &models.ModelUser{}

	err := m.Model(models.ModelUser{}).Where("user_id = ?", userId).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *UserStore) UpdateMapByUserID(userId uint32, mm map[string]interface{}) error {
	return m.Model(models.ModelUser{}).Where("user_id = ?", userId).Updates(mm).Error
}

func (m *UserStore) ClearAccessToken(userId uint32) error {
	mm := map[string]interface{}{
		"access_token": nil,
	}
	return m.Model(models.ModelUser{}).Where("user_id = ?", userId).Updates(mm).Error
}
