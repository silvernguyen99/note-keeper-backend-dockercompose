package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"note-keeper-backend/internal/models"
	"strconv"
)

func (s *Service) GetNote(w http.ResponseWriter, req *http.Request) {
	s.enableCors(&w)
	var err error
	res := models.GetNoteResponse{}

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error while Marshaling GetNoteResponse: " + err.Error())
			return
		}

		w.Write(resJson)
	}()

	userId := req.Header.Get("user-id")
	userIdUint, _ := strconv.ParseUint(userId, 10, 32)

	var r models.GetNoteRequest
	err = json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("error while decoding GetNoteRequest: " + err.Error())
		return
	}

	note, existed, err := s.noteStore.GetNote(r.NoteId, uint32(userIdUint))
	if err != nil {
		fmt.Println("error while getting note: " + err.Error())
		return
	}

	if !existed {
		res = models.GetNoteResponse{
			FlagInfo: models.FlagInfo{
				Flag:    144,
				Message: "Note does not existed or you are not allow to view this note",
			},
		}
		return
	}

	res.NoteData = models.Note{
		NoteId:   note.NoteId,
		NoteName: note.NoteName,
		Content:  note.Content,
	}

	res.FlagInfo = models.FlagInfo{
		Flag:    143,
		Message: "OK",
	}

	return
}

func (s *Service) GetAllNotes(w http.ResponseWriter, req *http.Request) {
	s.enableCors(&w)
	var err error
	var res models.GetAllNotesResponse

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error while Marshaling GetNoteResponse: " + err.Error())
			return
		}

		w.Write(resJson)
	}()

	userId := req.Header.Get("user-id")
	userIdUint, _ := strconv.ParseUint(userId, 10, 32)
	fmt.Printf("userIdUint: %v", userIdUint)

	notes, existed, err := s.noteStore.GetNotesByUserId(uint32(userIdUint))
	if err != nil {
		fmt.Println("error while performing GetNotesByUserId: " + err.Error())
		return
	}

	// to notes reponse
	if existed {
		notesRes := make([]models.Note, len(notes))
		for i := 0; i < len(notes); i++ {
			notesRes[i].NoteId = notes[i].NoteId
			notesRes[i].NoteName = notes[i].NoteName
			notesRes[i].Content = notes[i].Content
		}

		res = models.GetAllNotesResponse{
			NotesData: notesRes,
			FlagInfo: models.FlagInfo{
				Flag:    143,
				Message: "OK",
			},
		}

		return
	} else {
		res.FlagInfo = models.FlagInfo{
			Flag:    143,
			Message: "no note",
		}
	}
}

func (s *Service) EditNote(w http.ResponseWriter, req *http.Request) {
	s.enableCors(&w)
	var err error
	var res models.EditNoteResponse

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error while Marshaling GetNoteResponse: " + err.Error())
			return
		}

		w.Write(resJson)
	}()

	userId := req.Header.Get("user-id")
	userIdUint, _ := strconv.ParseUint(userId, 10, 32)

	var r models.EditNoteRequest
	err = json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("error while decoding EditNoteRequest: " + err.Error())
		return
	}

	var note models.ModelNote
	note.NoteName = r.NoteData.NoteName
	note.Content = r.NoteData.Content

	if err = s.noteStore.UpdateNote(r.NoteData.NoteId, uint32(userIdUint), &note); err != nil {
		fmt.Println("error while editing note: " + err.Error())
		return
	}

	res.FlagInfo = models.FlagInfo{
		Flag:    143,
		Message: "Edit note successfully!!",
	}
}

func (s *Service) DeleteNote(w http.ResponseWriter, req *http.Request) {
	s.enableCors(&w)
	var err error
	var res models.DeleteNoteResponse

	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error while Marshaling GetNoteResponse: " + err.Error())
			return
		}

		w.Write(resJson)
	}()

	// get user-id
	userId := req.Header.Get("user-id")
	userIdUint, _ := strconv.ParseUint(userId, 10, 32)

	var r models.DeleteNoteRequest
	err = json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("error while decoding DeleteNoteRequest: " + err.Error())
		return
	}

	// delete
	if err = s.noteStore.DeleteNote(r.NoteId, uint32(userIdUint)); err != nil {
		fmt.Println("error while deleting note: " + err.Error())
		return
	}

	res.FlagInfo = models.FlagInfo{
		Flag:    143,
		Message: "delete note successfully",
	}
}

func (s *Service) CreateNote(w http.ResponseWriter, req *http.Request) {
	s.enableCors(&w)
	var err error
	var res models.CreateNoteResponse
	defer func() {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			fmt.Println("error while Marshaling GetNoteResponse: " + err.Error())
			return
		}

		w.Write(resJson)
	}()

	// get user-id
	userId := req.Header.Get("user-id")
	userIdUint, _ := strconv.ParseUint(userId, 10, 32)

	var r models.CreateNoteRequest
	err = json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		fmt.Println("error while decoding DeleteNoteRequest: " + err.Error())
		return
	}

	// create note
	newNote := &models.ModelNote{
		UserId:   uint32(userIdUint),
		NoteName: r.NoteData.NoteName,
		Content:  r.NoteData.Content,
	}

	if err = s.noteStore.Save(newNote); err != nil {
		fmt.Println("error while creating new note: " + err.Error())
		return
	}

	res.FlagInfo = models.FlagInfo{
		Flag:    143,
		Message: "Create note successfully!!",
	}
}
