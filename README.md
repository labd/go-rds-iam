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

func main() {

	dsn, err := dburl.Parse("postgres://iam-role@rds-server/my-database")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgresql:rdsiam", dsn.DSN)
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
```



