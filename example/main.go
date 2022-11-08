package main

import (
	"context"
	"fmt"

	"github.com/xo/dburl"

	_ "github.com/labd/go-rds-iam/rdsiam"
)

// Register the custom schema with dburl
func init() {
	dburl.Register(dburl.Scheme{
		Driver:    "postgresql-rdsiam",
		Generator: dburl.GenPostgres,
		Transport: dburl.TransportUnix,
		Opaque:    false,
		Aliases:   []string{"postgresql-rdsiam"},
		Override:  "",
	})
}

func main() {

	db, err := dburl.Open("postgres://iam-role@rds-server/my-database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx := context.Background()
	rows, err := db.QueryContext(ctx, `SELECT 'foobar' AS one`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var value string
		rows.Scan(&value)
		fmt.Println("Selected:", value)
	}
}
