package evolve

import (
	"database/sql"
	"fmt"
	"strings"
)

func NewBackEndSqlite3(db *sql.DB) BackEnd {
	return &backEndSqlite3{
		db:      db,
		dialect: NewDialectMysql(),
	}
}

type backEndSqlite3 struct {
	db      *sql.DB
	dialect Dialect
}

func (b *backEndSqlite3) BuildSchema() (Schema, error) {
	return nil, nil
}

func (b *backEndSqlite3) Connection() *sql.DB {
	return b.db
}

func (b *backEndSqlite3) CreateTable(table Table) error {
	createTableSQL, err := b.dialect.GetCreateTableSQL(table)
	if err != nil {
		return err
	}

	if b.db != nil {
		_, err := b.db.Exec(createTableSQL)
		return err
	}

	return nil
}

func (b *backEndSqlite3) DropTable(name string) error {
	dropTableSQL, err := b.dialect.GetDropTableSQL(name)
	if err != nil {
		return err
	}

	if b.db != nil {
		_, err := b.db.Exec(dropTableSQL)
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
