package evolve

import (
	"database/sql"
	"fmt"
	"strings"
)

func NewBackEndSqlite3(db *sql.DB) BackEnd {
	return &backEndSqlite3{
		db: db,
	}
}

type backEndSqlite3 struct {
	db *sql.DB
}

func (b *backEndSqlite3) ToSQL(s Schema) string {
	var result string

	for _, table := range s.Tables() {
		result += createTableSQLSqlite3(table)
	}

	return result
}

func (b *backEndSqlite3) Connection() *sql.DB {
	return b.db
}

func (b *backEndSqlite3) CreateTable(table Table) error {
	sql := fmt.Sprintf("CREATE TABLE `%s` (%s)",
		table.Name(),
		generateColumnLinesSqlite3(table),
	)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndSqlite3) CreateTableIfNotExists(table Table) error {
	sql := createTableSQLSqlite3(table)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndSqlite3) DropTable(name string) error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", name)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndSqlite3) InsertData(table string, columns []string, values []string) error {
	columnsStr := fmt.Sprintf("(`%s`)", strings.Join(columns, "`), (`"))
	valuesStr := fmt.Sprintf("('%s')", strings.Join(values, "'), ('"))

	sql := fmt.Sprintf("INSERT INTO `%s` %s VALUES %s", table, columnsStr, valuesStr)

	_, err := b.db.Exec(sql)

	return err
}

func generateColumnLinesSqlite3(table Table) string {
	columnLines := []string{}
	for _, column := range table.Columns() {
		line := fmt.Sprintf("`%s` %s %s",
			column.Name,
			columnTypeToStringSqlite3(column),
			nullOrNotNullSqlite3(column),
		)

		if column.IsPrimary {
			line = line + " PRIMARY KEY"
		}

		columnLines = append(columnLines, line)
	}

	return strings.Join(columnLines, ", ")
}

func columnTypeToStringSqlite3(column *Column) string {
	if column.Type == COLUMN_TYPE_INTEGER {
		return "INTEGER"
	}

	if column.Type == COLUMN_TYPE_UNSIGNED_INTEGER {
		return "INTEGER"
	}

	if column.Type == COLUMN_TYPE_STRING {
		return "TEXT"
	}

	if column.Type == COLUMN_TYPE_DATE_TIME {
		return "TIMESTAMP"
	}

	panic("Incorrect column type")
}

func nullOrNotNullSqlite3(column *Column) string {
	if column.AllowNull {
		return "NULL"
	}

	return "NOT NULL"
}

func createTableSQLSqlite3(table Table) string {
	return fmt.Sprintf("CREATE TABLE `%s` (%s)",
		table.Name(),
		generateColumnLinesSqlite3(table),
	)
}
