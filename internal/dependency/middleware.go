package dependency

import (
	"rainbow-love-memory/internal/delivery/restapi/middleware"

	"go.uber.org/fx"
)

var MiddlewareModule = fx.Options(
	fx.Provide(
		middleware.NewAuthentication,
	),
)
