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
