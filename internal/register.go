package internal

import (
	"github.com/thiruselvaa/yokai-worker-template/internal/worker"
	"github.com/ankorstore/yokai/fxworker"
	"go.uber.org/fx"
)

// Register is used to register the application dependencies.
func Register() fx.Option {
	return fx.Options(
		fxworker.AsWorker(worker.NewExampleWorker),
	)
}
