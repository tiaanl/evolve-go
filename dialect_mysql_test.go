package evolve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDialectMysql_GetCreateTableSQL(t *testing.T) {
	var err error
	var createTableSQL string
	dialect := NewDialectMysql()

	createTableSQL, err = dialect.GetCreateTableSQL(testTable)
	assert.NoError(t, err)
	assert.Contains(t, createTableSQL, "CREATE TABLE")
	assert.Contains(t, createTableSQL, "`test_table`")

	invalidTable := NewTable("invalid")
	invalidTable.AddColumns(testColumnInvalid)
	createTableSQL, err = dialect.GetCreateTableSQL(invalidTable)
	assert.Error(t, err)
}

func TestDialectMysql_ColumnTypeToString(t *testing.T) {
	var err error
	dialect := NewDialectMysql()

	columnTypeInt, err := dialect.ColumnTypeToString(COLUMN_TYPE_INTEGER)
	assert.NoError(t, err)
	assert.Equal(t, "INT", columnTypeInt)

	columnTypeString, err := dialect.ColumnTypeToString(COLUMN_TYPE_STRING)
	assert.NoError(t, err)
	assert.Equal(t, "VARCHAR", columnTypeString)

	columnTypeDateTime, err := dialect.ColumnTypeToString(COLUMN_TYPE_DATE_TIME)
	assert.NoError(t, err)
	assert.Equal(t, "TIMESTAMP", columnTypeDateTime)

	_, err = dialect.ColumnTypeToString(2345)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "2345")
}

func TestDialectMysql_StringToColumnType(t *testing.T) {
	var err error
	dialect := NewDialectMysql()

	columnTypeInt, err := dialect.StringToColumnType("int")
	assert.NoError(t, err)
	assert.Equal(t, COLUMN_TYPE_INTEGER, columnTypeInt)

	columnTypeString, err := dialect.StringToColumnType("varchar")
	assert.NoError(t, err)
	assert.Equal(t, COLUMN_TYPE_STRING, columnTypeString)

	columnTypeDateTime, err := dialect.StringToColumnType("timestamp")
	assert.NoError(t, err)
	assert.Equal(t, COLUMN_TYPE_DATE_TIME, columnTypeDateTime)

	_, err = dialect.StringToColumnType("unknown")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown")
}

func TestDialectMysql_ColumnToString(t *testing.T) {
	var err error
	var str string

	dialect := NewDialectMysql()

	str, err = dialect.ColumnToString(testColumnPrimary)
	assert.NoError(t, err)
	assert.Equal(t, "`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY", str)

	str, err = dialect.ColumnToString(testColumnString)
	assert.NoError(t, err)
	assert.Equal(t, "`username` VARCHAR(50) NULL", str)

	str, err = dialect.ColumnToString(testColumnInteger)
	assert.NoError(t, err)
	assert.Equal(t, "`age` INT NULL", str)

	str, err = dialect.ColumnToString(testColumnDateTime)
	assert.NoError(t, err)
	assert.Equal(t, "`last_login` TIMESTAMP NULL", str)

	// If the column has issues.
	str, err = dialect.ColumnToString(testColumnInvalid)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "2345")
}
