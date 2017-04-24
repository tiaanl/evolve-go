# evolve-go

## Install

```bash
go get github.com/tiaanl/evolve-go
```

## Example

Create the `Engine` and then add your migrations.

```go
engine := NewEngine()

engine.AddMigration(NewMigrationWrapper(
    func(schema Schema) {
        schema.CreateTable("users", func(table Table) {
            table.Primary("id")
            table.String("name", 100).NotNull(true)
            table.String("email", 150).NotNull(true)
            table.DateTime("created_at").NotNull(false)
        })
    },
    func(schema Schema) {
        schema.DropTable("users")
    },
))
```

Create the back end you want to run the migration with:

```go
connection, err := sql.Open("...", "...")

// mysql
backEnd := NewBackEndMysql(connection)
```

Now `Execute` the migration and specify the back end to use.

```go
// Up
err := engine.Up()
// or Down
err := engine.Down()
```
