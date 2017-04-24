# evolve-go

## Install

```bash
go get github.com/tiaanl/evolve-go
```

## Use it in your project

Create the `Engine` and then add your migrations.  `Execute` the migrations (up or down) and specify the back end to use. (currently only Mysql)

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

connection, err := sql.Open("...", "...")

err := engine.Up(NewBackEndMysql(connection))
```
