package evolve

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	regexVarchar = regexp.MustCompile(`(\S+)\((\d+)\)( unsigned)?`)
)

func NewBackEndMysql(db *sql.DB) BackEnd {
	return &backEndMysql{
		db:      db,
		dialect: NewDialectMysql(),
	}
}

type backEndMysql struct {
	db      *sql.DB
	dialect Dialect
}

func (b *backEndMysql) Connection() *sql.DB {
	return b.db
}

func (b *backEndMysql) Dialect() Dialect {
	return b.dialect
}

func (b *backEndMysql) CreateTable(table Table) error {
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

func (b *backEndMysql) DropTable(name string) error {
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

func (b *backEndMysql) InsertData(table string, columns []string, values []string) error {
	columnsStr := fmt.Sprintf("(`%s`)", strings.Join(columns, "`), (`"))
	valuesStr := fmt.Sprintf("('%s')", strings.Join(values, "'), ('"))

	query := fmt.Sprintf("INSERT INTO `%s` %s VALUES %s", table, columnsStr, valuesStr)

	_, err := b.db.Exec(query)

	return err
}

func (b *backEndMysql) BuildSchema() (Schema, error) {
	tables, err := b.buildTablesMysql()
	if err != nil {
		return nil, err
	}

	return NewSchemaWithTables(tables), nil
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
		err = rows.Scan(&columnName, &columnType, &columnNull, &columnKey, &columnDefault, &columnExtra)
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
		if results := regexVarchar.FindStringSubmatch(columnType); len(results) > 0 {
			cType, err := b.dialect.StringToColumnType(results[1])
			if err != nil {
				return nil, err
			}

			column.Type = cType

			cSize, err := strconv.Atoi(results[2])
			if err != nil {
				return nil, err
			}

			column.Size = cSize
		} else if strings.ToLower(columnType) == "timestamp" {
			column.Type = ColumnTypeDateTime
			column.Size = 0
		} else {
			return nil, fmt.Errorf("invalid column type. %q", columnType)
		}

		columns = append(columns, column)
	}

	return columns, nil
}
