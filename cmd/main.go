package main

import (
	"golang-clean-arch-template/internal/dependency"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		dependency.AppOptions,
	).Run()
}
