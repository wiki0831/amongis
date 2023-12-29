package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// ConnectDB connect to db
func ConnectDB(DATABASE_URL string) {
	var err error
	ctx := context.Background()
	DB, err = pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		fmt.Println("DATABASE_URL:", DATABASE_URL)
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
}
