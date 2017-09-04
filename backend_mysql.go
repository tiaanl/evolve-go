package evolve

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	regex_varchar = regexp.MustCompile(`(\S+)\((\d+)\)( unsigned)?`)
)

func NewBackEndMysql(db *sql.DB) BackEnd {
	return &backEndMysql{
		db: db,
	}
}

type backEndMysql struct {
	db *sql.DB
}

func (b *backEndMysql) ToSQL(s Schema) string {
	var result string

	for _, table := range s.Tables() {
		result += createTableSQLMysql(table) + "\n"
	}

	return result
}

func (b *backEndMysql) BuildSchema() (Schema, error) {
	tables, err := b.buildTablesMysql()
	if err != nil {
		return nil, err
	}

	return NewSchemaWithTables(tables), nil
}

func (b *backEndMysql) Connection() *sql.DB {
	return b.db
}

func (b *backEndMysql) CreateTable(table Table) error {
	sql := fmt.Sprintf("CREATE TABLE `%s` (%s)",
		table.Name(),
		generateColumnLinesMysql(table),
	)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndMysql) CreateTableIfNotExists(table Table) error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s)",
		table.Name(),
		generateColumnLinesMysql(table),
	)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndMysql) DropTable(name string) error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", name)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndMysql) InsertData(table string, columns []string, values []string) error {
	columnsStr := fmt.Sprintf("(`%s`)", strings.Join(columns, "`), (`"))
	valuesStr := fmt.Sprintf("('%s')", strings.Join(values, "'), ('"))

	sql := fmt.Sprintf("INSERT INTO `%s` %s VALUES %s", table, columnsStr, valuesStr)

	_, err := b.db.Exec(sql)

	return err
}

func createTableSQLMysql(table Table) string {
	return fmt.Sprintf("CREATE TABLE `%s` (%s)",
		table.Name(),
		generateColumnLinesMysql(table),
	)
}

func generateColumnLinesMysql(table Table) string {
	columnLines := []string{}
	for _, column := range table.Columns() {
		line := fmt.Sprintf("`%s` %s %s",
			column.Name,
			columnTypeToStringMysql(column),
			nullOrNotNullMysql(column),
		)

		if column.AutoIncrement {
			line = line + " AUTO_INCREMENT"
		}

		if column.IsPrimary {
			line = line + " PRIMARY KEY"
		}

		columnLines = append(columnLines, line)
	}

	return strings.Join(columnLines, ", ")
}

func columnTypeToStringMysql(column *Column) string {
	if column.Type == COLUMN_TYPE_INTEGER {
		return "INT"
	}

	if column.Type == COLUMN_TYPE_UNSIGNED_INTEGER {
		return "INT UNSIGNED"
	}

	if column.Type == COLUMN_TYPE_STRING {
		return fmt.Sprintf("VARCHAR(%d)", column.Size)
	}

	if column.Type == COLUMN_TYPE_DATE_TIME {
		return "TIMESTAMP"
	}

	panic("Incorrect column type")
}

func stringToColumnTypeMysql(columnType string) (ColumnType, error) {
	switch strings.ToLower(columnType) {
	case "varchar":
		return COLUMN_TYPE_STRING, nil

	case "int":
		return COLUMN_TYPE_INTEGER, nil

	case "timstamp":
		return COLUMN_TYPE_DATE_TIME, nil
	}

	return COLUMN_TYPE_INTEGER, fmt.Errorf("Invalid string to column type. %q", columnType)
}

func nullOrNotNullMysql(column *Column) string {
	if column.AllowNull {
		return "NULL"
	}

	return "NOT NULL"
}

func (b *backEndMysql) buildTablesMysql() ([]Table, error) {
	rows, err := b.db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := []Table{}
	var tableName string

	for rows.Next() {
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}

		columns, err := b.buildColumnsMysql(tableName)
		if err != nil {
			return nil, err
		}

		tables = append(tables, NewTableWithColumns(tableName, columns))
	}

	return tables, nil
}

func (b *backEndMysql) buildColumnsMysql(tableName string) ([]*Column, error) {
	rows, err := b.db.Query("DESCRIBE " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := []*Column{}

	var columnName, columnType, columnNull string
	var columnKey, columnDefault, columnExtra sql.NullString

	for rows.Next() {
		err = rows.Scan(
			&columnName,
			&columnType,
			&columnNull,
			&columnKey,
			&columnDefault,
			&columnExtra,
		)
		if err != nil {
			return nil, err
		}

		column := &Column{
			Name:          columnName,
			AllowNull:     strings.ToLower(columnNull) == "yes",
			IsPrimary:     columnKey.Valid && strings.ToLower(columnKey.String) == "pri",
			AutoIncrement: columnExtra.Valid && strings.Contains(strings.ToLower(columnExtra.String), "auto_increment"),
		}

		// varchar(100) unsigned
		if results := regex_varchar.FindStringSubmatch(columnType); len(results) > 0 {
			cType, err := stringToColumnTypeMysql(results[1])
			if err != nil {
				return nil, err
			}

			if cType == COLUMN_TYPE_INTEGER && strings.Contains(results[3], "unsigned") {
				cType = COLUMN_TYPE_UNSIGNED_INTEGER
			}

			column.Type = cType

			cSize, err := strconv.Atoi(results[2])
			if err != nil {
				return nil, err
			}

			column.Size = cSize
		} else if strings.ToLower(columnType) == "timestamp" {
			column.Type = COLUMN_TYPE_DATE_TIME
			column.Size = 0
		} else {
			return nil, fmt.Errorf("Invalid column type. %q", columnType)
		}

		columns = append(columns, column)
	}

	return columns, nil
}
