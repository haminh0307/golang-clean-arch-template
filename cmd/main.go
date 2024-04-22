package main

import (
	"github.com/haminh0307/golang-clean-arch-template/internal/dependency"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		dependency.AppOptions,
	).Run()
}
