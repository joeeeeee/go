package options

import (
	genericoptions "github.com/joe/iam/internal/pkg/options"
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
)

type Options struct {
	RPCServer               string                                 `json:"rpcserver"      mapstructure:"rpcserver"`
	ClientCA                string                                 `json:"client-ca-file" mapstructure:"client-ca-file"`
	GenericServerRunOptions *genericoptions.ServerRunOptions       `json:"server"         mapstructure:"server"`
	//InsecureServing         *genericoptions.InsecureServingOptions `json:"insecure"       mapstructure:"insecure"`
	//SecureServing           *genericoptions.SecureServingOptions   `json:"secure"         mapstructure:"secure"`
	RedisOptions            *genericoptions.RedisOptions           `json:"redis"          mapstructure:"redis"`
	//FeatureOptions          *genericoptions.FeatureOptions         `json:"feature"        mapstructure:"feature"`
	//Log                     *log.Options                           `json:"log"            mapstructure:"log"`
	//AnalyticsOptions        *analytics.AnalyticsOptions            `json:"analytics"      mapstructure:"analytics"`
}


func NewOptions() *Options {
	return &Options {
		RPCServer: "127.0.0.1:8001",
		ClientCA: "",
		GenericServerRunOptions : genericoptions.NewServerRunOptions(),
	}
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	return fss
}

func (o *Options) Validate() []error {
	return nil
}



