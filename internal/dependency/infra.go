package dependency

import (
	"rainbow-love-memory/internal/domain/interface/infra"
	"rainbow-love-memory/pkg/jwtprovider"
	mongopkg "rainbow-love-memory/pkg/mongo"

	"go.uber.org/fx"
)

var InfraModule = fx.Options(
	fx.Provide(
		fx.Annotate(
			jwtprovider.NewJwtProvider,
			fx.As(new(infra.JwtProvider)),
			fx.ParamTags(`name:"JWT_ALG"`, `name:"JWT_KEY"`),
		),
		fx.Annotate(
			mongopkg.New,
			fx.ParamTags("", `name:"MONGO_CONN_URI"`, `name:"MONGO_DB_NAME"`),
		),
	),
)
