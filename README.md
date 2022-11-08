# RDS IAM SQL Driver for Go


## Usage

Make sure you have set the env var `AWS_REGION` to the region you want to
connect to.


## Example

```go
package main

import (
	"context"
	"database/sql"
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
	db, err := dburl.Open("postgresql-rdsiam://iam-role@rds-server/my-database")
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

// alternative skips registering the url with dburl
func alternative() {
	dsn, err := dburl.Parse("postgres://iam-role@rds-server/my-database")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgresql-rdsiam", dsn.DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
```



