package evolve

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMigrationList_AddMigrations(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	for _, backEnd := range []BackEnd{NewBackEndMysql(db), NewBackEndSqlite3(db)} {
		mock.ExpectExec("CREATE TABLE IF NOT EXISTS")
		mock.ExpectExec("INSERT INTO `migrations`").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `migrations`").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `migrations`").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery("SELECT name FROM migrations").WillReturnRows(
			sqlmock.NewRows([]string{"name"}).
				AddRow("migration1").
				AddRow("migration2").
				AddRow("migration3"),
		)

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
}
