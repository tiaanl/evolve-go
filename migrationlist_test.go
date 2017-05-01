package evolve

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrationList_AddMigrations(t *testing.T) {
	connection, err := sql.Open("mysql", "root:@tcp(localhost:3306)/evolve-test")
	if err != nil {
		t.Error(err)
	}
	defer connection.Close()

	backEnd := NewBackEndMysql(connection)

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

	checks := map[string]bool{}

	// Make sure all the migrations are there.
	for _, migrationName := range migrations {
		checks[migrationName] = true
	}

	assert.Equal(t, true, checks["migration1"])
	assert.Equal(t, true, checks["migration2"])
	assert.Equal(t, true, checks["migration3"])

	backEnd.DropTable("migrations")
}
