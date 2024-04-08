package dependency

import (
	"rainbow-love-memory/internal/domain/interface/repository"
	"rainbow-love-memory/internal/mongorepo"

	"go.uber.org/fx"
)

var RepositoryModule = fx.Options(
	fx.Provide(
		fx.Annotate(mongorepo.NewUser, fx.As(new(repository.User))),
	),
)
