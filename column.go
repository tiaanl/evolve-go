package evolve

type ColumnType int

const (
	COLUMN_TYPE_INTEGER ColumnType = iota
	COLUMN_TYPE_UNSIGNED_INTEGER
	COLUMN_TYPE_STRING
	COLUMN_TYPE_DATE_TIME
)

type Column struct {
	Name          string
	Type          ColumnType
	Size          int
	AllowNull     bool
	IsPrimary     bool
	AutoIncrement bool
}

func newFluentColumn(column *Column) *fluentColumn {
	return &fluentColumn{
		column: column,
	}
}

type fluentColumn struct {
	column *Column
}

func (f *fluentColumn) AllowNull(allowNull bool) *fluentColumn {
	f.column.AllowNull = allowNull
	return f
}

func (f *fluentColumn) IsPrimary(isPrimary bool) *fluentColumn {
	f.column.IsPrimary = isPrimary
	return f
}

func (f *fluentColumn) AutoIncrement(autoIncrement bool) *fluentColumn {
	f.column.AutoIncrement = autoIncrement
	return f
}
