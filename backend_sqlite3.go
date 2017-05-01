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

	fmt.Println(sql)

	return nil
}

func (b *backEndSqlite3) CreateTableIfNotExists(table Table) error {
	sql := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s)",
		table.Name(),
		generateColumnLinesSqlite3(table),
	)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	fmt.Println(sql)

	return nil
}

func (b *backEndSqlite3) DropTable(name string) error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", name)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	fmt.Println(sql)

	return nil
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
