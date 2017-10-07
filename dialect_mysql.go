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

	for _, column := range atc.toAlter {
		query, err := d.ColumnToString(column)
		if err != nil {
			return "", err
		}
		lines = append(lines, fmt.Sprintf("ALTER %s", query))
	}

	for _, column := range atc.toAdd {
		query, err := d.ColumnToString(column)
		if err != nil {
			return "", err
		}
		lines = append(lines, fmt.Sprintf("ADD COLUMN %s", query))
	}

	return fmt.Sprintf("ALTER TABLE `%s` %s", tableName, strings.Join(lines, " ")), nil
}

func (d *dialectMysql) StringToColumnType(str string) (ColumnType, error) {
	switch strings.ToLower(str) {
	case "varchar":
		return ColumnTypeString, nil

	case "int":
		return ColumnTypeInteger, nil

	case "timestamp":
		return ColumnTypeDateTime, nil
	}

	return ColumnTypeInteger, fmt.Errorf("invalid string to column type. %q", str)
}

func (d *dialectMysql) ColumnTypeToString(columnType ColumnType) (string, error) {
	if columnType == ColumnTypeString {
		return "VARCHAR", nil
	}

	if columnType == ColumnTypeInteger {
		return "INT", nil
	}

	if columnType == ColumnTypeFloat {
		return "FLOAT", nil
	}

	if columnType == ColumnTypeDateTime {
		return "TIMESTAMP", nil
	}

	return "", fmt.Errorf("invalid column type to string. %d", columnType)
}

func (d *dialectMysql) ColumnToString(column *Column) (string, error) {
	columnTypeString, err := d.ColumnTypeToString(column.Type)
	if err != nil {
		return "", err
	}

	if column.Type == ColumnTypeString {
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
