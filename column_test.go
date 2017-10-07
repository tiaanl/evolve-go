package evolve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFluentColumnIsPrimary(t *testing.T) {
	column := Column{
		Name:      "column",
		Type:      ColumnTypeString,
		Size:      100,
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

func TestFluentColumnAllowNull(t *testing.T) {
	column := Column{
		Name:      "column",
		Type:      ColumnTypeString,
		Size:      100,
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
