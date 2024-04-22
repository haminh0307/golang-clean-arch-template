package dependency

import (
	"golang-clean-arch-template/internal/delivery/restapi"
	"golang-clean-arch-template/internal/delivery/restapi/controller"

	"go.uber.org/fx"
)

var ControllerModule = fx.Options(
	fx.Provide(
		fx.Annotate(controller.NewAuthentication, fx.As(new(restapi.Controller)), fx.ResultTags(`group:"controllers"`)),
		fx.Annotate(controller.NewUser, fx.As(new(restapi.Controller)), fx.ResultTags(`group:"controllers"`)),
	),
)
