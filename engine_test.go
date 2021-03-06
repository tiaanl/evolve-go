package evolve

/*
import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestEngine(t *testing.T) {
	db, mock, _ := sqlmock.New()
	//if err != nil {
	//	t.Error(err)
	//}
	//defer db.Close()

	//for _, backEnd := range []BackEnd{NewBackEndSqlite3(db), NewBackEndMysql(db)} {
	for _, backEnd := range []BackEnd{NewBackEndMysql(db)} {
		mock.ExpectExec("CREATE TABLE IF NOT EXISTS `migrations`").WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectQuery("SELECT name FROM migrations ORDER BY name ASC").WillReturnRows(sqlmock.NewRows([]string{"name"}))
		mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("CREATE TABLE IF NOT EXISTS `migrations`").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DROP TABLE IF EXISTS").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec("DROP TABLE IF EXISTS").WillReturnResult(sqlmock.NewResult(1, 1))

		engine := NewEngine(backEnd)

		engine.AddMigration("create_users_table", NewMigrationWrapper(
			func(changeSet ChangeSet) {
				changeSet.CreateTableWithFunc("users", func(table Table) {
					table.Primary("id")
					table.String("name", 100).AllowNull(false)
					table.String("email", 150).AllowNull(false)
					table.String("password", 100).AllowNull(false)
					table.DateTime("created_at").AllowNull(true)
					table.DateTime("updated_at").AllowNull(true)
				})
			},
			func(changeSet ChangeSet) {
				changeSet.DropTable("users")
			},
		))

		engine.AddMigration("create_accounts_table", NewMigrationWrapper(
			func(changeSet ChangeSet) {
				changeSet.CreateTableWithFunc("accounts", func(table Table) {
					table.Primary("id")
					table.String("name", 100).AllowNull(false)
					table.String("number", 150).AllowNull(false)
					table.DateTime("created_at").AllowNull(true)
					table.DateTime("updated_at").AllowNull(true)
				})
			},
			func(changeSet ChangeSet) {
				changeSet.DropTable("accounts")
			},
		))

		//err = engine.Update()
		// assert.NoError(t, err)

		backEnd.DropTable("migrations")
	}
}
*/
