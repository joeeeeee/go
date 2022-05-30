package apiserver

import (
	"github.com/joe/iam/internal/apiserver/config"
	"github.com/joe/iam/internal/apiserver/store"
	"github.com/joe/iam/internal/apiserver/store/mysql"
	genericoptions "github.com/joe/iam/internal/pkg/options"
	genericapiserver "github.com/joe/iam/internal/pkg/server"
	"google.golang.org/grpc"
)

type apiServer struct {
	genericAPIServer interface{}
}

// ExtraConfig defines extra configuration for the iam-apiserver.
type ExtraConfig struct {
	Addr         string
	MaxMsgSize   int
	ServerCert   interface{}
	mysqlOptions *genericoptions.MySQLOptions
	// etcdOptions      *genericoptions.EtcdOptions
}

type completedExtraConfig struct {
	*ExtraConfig
}


func CreateApiServer (cfg *config.Config) (*apiServer, error){
	_, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	extraConfig, err := buildExtraConfig(cfg)
	if err != nil {
		return nil, err
	}

	//genericServer, err := genericConfig.Complete().New()
	if err != nil {
		return nil, err
	}
	_, err = extraConfig.complete().New()
	if err != nil {
		return nil, err
	}

	server := &apiServer{
		genericAPIServer: "",
	}

	return server , nil
}

//nolint: unparam
func buildExtraConfig(cfg *config.Config) (*ExtraConfig, error) {
	return &ExtraConfig{
		//Addr:         fmt.Sprintf("%s:%d", cfg.GRPCOptions.BindAddress, cfg.GRPCOptions.BindPort),
		//MaxMsgSize:   cfg.GRPCOptions.MaxMsgSize,
		//ServerCert:   cfg.SecureServing.ServerCert,
		mysqlOptions: cfg.MySQLOptions,
		// etcdOptions:      cfg.EtcdOptions,
	}, nil
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericapiserver.Config, lastErr error) {
	genericConfig = genericapiserver.NewConfig()
	if lastErr = cfg.GenericServerRunOptions.ApplyTo(genericConfig); lastErr != nil {
		return
	}
	//
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

func (s *apiServer) PrepareRun() error {
	initRouter()

	return nil
}

// Complete fills in any fields not set that are required to have valid data and can be derived from other fields.
func (c *ExtraConfig) complete() *completedExtraConfig {
	if c.Addr == "" {
		c.Addr = "127.0.0.1:8081"
	}

	return &completedExtraConfig{c}
}

// New create a grpcAPIServer instance.
func (c *completedExtraConfig) New() (interface{}, error) {

	storeIns, _ := mysql.GetMySQLFactoryOr(c.mysqlOptions)
	// storeIns, _ := etcd.GetEtcdFactoryOr(c.etcdOptions, nil)
	store.SetClient(storeIns)

	grpcServer := grpc.NewServer()
	//cacheIns, err := cachev1.GetCacheInsOr(storeIns)
	//if err != nil {
	//	log.Fatalf("Failed to get cache instance: %s", err.Error())
	//}


	return &grpcAPIServer{grpcServer, c.Addr}, nil
}

