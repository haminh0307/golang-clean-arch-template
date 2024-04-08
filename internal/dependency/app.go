package dependency

import (
	"context"
	"rainbow-love-memory/config"
	"rainbow-love-memory/internal/delivery/restapi"

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
