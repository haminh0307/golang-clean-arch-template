package dependency

import (
	"github.com/haminh0307/golang-clean-arch-template/internal/delivery/restapi/middleware"

	"go.uber.org/fx"
)

var MiddlewareModule = fx.Options(
	fx.Provide(
		middleware.NewAuthentication,
	),
)
