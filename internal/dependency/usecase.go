package dependency

import (
	"rainbow-love-memory/internal/domain/interface/usecase"
	ucimpl "rainbow-love-memory/internal/usecase"

	"go.uber.org/fx"
)

var UseCaseModule = fx.Options(
	fx.Provide(
		fx.Annotate(ucimpl.NewAuthentication, fx.As(new(usecase.Authentication)), fx.ParamTags("", "", `name:"JWT_EXPIRY"`)),
		fx.Annotate(ucimpl.NewUser, fx.As(new(usecase.User))),
	),
)
