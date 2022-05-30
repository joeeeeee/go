package authzserver

import (
	"github.com/joe/iam/internal/authzserver/config"
	"github.com/joe/iam/internal/authzserver/options"
	"github.com/joe/iam/pkg/app"
)

const commandDesc = `auth`

func NewApp(basename string) *app.App {
	opts := options.NewOptions()

	application := app.NewApp("IAM AUTHZ Server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithRunFunc(run(opts)),
	)


	return application
}

func run(opts *options.Options) app.RunFunc {
	return func (basename string) error {
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}