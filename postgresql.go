package main

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Conn_to_DB() error {

	var err error
	pool, err = pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/my_DB")
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		return err
	}
	//defer pool.Close()

	// Verify the connection
	if err := pool.Ping(ctx); err != nil {
		fmt.Println("Unable to ping database:", err)
		return err
	}

	fmt.Println("Connected to PostgreSQL database!")
	return nil
}
