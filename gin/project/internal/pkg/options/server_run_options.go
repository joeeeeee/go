package options

import "github.com/joe/iam/internal/pkg/server"

type ServerRunOptions struct {
	Mode        string   `json:"mode"  mapstructure:"mode"`
	Healthz     bool     `json:"healthz"  mapstructure:"healthz"`
	Middlewares []string `json:"middlewares"  mapstructure:"middlewares"`
}

func NewServerRunOptions() *ServerRunOptions {
	defaults := server.NewConfig()

	return &ServerRunOptions{
		Mode:        defaults.Mode,
		Healthz:     defaults.Healthz,
		Middlewares: defaults.Middlewares,
	}
}

// ApplyTo applies the run options to the method receiver and returns self.
func (s *ServerRunOptions) ApplyTo(c *server.Config) error {
	c.Mode = s.Mode
	c.Healthz = s.Healthz
	c.Middlewares = s.Middlewares

	return nil
}
