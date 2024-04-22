package dependency

import (
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/interface/repository"
	"github.com/haminh0307/golang-clean-arch-template/internal/mongorepo"

	"go.uber.org/fx"
)

var RepositoryModule = fx.Options(
	fx.Provide(
		fx.Annotate(mongorepo.NewUser, fx.As(new(repository.User))),
	),
)
