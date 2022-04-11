package stores

import (
	"note-keeper-backend/internal/models"

	"gorm.io/gorm"
)

type NoteStore struct {
	*gorm.DB
}

func NewNoteStore(db *gorm.DB) *NoteStore {
	return &NoteStore{db}
}

func (m *NoteStore) Save(object *models.ModelNote) error {
	object.BeforeCreate()
	return m.save(object)
}

func (m *NoteStore) save(object *models.ModelNote) error {
	return m.Create(object).Error
}

func (m *NoteStore) GetNotesByUserId(userId uint32) ([]*models.ModelNote, bool, error) {
	var notes = make([]*models.ModelNote, 0)
	err := m.Model(models.ModelNote{}).Where("user_id = ?", userId).Order("note_id asc").Find(&notes).Error
	if err == gorm.ErrRecordNotFound || len(notes) == 0 {
		return notes, false, nil
	}
	return notes, true, err
}

func (m *NoteStore) GetNote(noteId uint32, userId uint32) (*models.ModelNote, bool, error) {
	var object = &models.ModelNote{}
	err := m.Model(models.ModelNote{}).Where("note_id = ? and user_id = ?", noteId, userId).First(object).Error
	if err == gorm.ErrRecordNotFound {
		return object, false, nil
	}
	return object, true, err
}

func (m *NoteStore) UpdateNote(noteId uint32, userId uint32, object *models.ModelNote) error {
	return m.Model(models.ModelNote{}).Where("note_id = ? and user_id = ?", noteId, userId).Updates(&object).Error
}

func (m *NoteStore) DeleteNote(noteId uint32, userId uint32) error {
	return m.Model(models.ModelNote{}).Where("note_id = ? and user_id = ?", noteId, userId).Delete(&models.ModelNote{}).Error
}
