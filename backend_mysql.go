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

func (b *backEndMysql) CreateTable(table Table) error {
	sql := fmt.Sprintf("CREATE TABLE `%s` (%s)",
		table.Name(),
		generateColumnLines(table),
	)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	} else {
		fmt.Println(sql)
	}

	return nil
}

func (b *backEndMysql) DropTable(name string) error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS `%s`", name)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	} else {
		fmt.Println(sql)
	}

	return nil
}

func generateColumnLines(table Table) string {
	columnLines := []string{}
	for _, column := range table.Columns() {
		line := fmt.Sprintf("`%s` %s %s",
			column.Name,
			columnTypeToString(column),
			nullOrNotNull(column),
		)

		if column.IsPrimary {
			line = line + " AUTO_INCREMENT PRIMARY KEY"
		}

		columnLines = append(columnLines, line)
	}

	return strings.Join(columnLines, ", ")
}

func columnTypeToString(column Column) string {
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

func nullOrNotNull(column Column) string {
	if column.AllowNull {
		return "NULL"
	}

	return "NOT NULL"
}
