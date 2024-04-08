package dependency

import (
	"rainbow-love-memory/internal/delivery/restapi/handler"

	"go.uber.org/fx"
)

var HandlerModule = fx.Options(
	fx.Provide(
		handler.NewAuthentication,
		handler.NewUser,
	),
)
