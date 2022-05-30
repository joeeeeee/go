package authzserver

import (
	"context"
	"github.com/joe/iam/internal/authzserver/config"
	"github.com/joe/iam/internal/authzserver/load/cache"
	"github.com/joe/iam/internal/authzserver/store/apiserver"
	genericoptions "github.com/joe/iam/internal/pkg/options"
	genericapiserver "github.com/joe/iam/internal/pkg/server"
	"github.com/joe/iam/pkg/storage"
)

type authzServer struct {
	gs           string
	rpcServer    string
	clientCA     string
	redisOptions  *genericoptions.RedisOptions
	genericAPIServer *genericapiserver.GenericAPIServer
	analyticsOptions string
}

func createAuthzServer(cfg *config.Config) (*authzServer, error) {

	genericConfig, err := buildGenericConfig(cfg)

	if err != nil {
		return nil, err
	}

	genericServer, err := genericConfig.Complete().New()
	server := &authzServer{
		gs:               "",
		redisOptions:     cfg.RedisOptions,
		analyticsOptions: nil,
		rpcServer:        cfg.RPCServer,
		clientCA:         cfg.ClientCA,
		genericAPIServer: genericServer,
	}

	return server, nil
}

func (s *authzServer) PrepareRun() {
	s.initialize()
}

func (s *authzServer) initialize() {
	ctx, cancel := context.WithCancel(context.Background())

	go storage.ConnectToRedis(ctx, s.buildStorageConfig())

	cacheIns, err := cache.GetCacheInsOr(apiserver.GetAPIServerFactoryOrDie(s.rpcServer, s.clientCA))
}


func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}

	//if lastErr = cfg.FeatureOptions.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.SecureServing.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}
	//
	//if lastErr = cfg.InsecureServing.ApplyTo(genericConfig); lastErr != nil {
	//	return
	//}

	return
}


func (s *authzServer) buildStorageConfig() *storage.Config {
	return &storage.Config{
		Host:                  s.redisOptions.Host,
		Port:                  s.redisOptions.Port,
		Addrs:                 s.redisOptions.Addrs,
		MasterName:            s.redisOptions.MasterName,
		Username:              s.redisOptions.Username,
		Password:              s.redisOptions.Password,
		Database:              s.redisOptions.Database,
		MaxIdle:               s.redisOptions.MaxIdle,
		MaxActive:             s.redisOptions.MaxActive,
		Timeout:               s.redisOptions.Timeout,
		EnableCluster:         s.redisOptions.EnableCluster,
		UseSSL:                s.redisOptions.UseSSL,
		SSLInsecureSkipVerify: s.redisOptions.SSLInsecureSkipVerify,
	}
}


