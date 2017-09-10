package evolve

import (
	"fmt"
	"strings"
)

func NewDialectMysql() Dialect {
	return &dialectMysql{}
}

type dialectMysql struct{}

func (d *dialectMysql) GetCreateTableSQL(table Table) (string, error) {
	columnLines, err := d.generateColumnLines(table)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("CREATE TABLE `%s` (%s)", table.Name(), columnLines), nil
}

func (d *dialectMysql) GetDropTableSQL(tableName string) (string, error) {
	return fmt.Sprintf("DROP TABLE `%s`", tableName), nil
}

func (d *dialectMysql) GetAlterTableSQL(tableName string, atc *alterTableColumns) (string, error) {
	lines := []string{}

	for _, name := range atc.toDrop {
		lines = append(lines, fmt.Sprintf("DROP COLUMN `%s`", name))
	}

	for _, column := range atc.toAdd {
		query, err := d.ColumnToString(column)
		if err != nil {
			return "", err
		}
		lines = append(lines, query)
	}

	return fmt.Sprintf("ALTER TABLE `%s` %s", strings.Join(lines, " ")), nil
}

func (d *dialectMysql) StringToColumnType(str string) (ColumnType, error) {
	switch strings.ToLower(str) {
	case "varchar":
		return COLUMN_TYPE_STRING, nil

	case "int":
		return COLUMN_TYPE_INTEGER, nil

	case "timestamp":
		return COLUMN_TYPE_DATE_TIME, nil
	}

	return COLUMN_TYPE_INTEGER, fmt.Errorf("Invalid string to column type. %q", str)
}

func (d *dialectMysql) ColumnTypeToString(columnType ColumnType) (string, error) {
	if columnType == COLUMN_TYPE_INTEGER {
		return "INT", nil
	}

	if columnType == COLUMN_TYPE_STRING {
		return "VARCHAR", nil
	}

	if columnType == COLUMN_TYPE_DATE_TIME {
		return "TIMESTAMP", nil
	}

	return "", fmt.Errorf("Invalid column type to string. %d", columnType)
}

func (d *dialectMysql) ColumnToString(column *Column) (string, error) {
	columnTypeString, err := d.ColumnTypeToString(column.Type)
	if err != nil {
		return "", err
	}

	if column.Type == COLUMN_TYPE_STRING {
		columnTypeString += fmt.Sprintf("(%d)", column.Size)
	}

	line := fmt.Sprintf("`%s` %s %s", column.Name, columnTypeString, d.nullOrNotNull(column))

	if column.AutoIncrement {
		line = line + " AUTO_INCREMENT"
	}

	if column.IsPrimary {
		line = line + " PRIMARY KEY"
	}

	return line, nil
}

func (d *dialectMysql) generateColumnLines(table Table) (string, error) {
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

func (d *dialectMysql) nullOrNotNull(column *Column) string {
	if column.AllowNull {
		return "NULL"
	}

	return "NOT NULL"
}
