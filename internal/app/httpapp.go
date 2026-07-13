package app

import (
	"github.com/mdtajulislammt/zvonsystem_backend/internal/modules/auth"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/modules/user"
	"go.uber.org/fx"
)

func BaseHTTPModules() fx.Option {
	return fx.Options(
		auth.HTTPModule,
		user.HTTPModule,
	)
}
