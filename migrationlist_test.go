package evolve

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestMigrationList_AddMigrations(t *testing.T) {
	connection, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Error(err)
		return
	}
	defer connection.Close()

	// Create a back end over the connection.
	backEnd := NewBackEndSqlite3(connection)

	migrationList := NewMigrationList(backEnd)

	err = migrationList.AddMigrations("migration1", "migration2", "migration3")
	if err != nil {
		t.Error(err)
		return
	}

	migrations, err := migrationList.GetMigrations()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(migrations)
}
