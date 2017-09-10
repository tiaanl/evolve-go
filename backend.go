package evolve

import (
	"database/sql"
)

type BackEnd interface {
	// Return the connection that this back end represents.
	Connection() *sql.DB

	// Return the dialect used by this backend.
	Dialect() Dialect

	// Create a table on the connection.
	CreateTable(table Table) error

	// Drop a table on the connection.
	DropTable(name string) error

	// Insert data into the the given table.
	InsertData(table string, columns []string, values []string) error

	// Return a Schema from the current connection.
	BuildSchema() (Schema, error)
}
