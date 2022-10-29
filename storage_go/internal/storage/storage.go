package storage

import (
	"context"
	"fmt"

	pgxp "github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxp.Pool

// getConnect - get connect to database from pool
func getConnect() (*pgxp.Conn, error) {
	dbConn, err := db.Acquire(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("error get connect from pool %v", err)
	}
	return dbConn, nil
}

func SetProvider(dbConn *pgxp.Pool) error {
	if dbConn != nil {
		db = dbConn
		return nil
	}

	return fmt.Errorf("db connect is null pointer")
}
