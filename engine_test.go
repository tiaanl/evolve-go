package evolve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEngine(t *testing.T) {
	backEnd := NewBackEndMysql(nil)

	engine := NewEngine()

	engine.AddMigration(NewMigrationWrapper(
		func(schema Schema) {
			schema.CreateTable("users", func(table Table) {
				table.Primary("id")
				table.String("tableName", 100, false)
				table.String("email", 150, false)
				table.String("password", 100, false)
				table.DateTime("created_at", true)
				table.DateTime("updated_at", true)
			})
		},
		func(schema Schema) {
			schema.DropTable("users")
		},
	))

	err := engine.Up(backEnd)
	assert.NoError(t, err)

	err = engine.Down(backEnd)
	assert.NoError(t, err)
}
