package dependency

import (
	"github.com/haminh0307/golang-clean-arch-template/internal/domain/interface/infra"
	"github.com/haminh0307/golang-clean-arch-template/pkg/jwtprovider"
	mongopkg "github.com/haminh0307/golang-clean-arch-template/pkg/mongo"

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
