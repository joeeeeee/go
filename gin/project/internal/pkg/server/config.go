package server

import (
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
	"time"
)

// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*Config
}


type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}


type JwtInfo struct {
	Realm string

	Key string

	Timeout time.Duration

	maxRefresh time.Duration
}

func NewConfig() *Config {
	return &Config{
		SecureServing:   nil,
		InsecureServing: nil,
		Jwt:             &JwtInfo{
			Realm:      "iam jwt",
			Key:        "",
			Timeout:    1 * time.Hour,
			maxRefresh: 1 * time.Hour,
		},
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		Healthz:         true,
		EnableProfiling: true,
		EnableMetrics:   true,
	}
}

// CertKey contains configuration items related to certificate.
type CertKey struct {
	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}


// SecureServingInfo holds configuration of the TLS server.
type SecureServingInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey
}



// Address join host IP address and host port number into a address string, like: 0.0.0.0:8443.
func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

// InsecureServingInfo holds configuration of the insecure http server.
type InsecureServingInfo struct {
	Address string
}

func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// New returns a new instance of GenericAPIServer from the given config.
func (c CompletedConfig) New() (*GenericAPIServer, error) {
	s := &GenericAPIServer{
		SecureServingInfo:   c.SecureServing,
		InsecureServingInfo: c.InsecureServing,
		mode:                c.Mode,
		healthz:             c.Healthz,
		enableMetrics:       c.EnableMetrics,
		enableProfiling:     c.EnableProfiling,
		middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	initGenericAPIServer(s)

	return s, nil
}
