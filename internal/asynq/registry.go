package asynqfx

import (
	authtask "github.com/mdtajulislammt/zvonsystem_backend/internal/modules/auth/task"
	"go.uber.org/fx"
)

var Registry = fx.Module("asynqregistryfx",
	fx.Provide(
		authtask.NewAuthProcessHandler,
	),
)
