package dependency

import (
	"golang-clean-arch-template/internal/delivery/restapi/handler"

	"go.uber.org/fx"
)

var HandlerModule = fx.Options(
	fx.Provide(
		handler.NewAuthentication,
		handler.NewUser,
	),
)
