package asynqfx

import (
	"github.com/hibiken/asynq"
	authtask "github.com/mdtajulislammt/zvonsystem_backend/internal/modules/auth/task" 
)

func NewMux(
	authProcessHander *authtask.AuthProcessHandler,
) *asynq.ServeMux {
	mux := asynq.NewServeMux()
	// add worker handler here
	mux.HandleFunc(authtask.TypeAuthProcess, authProcessHander.Handle)

	return mux
}
