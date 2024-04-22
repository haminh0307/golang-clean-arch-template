package dependency

import (
	"context"
	"golang-clean-arch-template/config"
	"golang-clean-arch-template/internal/delivery/restapi"

	"go.uber.org/fx"
)

var AppOptions = fx.Options(
	LifecycleModule,
	fx.Provide(config.NewConfig),
	fx.Provide(fx.Annotate(restapi.NewServer, fx.ParamTags(`name:"HTTP_HOST"`, `name:"HTTP_PORT"`, `group:"controllers"`))),
	ControllerModule,
	HandlerModule,
	MiddlewareModule,
	UseCaseModule,
	RepositoryModule,
	InfraModule,
	fx.Provide(context.Background),
)
