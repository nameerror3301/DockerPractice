package service

import (
	"context"
	"gateway/internal/prepare"
	"gateway/internal/rest"
	"gateway/pkg/sender"
	"gateway/pkg/server"
	"gateway/pkg/utility"

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

	logrus.Infof("Start service with configs:\n\t- Server address: %s\n\t- Server port: %d\n\t- Server version: %s",
		config.Service.Address,
		config.Service.Port,
		config.Service.Version,
	)
	logrus.Infof("Parameters out server:\n\t- http://%s:%d",
		config.Out.Address,
		config.Out.Port,
	)

	store, err := sender.NewClient(config.Out.Address, config.Out.Port)
	if err != nil {
		logrus.Panic("Fail create perparer")
	}

	prepare.SetPrep(store)

	preparer, err := sender.NewClient(config.Out.Address, config.Out.Port)
	if err != nil {
		logrus.Panic("Fail create perparer")
	}

	prepare.SetPrep(preparer)

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
