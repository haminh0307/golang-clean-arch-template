package dependency

import (
	"golang-clean-arch-template/internal/domain/interface/usecase"
	ucimpl "golang-clean-arch-template/internal/usecase"

	"go.uber.org/fx"
)

var UseCaseModule = fx.Options(
	fx.Provide(
		fx.Annotate(ucimpl.NewAuthentication, fx.As(new(usecase.Authentication)), fx.ParamTags("", "", `name:"JWT_EXPIRY"`)),
		fx.Annotate(ucimpl.NewUser, fx.As(new(usecase.User))),
	),
)
