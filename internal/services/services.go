package services

import (
	"fmt"
	"log"
	"net/http"
	"note-keeper-backend/config"
	"note-keeper-backend/internal/stores"
)

type Service struct {
	cfg              *config.Config
	mainStore        *stores.MainStore
	loginSocialStore *stores.LoginSocialStore
	userStore        *stores.UserStore
	noteStore        *stores.NoteStore
}

func New(
	config *config.Config,
	mainStore *stores.MainStore,
	loginSocialStore *stores.LoginSocialStore,
	userStore *stores.UserStore,
	noteStore *stores.NoteStore,
) *Service {
	s := &Service{
		cfg:              config,
		mainStore:        mainStore,
		loginSocialStore: loginSocialStore,
		userStore:        userStore,
		noteStore:        noteStore,
	}

	return s
}

func (s *Service) enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
}

func (s *Service) Run() {
	http.HandleFunc("/login_using_social_access_token", s.LoginUsingSocialAccessToken)

	// basic middleware
	logoutHandler := http.HandlerFunc(s.Logout)
	http.Handle("/logout", s.AuthHandler(logoutHandler))

	// get all notes api
	getAllNotesHandler := http.HandlerFunc(s.GetAllNotes)
	http.Handle("/me/get_all_notes", s.AuthHandler(getAllNotesHandler))

	// get note detail api
	getNoteHandler := http.HandlerFunc(s.GetNote)
	http.Handle("/me/get_note", s.AuthHandler(getNoteHandler))

	// create note api
	createNoteHandler := http.HandlerFunc(s.CreateNote)
	http.Handle("/me/create_note", s.AuthHandler(createNoteHandler))

	// edit note api
	editNoteHandler := http.HandlerFunc(s.EditNote)
	http.Handle("/me/edit_note", s.AuthHandler(editNoteHandler))

	// delete note api
	deleteNoteHandler := http.HandlerFunc(s.DeleteNote)
	http.Handle("/me/delete_note", s.AuthHandler(deleteNoteHandler))

	log.Println("Listening on :9000...")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Service) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.enableCors(&w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		token := r.Header.Get("Authorization")
		if token == "" {
			fmt.Println("Access token is empty")
			http.Error(w, http.StatusText(401), 401)
			return
		}

		fmt.Print("token " + token + "\n")

		userModel, existed, err := s.userStore.GetByAccessToken(token)
		if err != nil {
			fmt.Println("error while authentication user: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if !existed {
			fmt.Println("Invalid token")
			http.Error(w, http.StatusText(401), 401)
			return
		}

		r.Header.Set("user-id", fmt.Sprint(userModel.UserId))

		next.ServeHTTP(w, r)
	})
}
