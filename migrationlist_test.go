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
		mock.ExpectExec("CREATE TABLE IF NOT EXISTS").WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectExec("INSERT INTO `migrations`").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `migrations`").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("INSERT INTO `migrations`").WillReturnResult(sqlmock.NewResult(1, 1))

		migrationList := NewMigrationList(backEnd)

		err = migrationList.Add("migration1", "migration2", "migration3")
		if err != nil {
			t.Error(err)
			return
		}

		for _, name := range []string{"migration1", "migration2", "migration3"} {
			migrationExists, err := migrationList.Exists(name)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, true, migrationExists)
		}

		backEnd.DropTable("migrations")
	}
}
