package app

import (
	"github.com/mdtajulislammt/zvonsystem_backend/internal/config"
	asynqfx "github.com/mdtajulislammt/zvonsystem_backend/internal/asynq"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/modules/auth"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/modules/user"
	"github.com/mdtajulislammt/zvonsystem_backend/internal/repository"
	"github.com/mdtajulislammt/zvonsystem_backend/pkg/ORM"
	"github.com/mdtajulislammt/zvonsystem_backend/pkg/logger"
	"github.com/mdtajulislammt/zvonsystem_backend/pkg/redis"
	"github.com/mdtajulislammt/zvonsystem_backend/pkg/s3client"
	"go.uber.org/fx"
)

func BaseModules() fx.Option {
	return fx.Options(
		// infra
		config.Module,
		logger.Module,
		redis.Module,
		// ORM.Init,
		s3client.Module,

		// async / integrations
		asynqfx.Module,

		// data + domain
		repository.Module,

		// business modules
		auth.Module,
		user.Module,

		fx.Provide(
			ORM.Init,
		),
	)
}
