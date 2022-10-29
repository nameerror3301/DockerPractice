package database

import (
	"context"
	"errors"
	"fmt"

	pgxp "github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var (
	ErrConfigData = errors.New("")
	ErrConnect    = func(err error) error { return fmt.Errorf("failed connect to database:\n %v", err) }
)

// postgres://admin:admin@database.timetabler.ex:5432/timedb?sslmode=disable&pool_max_conns=50
func NewProvider(host string, port int, user, password, database string, pool int) (*pgxp.Pool, error) {
	logrus.Debug("connecting by url: ", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d", user, password, host, port, database, pool))

	pgConfigConnect, err := pgxp.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&pool_max_conns=%d", user, password, host, port, database, pool))
	if err != nil {
		return nil, ErrConfigData
	}
	dbPool, err := pgxp.ConnectConfig(context.Background(), pgConfigConnect)
	if err != nil {
		return nil, ErrConnect(err)
	}

	return dbPool, nil
}
