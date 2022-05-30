package authzserver

import "github.com/joe/iam/internal/authzserver/config"

func Run(cfg *config.Config) error {
	server, err := createAuthzServer(cfg)
	if err != nil {
		return err
	}

	return server.Run()

}