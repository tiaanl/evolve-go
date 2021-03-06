package evolve

import (
	"fmt"
	"strings"
)

func NewDialectSqlite3() Dialect {
	return &dialectSqlite3{}
}

type dialectSqlite3 struct{}

func (d *dialectSqlite3) GetCreateTableSQL(table Table) (string, error) {
	columnLines, err := d.generateColumnLines(table)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("CREATE TABLE `%s` (%s)", table.Name(), columnLines), nil
}

func (d *dialectSqlite3) GetDropTableSQL(tableName string) (string, error) {
	return fmt.Sprintf("DROP TABLE `%s`", tableName), nil
}

func (d *dialectSqlite3) GetAlterTableSQL(tableName string, atc *alterTableColumns) (string, error) {
	return "", nil
}

func (d *dialectSqlite3) ColumnTypeToString(columnType ColumnType) (string, error) {
	switch columnType {
	case ColumnTypeInteger:
		return "INTEGER", nil

	case ColumnTypeString:
		return "TEXT", nil

	case ColumnTypeDateTime:
		return "TEXT", nil
	}

	return "", fmt.Errorf("invalid column type to column name. %d", columnType)
}

func (d *dialectSqlite3) StringToColumnType(str string) (ColumnType, error) {
	switch strings.ToLower(str) {
	case "varchar":
		return ColumnTypeString, nil

	case "integer":
		return ColumnTypeInteger, nil
	}

	return ColumnTypeInteger, fmt.Errorf("invalid string to column type. %q", str)
}

func (d *dialectSqlite3) ColumnToString(column *Column) (string, error) {
	columnTypeName, err := d.ColumnTypeToString(column.Type)
	if err != nil {
		return "", err
	}

	line := fmt.Sprintf("`%s` %s %s",
		column.Name,
		columnTypeName,
		d.nullOrNotNull(column),
	)

	if column.IsPrimary {
		line = line + " PRIMARY KEY"
	}

	return line, nil
}

func (d *dialectSqlite3) generateColumnLines(table Table) (string, error) {
	columnLines := []string{}
	for _, column := range table.Columns() {
		line, err := d.ColumnToString(column)
		if err != nil {
			return "", err
		}

		columnLines = append(columnLines, line)
	}

	return strings.Join(columnLines, ", "), nil
}

func (d *dialectSqlite3) nullOrNotNull(column *Column) string {
	if column.AllowNull {
		return "NULL"
	}

	return "NOT NULL"
}
