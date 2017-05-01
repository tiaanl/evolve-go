package evolve

type MigrationList interface {
	GetMigrations() ([]string, error)
	AddMigrations(migrationNames ...string) error

	// Make sure the `migrations` table exists.  Returns true if the table already existed.
	EnsureMigrationsTableExists() error
}

func NewMigrationList(backEnd BackEnd) MigrationList {
	return &migrationList{
		backEnd:    backEnd,
		migrations: []string{},
	}
}

type migrationList struct {
	backEnd    BackEnd
	migrations []string
}

func (ml *migrationList) GetMigrations() ([]string, error) {
	migrations := []string{}

	sql := "SELECT name FROM migrations ORDER BY id ASC"

	rows, err := ml.backEnd.Connection().Query(sql)
	if err != nil {
		return migrations, err
	}
	defer rows.Close()

	for rows.Next() {
		var migrationName string
		if err := rows.Scan(&migrationName); err != nil {
			return []string{}, err
		}
		migrations = append(migrations, migrationName)
	}

	return migrations, nil
}

func (ml *migrationList) AddMigrations(migrationNames ...string) error {
	ml.EnsureMigrationsTableExists()

	for _, migrationName := range migrationNames {
		err := ml.backEnd.InsertData("migrations", []string{"name"}, []string{migrationName})
		if err != nil {
			return err
		}
	}

	return nil
}

func (ml *migrationList) EnsureMigrationsTableExists() error {
	// Make sure the migrations table exists.
	table := NewTable("migrations")
	table.Primary("id")
	table.String("name", 100)

	err := ml.backEnd.CreateTableIfNotExists(table)
	if err != nil {
		return err
	}

	return nil
}
