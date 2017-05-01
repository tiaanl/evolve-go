package evolve

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFluentColumn_IsPrimary(t *testing.T) {
	column := Column{
		Name: "column",
		Type: COLUMN_TYPE_STRING,
		Size: 100,
		AllowNull: true,
		IsPrimary: false,
	}

	fc := newFluentColumn(&column)
	assert.Equal(t, false, fc.column.IsPrimary)

	fc.IsPrimary(true)
	assert.Equal(t, true, fc.column.IsPrimary)

	fc.IsPrimary(false)
	assert.Equal(t, false, fc.column.IsPrimary)
}

func TestFluentColumn_AllowNull(t *testing.T) {
	column := Column{
		Name: "column",
		Type: COLUMN_TYPE_STRING,
		Size: 100,
		AllowNull: true,
		IsPrimary: false,
	}

	fc := newFluentColumn(&column)
	assert.Equal(t, true, fc.column.AllowNull)

	fc.AllowNull(false)
	assert.Equal(t, false, fc.column.AllowNull)

	fc.AllowNull(true)
	assert.Equal(t, true, fc.column.AllowNull)
}
