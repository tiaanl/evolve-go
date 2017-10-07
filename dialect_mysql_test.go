package evolve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDialectMysqlGetCreateTableSQL(t *testing.T) {
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

func TestDialectMysqlGetDropTableSQL(t *testing.T) {
	var err error
	var dropTableSQL string
	dialect := NewDialectMysql()

	dropTableSQL, err = dialect.GetDropTableSQL("test_table")
	assert.NoError(t, err)
	assert.Equal(t, "DROP TABLE `test_table`", dropTableSQL)
}

func TestDialectMysqlColumnTypeToString(t *testing.T) {
	var err error
	dialect := NewDialectMysql()

	columnTypeString, err := dialect.ColumnTypeToString(ColumnTypeString)
	assert.NoError(t, err)
	assert.Equal(t, "VARCHAR", columnTypeString)

	columnTypeInt, err := dialect.ColumnTypeToString(ColumnTypeInteger)
	assert.NoError(t, err)
	assert.Equal(t, "INT", columnTypeInt)

	columnTypeFloat, err := dialect.ColumnTypeToString(ColumnTypeFloat)
	assert.NoError(t, err)
	assert.Equal(t, "FLOAT", columnTypeFloat)

	columnTypeDateTime, err := dialect.ColumnTypeToString(ColumnTypeDateTime)
	assert.NoError(t, err)
	assert.Equal(t, "TIMESTAMP", columnTypeDateTime)

	_, err = dialect.ColumnTypeToString(2345)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "2345")
}

func TestDialectMysqlStringToColumnType(t *testing.T) {
	var err error
	dialect := NewDialectMysql()

	columnTypeInt, err := dialect.StringToColumnType("int")
	assert.NoError(t, err)
	assert.Equal(t, ColumnTypeInteger, columnTypeInt)

	columnTypeString, err := dialect.StringToColumnType("varchar")
	assert.NoError(t, err)
	assert.Equal(t, ColumnTypeString, columnTypeString)

	columnTypeDateTime, err := dialect.StringToColumnType("timestamp")
	assert.NoError(t, err)
	assert.Equal(t, ColumnTypeDateTime, columnTypeDateTime)

	_, err = dialect.StringToColumnType("unknown")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown")
}

func TestDialectMysqlColumnToString(t *testing.T) {
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
