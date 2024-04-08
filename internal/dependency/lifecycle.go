package dependency

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

var LifecycleModule = fx.Options(
	fx.Invoke(func(db *mongo.Database, lc fx.Lifecycle) {
		lc.Append(fx.StopHook(db.Client().Disconnect))
	}),
	fx.Invoke(func(server *http.Server, lc fx.Lifecycle) {
		lc.Append(
			fx.Hook{
				OnStart: func(context.Context) error {
					go func() {
						err := server.ListenAndServe()
						if err != nil {
							panic(err)
						}
					}()

					return nil
				},
				OnStop: func(ctx context.Context) error {
					return server.Shutdown(ctx)
				},
			},
		)
	}),
)
