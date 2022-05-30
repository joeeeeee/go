package mysql

import (
	"fmt"
	"github.com/joe/iam/internal/apiserver/store"
	genericoptions "github.com/joe/iam/internal/pkg/options"
	"github.com/joe/iam/pkg/db"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sync"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) Users() store.UserStore {
	return newUsers(ds)
}

func (ds *datastore) Secrets() store.SecretStore {
	//return newSecrets(ds)
	return nil
}

func (ds *datastore) Policies() store.PolicyStore {
	return newPolicy(ds)
}



func (ds *datastore) Close() error {
	db, err := ds.db.DB()

	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("fail to get mysql store factory")
	}

	var err error
	var dbIns *gorm.DB

	once.Do(func() {
		options := &db.Options {
			Host:                  viper.GetString("mysql.host"),
			Username:              viper.GetString("mysql.username"),
			Password:              viper.GetString("mysql.password"),
			Database:              viper.GetString("mysql.database"),
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
		}

		dbIns, err = db.New(options)

		mysqlFactory = &datastore{dbIns}
	})

	return mysqlFactory , nil
}
