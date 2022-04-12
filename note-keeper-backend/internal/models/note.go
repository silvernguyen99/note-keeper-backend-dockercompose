package models

import "time"

type ModelNote struct {
	UserId    uint32     `json:"user_id"`
	NoteId    uint32     `json:"note_id" gorm:"primaryKey"`
	NoteName  string     `json:"note_name"`
	Content   string     `json:"content"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (m *ModelNote) BeforeCreate() {

}

func (m *ModelNote) TableName() string {
	return "tb_notes"
}
