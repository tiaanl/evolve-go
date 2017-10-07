package evolve

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBackEndMysqlCreateTable(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	backEnd := NewBackEndMysql(db)

	mock.ExpectExec("CREATE TABLE `test_table`").WillReturnResult(sqlmock.NewResult(0, 0))

	err = backEnd.CreateTable(testTable)
	assert.NoError(t, err)
}

func TestBackEndMysqlDropTable(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	backEnd := NewBackEndMysql(db)

	mock.ExpectExec("DROP TABLE `test_table`").WillReturnResult(sqlmock.NewResult(0, 0))

	err = backEnd.DropTable("test_table")
	assert.NoError(t, err)
}

func TestBackEndMysqlInsertData(t *testing.T) {
	var err error

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	backEnd := NewBackEndMysql(db)

	mock.ExpectExec("INSERT INTO `test_table`").WillReturnResult(sqlmock.NewResult(1, 1))

	err = backEnd.InsertData(
		"test_table",
		[]string{"username", "age", "last_login"},
		[]string{"some username", "10", "2001-01-01 12:34:56"},
	)
	assert.NoError(t, err)
}
