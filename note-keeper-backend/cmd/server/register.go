package main

import (
	"note-keeper-backend/config"
	"note-keeper-backend/internal/services"
	"note-keeper-backend/internal/stores"
)

func registerService(cfg *config.Config) *services.Service {
	db := mustConnectPostgres(cfg)

	mainStore := stores.NewMainStore(db)
	userStore := stores.NewUserStore(db)
	loginSocialStore := stores.NewLoginSocialStore(db)
	noteStore := stores.NewNoteStore(db)
	return services.New(cfg, mainStore, loginSocialStore, userStore, noteStore)
}
