package models

type FlagInfo struct {
	Flag    uint32 `json:"flag"`
	Message string `json:"message"`
}

type LoginUsingSocialAccessTokenRequest struct {
	Type        string `json:"type"`
	SocialToken string `json:"social_token"`
}

type Account struct {
	AccessToken string `json:"access_token"`
	UserName    string `json:"user_name"`
}

type LoginUsingSocialAccessTokenResponse struct {
	Account  Account  `json:"account"`
	FlagInfo FlagInfo `json:"flag_info"`
}

type LogoutResponse struct {
	FlagInfo FlagInfo `json:"flag_info"`
}

type Note struct {
	NoteId    uint32 `json:"note_id"`
	NoteName  string `json:"note_name"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

// get note
type GetNoteRequest struct {
	NoteId uint32 `json:"note_id"`
}

type GetNoteResponse struct {
	NoteData Note     `json:"note_data"`
	FlagInfo FlagInfo `json:"flag_info"`
}

// GetAllNotesResponse
type GetAllNotesResponse struct {
	NotesData []Note   `json:"notes_data"`
	FlagInfo  FlagInfo `json:"flag_info"`
}

// Create note
type CreateNoteRequest struct {
	NoteData Note `json:"note_data"`
}

type CreateNoteResponse struct {
	FlagInfo FlagInfo `json:"flag_info"`
}

// edit note
type EditNoteRequest struct {
	NoteData Note `json:"note_data"`
}

type EditNoteResponse struct {
	FlagInfo FlagInfo `json:"flag_info"`
}

// delete note

type DeleteNoteRequest struct {
	NoteId uint32 `json:"note_id"`
}

type DeleteNoteResponse struct {
	FlagInfo FlagInfo `json:"flag_info"`
}
