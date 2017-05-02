package evolve

import (
	"database/sql"
	"fmt"
	"strings"
)

func NewBackEndMysql(db *sql.DB) BackEnd {
	return &backEndMysql{
		db: db,
	}
}

type backEndMysql struct {
	db *sql.DB
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

func generateColumnLinesMysql(table Table) string {
	columnLines := []string{}
	for _, column := range table.Columns() {
		line := fmt.Sprintf("`%s` %s %s",
			column.Name,
			columnTypeToStringMysql(column),
			nullOrNotNullMysql(column),
		)

		if column.IsPrimary {
			line = line + " AUTO_INCREMENT PRIMARY KEY"
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

func nullOrNotNullMysql(column *Column) string {
	if column.AllowNull {
		return "NULL"
	}

	return "NOT NULL"
}
