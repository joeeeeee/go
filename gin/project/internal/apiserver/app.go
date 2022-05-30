package apiserver

import (
	"github.com/joe/iam/internal/apiserver/config"
	"github.com/joe/iam/internal/apiserver/options"
	"github.com/joe/iam/pkg/app"
)

const CommandDesc = `iam description`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()

	application := app.NewApp("IAM API Server",
		basename,
		app.WithRunFunc(run(opts)),
		app.WithDescription(CommandDesc),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
