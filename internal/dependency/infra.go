package dependency

import (
	"golang-clean-arch-template/internal/domain/interface/infra"
	"golang-clean-arch-template/pkg/jwtprovider"
	mongopkg "golang-clean-arch-template/pkg/mongo"

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
