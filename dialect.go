package evolve

type Dialect interface {
	// Return SQL for creating a table.
	GetCreateTableSQL(table Table) (string, error)

	// Return SQL for dropping a table.
	GetDropTableSQL(tableName string) (string, error)

	// Convert string with column type name to ColumnType.
	StringToColumnType(columnName string) (ColumnType, error)

	// Convert a ColumnType to a string with a column type name.
	ColumnTypeToString(columnType ColumnType) (string, error)

	// Convert a Column to it's string representation.
	ColumnToString(column *Column) (string, error)
}
