package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)


func DB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(
		context.Background(),
		"postgresql://postgres:1@localhost:5432/soal5?sslmode=disable",
	)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}
