package service

import (
	"context"
	"prep/internal/rest"
	"prep/internal/storage"
	"prep/pkg/sender"
	"prep/pkg/server"
	"prep/pkg/utility"

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

	logrus.Infof("Start service with configs:\n\t- Server address: %s\n\t- Server port: %d\n\t- Server version: %s",
		config.Service.Address,
		config.Service.Port,
		config.Service.Version,
	)
	logrus.Infof("Parameters out server:\n\t- %s:%d",
		config.Out.Address,
		config.Out.Port,
	)

	store, err := sender.NewClient(config.Out.Address, config.Out.Port)
	if err != nil {
		logrus.Panic("Fail create perparer")
	}

	storage.SetPrep(store)

	logrus.Infof("Base service version: %s", config.Service.Version)

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
