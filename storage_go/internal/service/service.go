package service

import (
	"context"
	"store/internal/rest"
	"store/internal/storage"
	"store/pkg/database"
	"store/pkg/server"
	"store/pkg/utility"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// Пример функции в качестве callback
func ServerHttpCallback() {
	logrus.Info("Server state is down")
}

func Service(ctx context.Context) error {
	logrus.Info("Service")

	// Use export DRIVER_CONFIG = {json,yaml,toml}
	config, err := utility.ReadConfig()
	if err != nil {
		return err
	}

	logrus.Infof("Base service version: %s", config.Service.Version)

	//
	// Database provider
	//
	logrus.Info("Connecting to database")
	var dbConf = config.Database
	var db *pgxpool.Pool
	for i := 0; true; i++ {
		db, err = database.NewProvider(dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Database, dbConf.Pool)
		if err == nil {
			break
		} else if i == 10 && err != nil {
			logrus.Fatal(err)
		}

		logrus.Info("...")
		time.Sleep(time.Second * 2)
	}
	logrus.Info("Database connect success")

	storage.SetProvider(db)
	// defer db.Close()

	if _, err = server.NewServer(
		nil,
		config.Service.Address,
		config.Service.Port,
		ServerHttpCallback,
		rest.InitHandlers(),
	); err != nil {
		return err
	}

	return nil
}
