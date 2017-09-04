package evolve

import "database/sql"

type BackEnd interface {
	// Return SQL for the given schema.
	ToSQL(s Schema) string

	// Return the connection that this back end represents.
	Connection() *sql.DB

	// Create a table on the connection.
	CreateTable(table Table) error

	// Create a table on the connection if it doesn't already exits.
	CreateTableIfNotExists(table Table) error

	// Drop a table on the connection.
	DropTable(name string) error

	// Insert data into the the given table.
	InsertData(table string, columns []string, values []string) error
}
