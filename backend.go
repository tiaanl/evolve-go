package evolve

import (
	"database/sql"
)

type BackEnd interface {
	// Return SQL for the given schema.
	ToSQL(s Schema) (string, error)

	// Return a Schema from the current connection.
	BuildSchema() (Schema, error)

	// Return the connection that this back end represents.
	Connection() *sql.DB

	// Create a table on the connection.
	CreateTable(table Table) error

	// Drop a table on the connection.
	DropTable(name string) error

	// Insert data into the the given table.
	InsertData(table string, columns []string, values []string) error
}

/*
func (b *backEndImpl) ToSQL(s Schema) string {
	var result string

	for _, table := range s.Tables() {
		result += b.dialect.GetCreateTableSQL(table) + "\n"
	}

	return result
}

func (b *backEndImpl) BuildSchema() (Schema, error) {
	tables, err := buildTablesMysql()
	if err != nil {
		return nil, err
	}

	return NewSchemaWithTables(tables), nil
}

func (b *backEndImpl) Connection() *sql.DB {
	return b.db
}

func (b *backEndImpl) CreateTable(table Table) error {
	sql := b.dialect.GetCreateTableSQL(table)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndImpl) DropTable(name string) error {
	sql := b.dialect.GetDropTableSQL(name)

	if b.db != nil {
		_, err := b.db.Exec(sql)
		return err
	}

	return nil
}

func (b *backEndImpl) InsertData(table string, columns []string, values []string) error {
	columnsStr := fmt.Sprintf("(`%s`)", strings.Join(columns, "`), (`"))
	valuesStr := fmt.Sprintf("('%s')", strings.Join(values, "'), ('"))

	sql := fmt.Sprintf("INSERT INTO `%s` %s VALUES %s", table, columnsStr, valuesStr)

	_, err := b.db.Exec(sql)

	return err
}
*/
