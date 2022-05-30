package options

import (
	genericoptions "github.com/joe/iam/internal/pkg/options"
)


type Options struct {
	GenericServerRunOptions *genericoptions.ServerRunOptions `json:"server" mapstructure:"server"`
	GRPCOptions interface{}
	InsecureServing interface{}
	SecureServing interface{}
	MySQLOptions *genericoptions.MySQLOptions `json:"server" mapstructure:"server"`
	RedisOptions interface{}
	JwtOptions interface{}
	Log interface{}
	FeatureOptions interface{}
}

func NewOptions() *Options {
	return &Options{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		GRPCOptions:             genericoptions.NewGRPCOptions(),
		InsecureServing:         nil,
		SecureServing:           nil,
		MySQLOptions:            genericoptions.NewMySQLOptions(),
		RedisOptions:            nil,
		JwtOptions:              nil,
		Log:                     nil,
		FeatureOptions:          nil,
	}
}
