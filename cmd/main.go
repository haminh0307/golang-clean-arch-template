package main

import (
	"rainbow-love-memory/internal/dependency"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		dependency.AppOptions,
	).Run()
}
