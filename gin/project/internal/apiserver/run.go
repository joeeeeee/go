package apiserver

import "github.com/joe/iam/internal/apiserver/config"

func Run(cfg *config.Config) error {
	server, err := CreateApiServer(cfg)
	if err != nil {
		return err
	}

	return server.PrepareRun()
}
