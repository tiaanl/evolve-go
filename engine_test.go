package evolve

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	// Create a sqlite connection for testing.
	//connection, err := sql.Open("sqlite3", ":memory:")
	//if err != nil {
	//	t.Error(err)
	//}

	connection, err := sql.Open("mysql", "root:@tcp(localhost:3306)/evolve-test")
	if err != nil {
		t.Error(err)
	}

	// Create the back end connected to the sqlite3 in memory database.
	backEnd := NewBackEndSqlite3(connection)

	engine := NewEngine(backEnd)

	engine.AddMigration("create_users_table", NewMigrationWrapper(
		func(schema Schema) {
			schema.CreateTable("users", func(table Table) {
				table.Primary("id")
				table.String("name", 100).AllowNull(false)
				table.String("email", 150).AllowNull(false)
				table.String("password", 100).AllowNull(false)
				table.DateTime("created_at").AllowNull(true)
				table.DateTime("updated_at").AllowNull(true)
			})
		},
		func(schema Schema) {
			schema.DropTable("users")
		},
	))

	engine.AddMigration("create_accounts_table", NewMigrationWrapper(
		func(schema Schema) {
			schema.CreateTable("accounts", func(table Table) {
				table.Primary("id")
				table.String("name", 100).AllowNull(false)
				table.String("number", 150).AllowNull(false)
				table.DateTime("created_at").AllowNull(true)
				table.DateTime("updated_at").AllowNull(true)
			})
		},
		func(schema Schema) {
			schema.DropTable("accounts")
		},
	))

	err = engine.Up()
	assert.NoError(t, err)

	err = engine.Down()
	assert.NoError(t, err)
}
