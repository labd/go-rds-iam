package rdsiam

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/lib/pq"
)

type Driver struct {
	// Wrapped driver
	wd driver.Driver
}

func init() {
	sql.Register("postgresql-rdsiam", &Driver{
		wd: pq.Driver{},
	})
}

// Open opens a new connection to the database. name is a connection string.
// Most users should only use it through database/sql package from the standard
// library.
func (d *Driver) Open(name string) (driver.Conn, error) {
	ctx := context.Background()
	dsn, err := createNewDSN(ctx, name)
	if err != nil {
		return nil, err
	}
	return d.wd.Open(dsn)
}
