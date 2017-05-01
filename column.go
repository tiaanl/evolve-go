package evolve

type ColumnType int

const (
	COLUMN_TYPE_INTEGER          = 1
	COLUMN_TYPE_UNSIGNED_INTEGER = 2
	COLUMN_TYPE_STRING           = 3
	COLUMN_TYPE_DATE_TIME        = 4
)

type Column struct {
	Name      string
	Type      ColumnType
	Size      int
	AllowNull bool
	IsPrimary bool
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
