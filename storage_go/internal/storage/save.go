package storage

import (
	"context"

	"github.com/sirupsen/logrus"
)

func SavePayload(ctx *context.Context, data []byte) error {
	dbConn, err := getConnect()
	if err != nil {
		return err
	}
	defer dbConn.Release()

	logrus.Debug("Data for save %s", string(data))

	if _, err := dbConn.Exec(*ctx, "INSERT INTO data_heap (payload) VALUES ($1)", data); err != nil {
		return err
	}

	return nil
}
