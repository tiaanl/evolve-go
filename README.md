# evolve-go

## Example

```go
connection, err := sql.Open("...", "...")

engine := NewEngine()

engine.AddMigration(NewMigrationWrapper(
    func(schema Schema) {
        schema.CreateTable("users", func(table Table) {
            table.Primary("id")
            table.String("name", 100, false)
            table.String("email", 150, false)
            table.DateTime("created_at", true)
        })
    },
    func(schema Schema) {
        schema.DropTable("users")
    },
))

err := engine.Up(NewBackEndMysql(connection))
```
