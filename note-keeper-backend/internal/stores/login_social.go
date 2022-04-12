package stores

import (
	"note-keeper-backend/internal/models"

	"gorm.io/gorm"
)

type LoginSocialStore struct {
	*gorm.DB
}

// for a service
func NewLoginSocialStore(db *gorm.DB) *LoginSocialStore {
	return &LoginSocialStore{db}
}

func (m *LoginSocialStore) Save(object *models.ModelLoginSocial) error {
	object.BeforeCreate()
	return m.save(object)
}

func (m *LoginSocialStore) save(object *models.ModelLoginSocial) error {
	return m.Create(object).Error
}

func (m *LoginSocialStore) GetBySocialId(socialId string) (*models.ModelLoginSocial, bool, error) {
	var object = &models.ModelLoginSocial{}
	err := m.Model(models.ModelLoginSocial{}).Where("social_id = ?", socialId).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}
