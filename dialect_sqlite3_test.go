package evolve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDialectSqlite3_GetCreateTableSQL(t *testing.T) {
	var err error
	var createTableSQL string
	dialect := NewDialectSqlite3()

	createTableSQL, err = dialect.GetCreateTableSQL(testTable)
	assert.NoError(t, err)
	assert.Contains(t, createTableSQL, "CREATE TABLE")
	assert.Contains(t, createTableSQL, "`test_table`")

	invalidTable := NewTable("invalid")
	invalidTable.AddColumns(testColumnInvalid)
	createTableSQL, err = dialect.GetCreateTableSQL(invalidTable)
	assert.Error(t, err)
}

func TestDialectSqlite3_GetDropTableSQL(t *testing.T) {
	var err error
	var dropTableSQL string
	dialect := NewDialectSqlite3()

	dropTableSQL, err = dialect.GetDropTableSQL("test_table")
	assert.NoError(t, err)
	assert.Equal(t, "DROP TABLE `test_table`", dropTableSQL)
}

func TestDialectSqlite3_ColumnTypeToString(t *testing.T) {
	var err error
	dialect := NewDialectSqlite3()

	columnTypeInt, err := dialect.ColumnTypeToString(COLUMN_TYPE_INTEGER)
	assert.NoError(t, err)
	assert.Equal(t, "INTEGER", columnTypeInt)

	columnTypeString, err := dialect.ColumnTypeToString(COLUMN_TYPE_STRING)
	assert.NoError(t, err)
	assert.Equal(t, "TEXT", columnTypeString)

	columnTypeDateTime, err := dialect.ColumnTypeToString(COLUMN_TYPE_DATE_TIME)
	assert.NoError(t, err)
	assert.Equal(t, "TEXT", columnTypeDateTime)

	_, err = dialect.ColumnTypeToString(2345)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "2345")
}

func TestDialectSqlite3_StringToColumnType(t *testing.T) {
	var err error
	dialect := NewDialectSqlite3()

	columnTypeInt, err := dialect.StringToColumnType("integer")
	assert.NoError(t, err)
	assert.Equal(t, COLUMN_TYPE_INTEGER, columnTypeInt)

	columnTypeString, err := dialect.StringToColumnType("varchar")
	assert.NoError(t, err)
	assert.Equal(t, COLUMN_TYPE_STRING, columnTypeString)

	_, err = dialect.StringToColumnType("unknown")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown")
}

func TestDialectSqlite3_ColumnToString(t *testing.T) {
	var err error
	var str string

	dialect := NewDialectSqlite3()

	str, err = dialect.ColumnToString(testColumnPrimary)
	assert.NoError(t, err)
	assert.Equal(t, "`id` INTEGER NOT NULL PRIMARY KEY", str)

	str, err = dialect.ColumnToString(testColumnString)
	assert.NoError(t, err)
	assert.Equal(t, "`username` TEXT NULL", str)

	str, err = dialect.ColumnToString(testColumnInteger)
	assert.NoError(t, err)
	assert.Equal(t, "`age` INTEGER NULL", str)

	str, err = dialect.ColumnToString(testColumnDateTime)
	assert.NoError(t, err)
	assert.Equal(t, "`last_login` TEXT NULL", str)

	// If the column has issues.
	str, err = dialect.ColumnToString(testColumnInvalid)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "2345")
}
